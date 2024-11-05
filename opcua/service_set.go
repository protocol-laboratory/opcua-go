package opcua

import (
	"errors"

	"github.com/protocol-laboratory/opcua-go/opcua/uamsg"
	"github.com/protocol-laboratory/opcua-go/opcua/util"
)

func (secChan *SecureChannel) handleOpenSecureChannelRequest(req *uamsg.Message) (*uamsg.Message, error) {
	svc, err := getService[uamsg.OpenSecureChannelRequest](req)
	if err != nil {
		return nil, err
	}

	err = secChan.interceptor.BeforeOpenSecureChannel(secChan.conn, svc)
	if err != nil {
		return nil, err
	}

	err = secChan.interceptor.AfterOpenSecureChannel(secChan.conn, svc, nil)
	if err != nil {
		return nil, err
	}

	return nil, errors.New("open new secure channel is not supported")
}

func (secChan *SecureChannel) handleCloseSecureChannelRequest(req *uamsg.Message) (*uamsg.Message, error) {
	svc, err := getService[uamsg.CloseSecureChannelRequest](req)
	if err != nil {
		return nil, err
	}

	err = secChan.interceptor.BeforeCloseSecureChannel(secChan.conn, svc)
	if err != nil {
		return nil, err
	}

	err = secChan.interceptor.AfterCloseSecureChannel(secChan.conn, svc, nil)
	if err != nil {
		return nil, err
	}

	// TODO return a dummy error just for closing connection, should refactor after ErrorCode redesign
	return nil, errors.New("secure channel closed")
}

func (secChan *SecureChannel) handleCreateSessionRequest(req *uamsg.Message) (*uamsg.Message, error) {
	svc, err := getService[uamsg.CreateSessionRequest](req)
	if err != nil {
		return nil, err
	}

	err = secChan.interceptor.BeforeCreateSession(secChan.conn, svc)
	if err != nil {
		return nil, err
	}
	session := newSession(svc.SessionName, svc.RequestedSessionTimeout, svc.MaxResponseMessageSize)
	token := getUniqueSessionAuthenticationToken()
	bytes, ok := token.Identifier.([]byte)
	if !ok {
		return nil, ErrInvalidMessageBody
	}
	secChan.sessionManager.add(string(bytes), session)

	ep := &uamsg.EndpointDescription{
		EndpointUrl: svc.EndpointUrl,
		Server: &uamsg.ApplicationDescription{
			ApplicationName: &uamsg.LocalizedText{},
		},
		SecurityMode:      uamsg.MessageSecurityModeNone,
		SecurityPolicyUri: uamsg.SecurityPolicyUriNone,
		UserIdentityTokens: []*uamsg.UserTokenPolicy{
			{
				PolicyId:  "username",
				TokenType: uamsg.UserTokenTypeUsername,
			},
		},
		TransportProfileUri: "http://opcfoundation.org/UA-Profile/Transport/uatcp-uasc-uabinary",
		SecurityLevel:       2,
	}

	respSvc := &uamsg.CreateSessionResponse{
		Header: &uamsg.ResponseHeader{
			Timestamp:     util.GetCurrentUaTimestamp(),
			RequestHandle: svc.Header.RequestHandle,
			ServiceResult: uint32(uamsg.ErrorCodeGood),
			ServiceDiagnostics: &uamsg.DiagnosticInfo{
				EncodingMask: 0x00,
			},
			StringTable: nil,
			AdditionalHeader: &uamsg.ExtensionObject{
				TypeId: &uamsg.NodeId{
					EncodingType: uamsg.TwoByte,
					Identifier:   byte(0),
				},
				Encoding: 0x00,
			},
		},
		SessionId:             &session.sessionId,
		AuthenticationToken:   token,
		RevisedSessionTimeout: uamsg.Duration(session.requestedSessionTimeout.Milliseconds()),
		ServerNonce:           session.getServerNonce(),
		MaxRequestMessageSize: session.maxResponseMessageSize,
		// TODO endpoints description needed
		ServerEndpoints: []*uamsg.EndpointDescription{ep},
		ServerSignature: &uamsg.SignatureData{
			Algorithm: "",
			Signature: nil,
		},
	}

	err = secChan.interceptor.AfterCreateSession(secChan.conn, svc, respSvc)
	if err != nil {
		return nil, err
	}

	rsp := &uamsg.Message{
		MessageHeader: &uamsg.MessageHeader{
			MessageType:     uamsg.MsgMessageType,
			SecureChannelId: &secChan.channelId,
		},
		SecurityHeader: &uamsg.SymmetricSecurityHeader{
			TokenId: secChan.getCurrentTokenId(),
		},
		SequenceHeader: &uamsg.SequenceHeader{
			SequenceNumber: secChan.getNextSequenceNumber(),
			RequestId:      req.RequestId,
		},
		MessageBody: &uamsg.GenericBody{
			TypeId: &uamsg.ExpandedNodeId{
				NodeId: &uamsg.ObjectCreateSessionResponse_Encoding_DefaultBinary,
			},
			Service: respSvc,
		},
	}

	return rsp, nil
}

func (secChan *SecureChannel) handleActivateSessionRequest(req *uamsg.Message) (*uamsg.Message, error) {
	svc, err := getService[uamsg.ActivateSessionRequest](req)
	if err != nil {
		return nil, err
	}
	err = secChan.interceptor.BeforeActivateSession(secChan.conn, svc)
	if err != nil {
		return nil, err
	}

	token, ok := svc.Header.AuthenticationToken.Identifier.([]byte)
	if !ok {
		return nil, ErrInvalidMessageBody
	}

	session, ok := secChan.sessionManager.get(string(token))
	if !ok {
		return secChan.createErrorMessage(req.RequestId, svc.Header.RequestHandle, uamsg.ErrorCodeBadSessionIdInvalid), nil
	}

	nextServerNonce := util.GenerateRandomBytes(32)
	session.setServerNonce(nextServerNonce)

	respSvc := &uamsg.ActivateSessionResponse{
		Header: &uamsg.ResponseHeader{
			Timestamp:     util.GetCurrentUaTimestamp(),
			RequestHandle: svc.Header.RequestHandle,
			ServiceResult: uint32(uamsg.ErrorCodeGood),
			ServiceDiagnostics: &uamsg.DiagnosticInfo{
				EncodingMask: 0x00,
			},
			StringTable: nil,
			AdditionalHeader: &uamsg.ExtensionObject{
				TypeId: &uamsg.NodeId{
					EncodingType: uamsg.TwoByte,
					Identifier:   byte(0),
				},
				Encoding: 0x00,
			},
		},
		ServerNonce: nextServerNonce,
	}

	err = secChan.interceptor.AfterActivateSession(secChan.conn, svc, respSvc)
	if err != nil {
		return nil, err
	}

	return &uamsg.Message{
		MessageHeader: &uamsg.MessageHeader{
			MessageType:     uamsg.MsgMessageType,
			SecureChannelId: &secChan.channelId,
		},
		SecurityHeader: &uamsg.SymmetricSecurityHeader{
			TokenId: secChan.getCurrentTokenId(),
		},
		SequenceHeader: &uamsg.SequenceHeader{
			SequenceNumber: secChan.getNextSequenceNumber(),
			RequestId:      req.RequestId,
		},
		MessageBody: &uamsg.GenericBody{
			TypeId: &uamsg.ExpandedNodeId{
				NodeId: &uamsg.ObjectActivateSessionResponse_Encoding_DefaultBinary,
			},
			Service: respSvc,
		},
	}, nil
}

func (secChan *SecureChannel) handleCloseSessionRequest(req *uamsg.Message) (*uamsg.Message, error) {
	svc, err := getService[uamsg.CloseSessionRequest](req)
	if err != nil {
		return nil, err
	}
	err = secChan.interceptor.BeforeCloseSession(secChan.conn, svc)
	if err != nil {
		return nil, err
	}

	// TODO for any outstanding request, should return with BadSessionClosed error code
	// TODO SessionDiagnosticsArray

	// TODO handle svc.DeleteSubscriptions

	token, ok := svc.Header.AuthenticationToken.Identifier.([]byte)
	if !ok {
		return nil, ErrInvalidMessageBody
	}
	secChan.sessionManager.delete(string(token))

	respSvc := &uamsg.CloseSessionResponse{
		Header: &uamsg.ResponseHeader{
			Timestamp:     util.GetCurrentUaTimestamp(),
			RequestHandle: svc.Header.RequestHandle,
			ServiceResult: uint32(uamsg.ErrorCodeGood),
			ServiceDiagnostics: &uamsg.DiagnosticInfo{
				EncodingMask: 0x00,
			},
			StringTable: nil,
			AdditionalHeader: &uamsg.ExtensionObject{
				TypeId: &uamsg.NodeId{
					EncodingType: uamsg.TwoByte,
					Identifier:   byte(0),
				},
				Encoding: 0x00,
			},
		},
	}
	err = secChan.interceptor.AfterCloseSession(secChan.conn, svc, respSvc)
	if err != nil {
		return nil, err
	}

	return &uamsg.Message{
		MessageHeader: &uamsg.MessageHeader{
			MessageType:     uamsg.MsgMessageType,
			SecureChannelId: &secChan.channelId,
		},
		SecurityHeader: &uamsg.SymmetricSecurityHeader{
			TokenId: secChan.getCurrentTokenId(),
		},
		SequenceHeader: &uamsg.SequenceHeader{
			SequenceNumber: secChan.getNextSequenceNumber(),
			RequestId:      req.RequestId,
		},
		MessageBody: &uamsg.GenericBody{
			TypeId: &uamsg.ExpandedNodeId{
				NodeId: &uamsg.ObjectCloseSessionResponse_Encoding_DefaultBinary,
			},
			Service: respSvc,
		},
	}, nil
}

func (secChan *SecureChannel) handleGetEndpoints(req *uamsg.Message) (*uamsg.Message, error) {
	svc, err := getService[uamsg.GetEndpointsRequest](req)
	if err != nil {
		return nil, err
	}
	err = secChan.interceptor.BeforeGetEndpoints(secChan.conn, svc)
	if err != nil {
		return nil, err
	}
	respSvc := &uamsg.GetEndpointsResponse{
		Header: &uamsg.ResponseHeader{
			Timestamp:     util.GetCurrentUaTimestamp(),
			RequestHandle: svc.Header.RequestHandle,
			ServiceResult: uint32(uamsg.ErrorCodeGood),
			ServiceDiagnostics: &uamsg.DiagnosticInfo{
				EncodingMask: 0x00,
			},
			StringTable: nil,
			AdditionalHeader: &uamsg.ExtensionObject{
				TypeId: &uamsg.NodeId{
					EncodingType: uamsg.TwoByte,
					Identifier:   byte(0),
				},
				Encoding: 0x00,
			},
		},
		// TODO can support build endpoint description from callback
		Endpoints: []uamsg.EndpointDescription{
			{
				EndpointUrl: svc.EndpointUrl,
				Server: &uamsg.ApplicationDescription{
					ApplicationName: &uamsg.LocalizedText{},
					ApplicationType: 0,
				},
				SecurityMode:      uamsg.MessageSecurityModeNone,
				SecurityPolicyUri: uamsg.SecurityPolicyUriNone,
				UserIdentityTokens: []*uamsg.UserTokenPolicy{
					{
						PolicyId:  "username",
						TokenType: uamsg.UserTokenTypeUsername,
					},
				},
				TransportProfileUri: "http://opcfoundation.org/UA-Profile/Transport/uatcp-uasc-uabinary",
			},
		},
	}
	err = secChan.interceptor.AfterGetEndpoints(secChan.conn, svc, respSvc)
	if err != nil {
		return nil, err
	}

	return &uamsg.Message{
		MessageHeader: &uamsg.MessageHeader{
			MessageType:     uamsg.MsgMessageType,
			SecureChannelId: &secChan.channelId,
		},
		SecurityHeader: &uamsg.SymmetricSecurityHeader{
			TokenId: secChan.getCurrentTokenId(),
		},
		SequenceHeader: &uamsg.SequenceHeader{
			SequenceNumber: secChan.getNextSequenceNumber(),
			RequestId:      req.RequestId,
		},
		MessageBody: &uamsg.GenericBody{
			TypeId: &uamsg.ExpandedNodeId{
				NodeId: &uamsg.ObjectGetEndpointsResponse_Encoding_DefaultBinary,
			},
			Service: respSvc,
		},
	}, nil
}

func (secChan *SecureChannel) handleBrowseRequest(req *uamsg.Message) (*uamsg.Message, error) {
	svc, err := getService[uamsg.BrowseRequest](req)
	if err != nil {
		return nil, err
	}
	err = secChan.interceptor.BeforeBrowse(secChan.conn, svc)
	if err != nil {
		return nil, err
	}
	err = secChan.interceptor.AfterBrowse(secChan.conn, svc, nil)
	if err != nil {
		return nil, err
	}
	return nil, nil
}

func (secChan *SecureChannel) handleReadRequest(req *uamsg.Message) (*uamsg.Message, error) {
	svc, err := getService[uamsg.ReadRequest](req)
	if err != nil {
		return nil, err
	}
	err = secChan.interceptor.BeforeRead(secChan.conn, svc)
	if err != nil {
		return nil, err
	}
	respSvc, err := secChan.handler.HandleRead(secChan.conn, svc)
	if err != nil {
		return nil, err
	}
	err = secChan.interceptor.AfterRead(secChan.conn, svc, respSvc)
	if err != nil {
		return nil, err
	}
	return &uamsg.Message{
		MessageHeader: &uamsg.MessageHeader{
			MessageType:     uamsg.MsgMessageType,
			SecureChannelId: &secChan.channelId,
		},
		SecurityHeader: &uamsg.SymmetricSecurityHeader{
			TokenId: secChan.getCurrentTokenId(),
		},
		SequenceHeader: &uamsg.SequenceHeader{
			SequenceNumber: secChan.getNextSequenceNumber(),
			RequestId:      req.RequestId,
		},
		MessageBody: &uamsg.GenericBody{
			TypeId: &uamsg.ExpandedNodeId{
				NodeId: &uamsg.ObjectReadResponse_Encoding_DefaultBinary,
			},
			Service: respSvc,
		},
		MessageFooter: nil,
	}, nil
}

func getService[T any](msg *uamsg.Message) (*T, error) {
	generic, ok := msg.MessageBody.(uamsg.GenericBody)
	if !ok {
		return nil, ErrInvalidMessageBody
	}

	service, ok := generic.Service.(*T)
	if !ok {
		return nil, ErrInvalidMessageBody
	}

	return service, nil
}
