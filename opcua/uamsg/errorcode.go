package uamsg

type ErrorCode uint32

const (
	ErrorCodeGood                                                            ErrorCode = 0x00000000
	ErrorCodeUncertain                                                       ErrorCode = 0x40000000
	ErrorCodeBad                                                             ErrorCode = 0x80000000
	ErrorCodeBadUnexpectedError                                              ErrorCode = 0x80010000
	ErrorCodeBadInternalError                                                ErrorCode = 0x80020000
	ErrorCodeBadOutOfMemory                                                  ErrorCode = 0x80030000
	ErrorCodeBadResourceUnavailable                                          ErrorCode = 0x80040000
	ErrorCodeBadCommunicationError                                           ErrorCode = 0x80050000
	ErrorCodeBadEncodingError                                                ErrorCode = 0x80060000
	ErrorCodeBadDecodingError                                                ErrorCode = 0x80070000
	ErrorCodeBadEncodingLimitsExceeded                                       ErrorCode = 0x80080000
	ErrorCodeBadRequestTooLarge                                              ErrorCode = 0x80B80000
	ErrorCodeBadResponseTooLarge                                             ErrorCode = 0x80B90000
	ErrorCodeBadUnknownResponse                                              ErrorCode = 0x80090000
	ErrorCodeBadTimeout                                                      ErrorCode = 0x800A0000
	ErrorCodeBadServiceUnsupported                                           ErrorCode = 0x800B0000
	ErrorCodeBadShutdown                                                     ErrorCode = 0x800C0000
	ErrorCodeBadServerNotConnected                                           ErrorCode = 0x800D0000
	ErrorCodeBadServerHalted                                                 ErrorCode = 0x800E0000
	ErrorCodeBadNothingToDo                                                  ErrorCode = 0x800F0000
	ErrorCodeBadTooManyOperations                                            ErrorCode = 0x80100000
	ErrorCodeBadTooManyMonitoredItems                                        ErrorCode = 0x80DB0000
	ErrorCodeBadDataTypeIdUnknown                                            ErrorCode = 0x80110000
	ErrorCodeBadCertificateInvalid                                           ErrorCode = 0x80120000
	ErrorCodeBadSecurityChecksFailed                                         ErrorCode = 0x80130000
	ErrorCodeBadCertificatePolicyCheckFailed                                 ErrorCode = 0x81140000
	ErrorCodeBadCertificateTimeInvalid                                       ErrorCode = 0x80140000
	ErrorCodeBadCertificateIssuerTimeInvalid                                 ErrorCode = 0x80150000
	ErrorCodeBadCertificateHostNameInvalid                                   ErrorCode = 0x80160000
	ErrorCodeBadCertificateUriInvalid                                        ErrorCode = 0x80170000
	ErrorCodeBadCertificateUseNotAllowed                                     ErrorCode = 0x80180000
	ErrorCodeBadCertificateIssuerUseNotAllowed                               ErrorCode = 0x80190000
	ErrorCodeBadCertificateUntrusted                                         ErrorCode = 0x801A0000
	ErrorCodeBadCertificateRevocationUnknown                                 ErrorCode = 0x801B0000
	ErrorCodeBadCertificateIssuerRevocationUnknown                           ErrorCode = 0x801C0000
	ErrorCodeBadCertificateRevoked                                           ErrorCode = 0x801D0000
	ErrorCodeBadCertificateIssuerRevoked                                     ErrorCode = 0x801E0000
	ErrorCodeBadCertificateChainIncomplete                                   ErrorCode = 0x810D0000
	ErrorCodeBadUserAccessDenied                                             ErrorCode = 0x801F0000
	ErrorCodeBadIdentityTokenInvalid                                         ErrorCode = 0x80200000
	ErrorCodeBadIdentityTokenRejected                                        ErrorCode = 0x80210000
	ErrorCodeBadSecureChannelIdInvalid                                       ErrorCode = 0x80220000
	ErrorCodeBadInvalidTimestamp                                             ErrorCode = 0x80230000
	ErrorCodeBadNonceInvalid                                                 ErrorCode = 0x80240000
	ErrorCodeBadSessionIdInvalid                                             ErrorCode = 0x80250000
	ErrorCodeBadSessionClosed                                                ErrorCode = 0x80260000
	ErrorCodeBadSessionNotActivated                                          ErrorCode = 0x80270000
	ErrorCodeBadSubscriptionIdInvalid                                        ErrorCode = 0x80280000
	ErrorCodeBadRequestHeaderInvalid                                         ErrorCode = 0x802A0000
	ErrorCodeBadTimestampsToReturnInvalid                                    ErrorCode = 0x802B0000
	ErrorCodeBadRequestCancelledByClient                                     ErrorCode = 0x802C0000
	ErrorCodeBadTooManyArguments                                             ErrorCode = 0x80E50000
	ErrorCodeBadLicenseExpired                                               ErrorCode = 0x810E0000
	ErrorCodeBadLicenseLimitsExceeded                                        ErrorCode = 0x810F0000
	ErrorCodeBadLicenseNotAvailable                                          ErrorCode = 0x81100000
	ErrorCodeBadServerTooBusy                                                ErrorCode = 0x80EE0000
	ErrorCodeGoodPasswordChangeRequired                                      ErrorCode = 0x00EF0000
	ErrorCodeGoodSubscriptionTransferred                                     ErrorCode = 0x002D0000
	ErrorCodeGoodCompletesAsynchronously                                     ErrorCode = 0x002E0000
	ErrorCodeGoodOverload                                                    ErrorCode = 0x002F0000
	ErrorCodeGoodClamped                                                     ErrorCode = 0x00300000
	ErrorCodeBadNoCommunication                                              ErrorCode = 0x80310000
	ErrorCodeBadWaitingForInitialData                                        ErrorCode = 0x80320000
	ErrorCodeBadNodeIdInvalid                                                ErrorCode = 0x80330000
	ErrorCodeBadNodeIdUnknown                                                ErrorCode = 0x80340000
	ErrorCodeBadAttributeIdInvalid                                           ErrorCode = 0x80350000
	ErrorCodeBadIndexRangeInvalid                                            ErrorCode = 0x80360000
	ErrorCodeBadIndexRangeNoData                                             ErrorCode = 0x80370000
	ErrorCodeBadIndexRangeDataMismatch                                       ErrorCode = 0x80EA0000
	ErrorCodeBadDataEncodingInvalid                                          ErrorCode = 0x80380000
	ErrorCodeBadDataEncodingUnsupported                                      ErrorCode = 0x80390000
	ErrorCodeBadNotReadable                                                  ErrorCode = 0x803A0000
	ErrorCodeBadNotWritable                                                  ErrorCode = 0x803B0000
	ErrorCodeBadOutOfRange                                                   ErrorCode = 0x803C0000
	ErrorCodeBadNotSupported                                                 ErrorCode = 0x803D0000
	ErrorCodeBadNotFound                                                     ErrorCode = 0x803E0000
	ErrorCodeBadObjectDeleted                                                ErrorCode = 0x803F0000
	ErrorCodeBadNotImplemented                                               ErrorCode = 0x80400000
	ErrorCodeBadMonitoringModeInvalid                                        ErrorCode = 0x80410000
	ErrorCodeBadMonitoredItemIdInvalid                                       ErrorCode = 0x80420000
	ErrorCodeBadMonitoredItemFilterInvalid                                   ErrorCode = 0x80430000
	ErrorCodeBadMonitoredItemFilterUnsupported                               ErrorCode = 0x80440000
	ErrorCodeBadFilterNotAllowed                                             ErrorCode = 0x80450000
	ErrorCodeBadStructureMissing                                             ErrorCode = 0x80460000
	ErrorCodeBadEventFilterInvalid                                           ErrorCode = 0x80470000
	ErrorCodeBadContentFilterInvalid                                         ErrorCode = 0x80480000
	ErrorCodeBadFilterOperatorInvalid                                        ErrorCode = 0x80C10000
	ErrorCodeBadFilterOperatorUnsupported                                    ErrorCode = 0x80C20000
	ErrorCodeBadFilterOperandCountMismatch                                   ErrorCode = 0x80C30000
	ErrorCodeBadFilterOperandInvalid                                         ErrorCode = 0x80490000
	ErrorCodeBadFilterElementInvalid                                         ErrorCode = 0x80C40000
	ErrorCodeBadFilterLiteralInvalid                                         ErrorCode = 0x80C50000
	ErrorCodeBadContinuationPointInvalid                                     ErrorCode = 0x804A0000
	ErrorCodeBadNoContinuationPoints                                         ErrorCode = 0x804B0000
	ErrorCodeBadReferenceTypeIdInvalid                                       ErrorCode = 0x804C0000
	ErrorCodeBadBrowseDirectionInvalid                                       ErrorCode = 0x804D0000
	ErrorCodeBadNodeNotInView                                                ErrorCode = 0x804E0000
	ErrorCodeBadNumericOverflow                                              ErrorCode = 0x81120000
	ErrorCodeBadLocaleNotSupported                                           ErrorCode = 0x80ED0000
	ErrorCodeBadNoValue                                                      ErrorCode = 0x80F00000
	ErrorCodeBadServerUriInvalid                                             ErrorCode = 0x804F0000
	ErrorCodeBadServerNameMissing                                            ErrorCode = 0x80500000
	ErrorCodeBadDiscoveryUrlMissing                                          ErrorCode = 0x80510000
	ErrorCodeBadSempahoreFileMissing                                         ErrorCode = 0x80520000
	ErrorCodeBadRequestTypeInvalid                                           ErrorCode = 0x80530000
	ErrorCodeBadSecurityModeRejected                                         ErrorCode = 0x80540000
	ErrorCodeBadSecurityPolicyRejected                                       ErrorCode = 0x80550000
	ErrorCodeBadTooManySessions                                              ErrorCode = 0x80560000
	ErrorCodeBadUserSignatureInvalid                                         ErrorCode = 0x80570000
	ErrorCodeBadApplicationSignatureInvalid                                  ErrorCode = 0x80580000
	ErrorCodeBadNoValidCertificates                                          ErrorCode = 0x80590000
	ErrorCodeBadIdentityChangeNotSupported                                   ErrorCode = 0x80C60000
	ErrorCodeBadRequestCancelledByRequest                                    ErrorCode = 0x805A0000
	ErrorCodeBadParentNodeIdInvalid                                          ErrorCode = 0x805B0000
	ErrorCodeBadReferenceNotAllowed                                          ErrorCode = 0x805C0000
	ErrorCodeBadNodeIdRejected                                               ErrorCode = 0x805D0000
	ErrorCodeBadNodeIdExists                                                 ErrorCode = 0x805E0000
	ErrorCodeBadNodeClassInvalid                                             ErrorCode = 0x805F0000
	ErrorCodeBadBrowseNameInvalid                                            ErrorCode = 0x80600000
	ErrorCodeBadBrowseNameDuplicated                                         ErrorCode = 0x80610000
	ErrorCodeBadNodeAttributesInvalid                                        ErrorCode = 0x80620000
	ErrorCodeBadTypeDefinitionInvalid                                        ErrorCode = 0x80630000
	ErrorCodeBadSourceNodeIdInvalid                                          ErrorCode = 0x80640000
	ErrorCodeBadTargetNodeIdInvalid                                          ErrorCode = 0x80650000
	ErrorCodeBadDuplicateReferenceNotAllowed                                 ErrorCode = 0x80660000
	ErrorCodeBadInvalidSelfReference                                         ErrorCode = 0x80670000
	ErrorCodeBadReferenceLocalOnly                                           ErrorCode = 0x80680000
	ErrorCodeBadNoDeleteRights                                               ErrorCode = 0x80690000
	ErrorCodeUncertainReferenceNotDeleted                                    ErrorCode = 0x40BC0000
	ErrorCodeBadServerIndexInvalid                                           ErrorCode = 0x806A0000
	ErrorCodeBadViewIdUnknown                                                ErrorCode = 0x806B0000
	ErrorCodeBadViewTimestampInvalid                                         ErrorCode = 0x80C90000
	ErrorCodeBadViewParameterMismatch                                        ErrorCode = 0x80CA0000
	ErrorCodeBadViewVersionInvalid                                           ErrorCode = 0x80CB0000
	ErrorCodeUncertainNotAllNodesAvailable                                   ErrorCode = 0x40C00000
	ErrorCodeGoodResultsMayBeIncomplete                                      ErrorCode = 0x00BA0000
	ErrorCodeBadNotTypeDefinition                                            ErrorCode = 0x80C80000
	ErrorCodeUncertainReferenceOutOfServer                                   ErrorCode = 0x406C0000
	ErrorCodeBadTooManyMatches                                               ErrorCode = 0x806D0000
	ErrorCodeBadQueryTooComplex                                              ErrorCode = 0x806E0000
	ErrorCodeBadNoMatch                                                      ErrorCode = 0x806F0000
	ErrorCodeBadMaxAgeInvalid                                                ErrorCode = 0x80700000
	ErrorCodeBadSecurityModeInsufficient                                     ErrorCode = 0x80E60000
	ErrorCodeBadHistoryOperationInvalid                                      ErrorCode = 0x80710000
	ErrorCodeBadHistoryOperationUnsupported                                  ErrorCode = 0x80720000
	ErrorCodeBadInvalidTimestampArgument                                     ErrorCode = 0x80BD0000
	ErrorCodeBadWriteNotSupported                                            ErrorCode = 0x80730000
	ErrorCodeBadTypeMismatch                                                 ErrorCode = 0x80740000
	ErrorCodeBadMethodInvalid                                                ErrorCode = 0x80750000
	ErrorCodeBadArgumentsMissing                                             ErrorCode = 0x80760000
	ErrorCodeBadNotExecutable                                                ErrorCode = 0x81110000
	ErrorCodeBadTooManySubscriptions                                         ErrorCode = 0x80770000
	ErrorCodeBadTooManyPublishRequests                                       ErrorCode = 0x80780000
	ErrorCodeBadNoSubscription                                               ErrorCode = 0x80790000
	ErrorCodeBadSequenceNumberUnknown                                        ErrorCode = 0x807A0000
	ErrorCodeGoodRetransmissionQueueNotSupported                             ErrorCode = 0x00DF0000
	ErrorCodeBadMessageNotAvailable                                          ErrorCode = 0x807B0000
	ErrorCodeBadInsufficientClientProfile                                    ErrorCode = 0x807C0000
	ErrorCodeBadStateNotActive                                               ErrorCode = 0x80BF0000
	ErrorCodeBadAlreadyExists                                                ErrorCode = 0x81150000
	ErrorCodeBadTcpServerTooBusy                                             ErrorCode = 0x807D0000
	ErrorCodeBadTcpMessageTypeInvalid                                        ErrorCode = 0x807E0000
	ErrorCodeBadTcpSecureChannelUnknown                                      ErrorCode = 0x807F0000
	ErrorCodeBadTcpMessageTooLarge                                           ErrorCode = 0x80800000
	ErrorCodeBadTcpNotEnoughResources                                        ErrorCode = 0x80810000
	ErrorCodeBadTcpInternalError                                             ErrorCode = 0x80820000
	ErrorCodeBadTcpEndpointUrlInvalid                                        ErrorCode = 0x80830000
	ErrorCodeBadRequestInterrupted                                           ErrorCode = 0x80840000
	ErrorCodeBadRequestTimeout                                               ErrorCode = 0x80850000
	ErrorCodeBadSecureChannelClosed                                          ErrorCode = 0x80860000
	ErrorCodeBadSecureChannelTokenUnknown                                    ErrorCode = 0x80870000
	ErrorCodeBadSequenceNumberInvalid                                        ErrorCode = 0x80880000
	ErrorCodeBadProtocolVersionUnsupported                                   ErrorCode = 0x80BE0000
	ErrorCodeBadConfigurationError                                           ErrorCode = 0x80890000
	ErrorCodeBadNotConnected                                                 ErrorCode = 0x808A0000
	ErrorCodeBadDeviceFailure                                                ErrorCode = 0x808B0000
	ErrorCodeBadSensorFailure                                                ErrorCode = 0x808C0000
	ErrorCodeBadOutOfService                                                 ErrorCode = 0x808D0000
	ErrorCodeBadDeadbandFilterInvalid                                        ErrorCode = 0x808E0000
	ErrorCodeUncertainNoCommunicationLastUsableValue                         ErrorCode = 0x408F0000
	ErrorCodeUncertainLastUsableValue                                        ErrorCode = 0x40900000
	ErrorCodeUncertainSubstituteValue                                        ErrorCode = 0x40910000
	ErrorCodeUncertainInitialValue                                           ErrorCode = 0x40920000
	ErrorCodeUncertainSensorNotAccurate                                      ErrorCode = 0x40930000
	ErrorCodeUncertainEngineeringUnitsExceeded                               ErrorCode = 0x40940000
	ErrorCodeUncertainSubNormal                                              ErrorCode = 0x40950000
	ErrorCodeGoodLocalOverride                                               ErrorCode = 0x00960000
	ErrorCodeGoodSubNormal                                                   ErrorCode = 0x00EB0000
	ErrorCodeBadRefreshInProgress                                            ErrorCode = 0x80970000
	ErrorCodeBadConditionAlreadyDisabled                                     ErrorCode = 0x80980000
	ErrorCodeBadConditionAlreadyEnabled                                      ErrorCode = 0x80CC0000
	ErrorCodeBadConditionDisabled                                            ErrorCode = 0x80990000
	ErrorCodeBadEventIdUnknown                                               ErrorCode = 0x809A0000
	ErrorCodeBadEventNotAcknowledgeable                                      ErrorCode = 0x80BB0000
	ErrorCodeBadDialogNotActive                                              ErrorCode = 0x80CD0000
	ErrorCodeBadDialogResponseInvalid                                        ErrorCode = 0x80CE0000
	ErrorCodeBadConditionBranchAlreadyAcked                                  ErrorCode = 0x80CF0000
	ErrorCodeBadConditionBranchAlreadyConfirmed                              ErrorCode = 0x80D00000
	ErrorCodeBadConditionAlreadyShelved                                      ErrorCode = 0x80D10000
	ErrorCodeBadConditionNotShelved                                          ErrorCode = 0x80D20000
	ErrorCodeBadShelvingTimeOutOfRange                                       ErrorCode = 0x80D30000
	ErrorCodeBadNoData                                                       ErrorCode = 0x809B0000
	ErrorCodeBadBoundNotFound                                                ErrorCode = 0x80D70000
	ErrorCodeBadBoundNotSupported                                            ErrorCode = 0x80D80000
	ErrorCodeBadDataLost                                                     ErrorCode = 0x809D0000
	ErrorCodeBadDataUnavailable                                              ErrorCode = 0x809E0000
	ErrorCodeBadEntryExists                                                  ErrorCode = 0x809F0000
	ErrorCodeBadNoEntryExists                                                ErrorCode = 0x80A00000
	ErrorCodeBadTimestampNotSupported                                        ErrorCode = 0x80A10000
	ErrorCodeGoodEntryInserted                                               ErrorCode = 0x00A20000
	ErrorCodeGoodEntryReplaced                                               ErrorCode = 0x00A30000
	ErrorCodeUncertainDataSubNormal                                          ErrorCode = 0x40A40000
	ErrorCodeGoodNoData                                                      ErrorCode = 0x00A50000
	ErrorCodeGoodMoreData                                                    ErrorCode = 0x00A60000
	ErrorCodeBadAggregateListMismatch                                        ErrorCode = 0x80D40000
	ErrorCodeBadAggregateNotSupported                                        ErrorCode = 0x80D50000
	ErrorCodeBadAggregateInvalidInputs                                       ErrorCode = 0x80D60000
	ErrorCodeBadAggregateConfigurationRejected                               ErrorCode = 0x80DA0000
	ErrorCodeGoodDataIgnored                                                 ErrorCode = 0x00D90000
	ErrorCodeBadRequestNotAllowed                                            ErrorCode = 0x80E40000
	ErrorCodeBadRequestNotComplete                                           ErrorCode = 0x81130000
	ErrorCodeBadTransactionPending                                           ErrorCode = 0x80E80000
	ErrorCodeBadTicketRequired                                               ErrorCode = 0x811F0000
	ErrorCodeBadTicketInvalid                                                ErrorCode = 0x81200000
	ErrorCodeBadLocked                                                       ErrorCode = 0x80E90000
	ErrorCodeBadRequiresLock                                                 ErrorCode = 0x80EC0000
	ErrorCodeGoodEdited                                                      ErrorCode = 0x00DC0000
	ErrorCodeGoodPostActionFailed                                            ErrorCode = 0x00DD0000
	ErrorCodeUncertainDominantValueChanged                                   ErrorCode = 0x40DE0000
	ErrorCodeGoodDependentValueChanged                                       ErrorCode = 0x00E00000
	ErrorCodeBadDominantValueChanged                                         ErrorCode = 0x80E10000
	ErrorCodeUncertainDependentValueChanged                                  ErrorCode = 0x40E20000
	ErrorCodeBadDependentValueChanged                                        ErrorCode = 0x80E30000
	ErrorCodeGoodEdited_DependentValueChanged                                ErrorCode = 0x01160000
	ErrorCodeGoodEdited_DominantValueChanged                                 ErrorCode = 0x01170000
	ErrorCodeGoodEdited_DominantValueChanged_DependentValueChanged           ErrorCode = 0x01180000
	ErrorCodeBadEdited_OutOfRange                                            ErrorCode = 0x81190000
	ErrorCodeBadInitialValue_OutOfRange                                      ErrorCode = 0x811A0000
	ErrorCodeBadOutOfRange_DominantValueChanged                              ErrorCode = 0x811B0000
	ErrorCodeBadEdited_OutOfRange_DominantValueChanged                       ErrorCode = 0x811C0000
	ErrorCodeBadOutOfRange_DominantValueChanged_DependentValueChanged        ErrorCode = 0x811D0000
	ErrorCodeBadEdited_OutOfRange_DominantValueChanged_DependentValueChanged ErrorCode = 0x811E0000
	ErrorCodeGoodCommunicationEvent                                          ErrorCode = 0x00A70000
	ErrorCodeGoodShutdownEvent                                               ErrorCode = 0x00A80000
	ErrorCodeGoodCallAgain                                                   ErrorCode = 0x00A90000
	ErrorCodeGoodNonCriticalTimeout                                          ErrorCode = 0x00AA0000
	ErrorCodeBadInvalidArgument                                              ErrorCode = 0x80AB0000
	ErrorCodeBadConnectionRejected                                           ErrorCode = 0x80AC0000
	ErrorCodeBadDisconnect                                                   ErrorCode = 0x80AD0000
	ErrorCodeBadConnectionClosed                                             ErrorCode = 0x80AE0000
	ErrorCodeBadInvalidState                                                 ErrorCode = 0x80AF0000
	ErrorCodeBadEndOfStream                                                  ErrorCode = 0x80B00000
	ErrorCodeBadNoDataAvailable                                              ErrorCode = 0x80B10000
	ErrorCodeBadWaitingForResponse                                           ErrorCode = 0x80B20000
	ErrorCodeBadOperationAbandoned                                           ErrorCode = 0x80B30000
	ErrorCodeBadExpectedStreamToBlock                                        ErrorCode = 0x80B40000
	ErrorCodeBadWouldBlock                                                   ErrorCode = 0x80B50000
	ErrorCodeBadSyntaxError                                                  ErrorCode = 0x80B60000
	ErrorCodeBadMaxConnectionsReached                                        ErrorCode = 0x80B70000
	ErrorCodeUncertainTransducerInManual                                     ErrorCode = 0x42080000
	ErrorCodeUncertainSimulatedValue                                         ErrorCode = 0x42090000
	ErrorCodeUncertainSensorCalibration                                      ErrorCode = 0x420A0000
	ErrorCodeUncertainConfigurationError                                     ErrorCode = 0x420F0000
	ErrorCodeGoodCascadeInitializationAcknowledged                           ErrorCode = 0x04010000
	ErrorCodeGoodCascadeInitializationRequest                                ErrorCode = 0x04020000
	ErrorCodeGoodCascadeNotInvited                                           ErrorCode = 0x04030000
	ErrorCodeGoodCascadeNotSelected                                          ErrorCode = 0x04040000
	ErrorCodeGoodFaultStateActive                                            ErrorCode = 0x04070000
	ErrorCodeGoodInitiateFaultState                                          ErrorCode = 0x04080000
	ErrorCodeGoodCascade                                                     ErrorCode = 0x04090000
	ErrorCodeBadDataSetIdInvalid                                             ErrorCode = 0x80E70000
)

var ErrorCodes = map[ErrorCode]string{
	ErrorCodeGood:                                                            "The operation succeeded.",
	ErrorCodeUncertain:                                                       "The operation was uncertain.",
	ErrorCodeBad:                                                             "The operation failed.",
	ErrorCodeBadUnexpectedError:                                              "An unexpected error occurred.",
	ErrorCodeBadInternalError:                                                "An internal error occurred as a result of a programming or configuration error.",
	ErrorCodeBadOutOfMemory:                                                  "Not enough memory to complete the operation.",
	ErrorCodeBadResourceUnavailable:                                          "An operating system resource is not available.",
	ErrorCodeBadCommunicationError:                                           "A low level communication error occurred.",
	ErrorCodeBadEncodingError:                                                "Encoding halted because of invalid data in the objects being serialized.",
	ErrorCodeBadDecodingError:                                                "Decoding halted because of invalid data in the stream.",
	ErrorCodeBadEncodingLimitsExceeded:                                       "The message encoding/decoding limits imposed by the stack have been exceeded.",
	ErrorCodeBadRequestTooLarge:                                              "The request message size exceeds limits set by the server.",
	ErrorCodeBadResponseTooLarge:                                             "The response message size exceeds limits set by the client or server.",
	ErrorCodeBadUnknownResponse:                                              "An unrecognized response was received from the server.",
	ErrorCodeBadTimeout:                                                      "The operation timed out.",
	ErrorCodeBadServiceUnsupported:                                           "The server does not support the requested service.",
	ErrorCodeBadShutdown:                                                     "The operation was cancelled because the application is shutting down.",
	ErrorCodeBadServerNotConnected:                                           "The operation could not complete because the client is not connected to the server.",
	ErrorCodeBadServerHalted:                                                 "The server has stopped and cannot process any requests.",
	ErrorCodeBadNothingToDo:                                                  "No processing could be done because there was nothing to do.",
	ErrorCodeBadTooManyOperations:                                            "The request could not be processed because it specified too many operations.",
	ErrorCodeBadTooManyMonitoredItems:                                        "The request could not be processed because there are too many monitored items in the subscription.",
	ErrorCodeBadDataTypeIdUnknown:                                            "The extension object cannot be (de)serialized because the data type id is not recognized.",
	ErrorCodeBadCertificateInvalid:                                           "The certificate provided as a parameter is not valid.",
	ErrorCodeBadSecurityChecksFailed:                                         "An error occurred verifying security.",
	ErrorCodeBadCertificatePolicyCheckFailed:                                 "The certificate does not meet the requirements of the security policy.",
	ErrorCodeBadCertificateTimeInvalid:                                       "The certificate has expired or is not yet valid.",
	ErrorCodeBadCertificateIssuerTimeInvalid:                                 "An issuer certificate has expired or is not yet valid.",
	ErrorCodeBadCertificateHostNameInvalid:                                   "The HostName used to connect to a server does not match a HostName in the certificate.",
	ErrorCodeBadCertificateUriInvalid:                                        "The URI specified in the ApplicationDescription does not match the URI in the certificate.",
	ErrorCodeBadCertificateUseNotAllowed:                                     "The certificate may not be used for the requested operation.",
	ErrorCodeBadCertificateIssuerUseNotAllowed:                               "The issuer certificate may not be used for the requested operation.",
	ErrorCodeBadCertificateUntrusted:                                         "The certificate is not trusted.",
	ErrorCodeBadCertificateRevocationUnknown:                                 "It was not possible to determine if the certificate has been revoked.",
	ErrorCodeBadCertificateIssuerRevocationUnknown:                           "It was not possible to determine if the issuer certificate has been revoked.",
	ErrorCodeBadCertificateRevoked:                                           "The certificate has been revoked.",
	ErrorCodeBadCertificateIssuerRevoked:                                     "The issuer certificate has been revoked.",
	ErrorCodeBadCertificateChainIncomplete:                                   "The certificate chain is incomplete.",
	ErrorCodeBadUserAccessDenied:                                             "User does not have permission to perform the requested operation.",
	ErrorCodeBadIdentityTokenInvalid:                                         "The user identity token is not valid.",
	ErrorCodeBadIdentityTokenRejected:                                        "The user identity token is valid but the server has rejected it.",
	ErrorCodeBadSecureChannelIdInvalid:                                       "The specified secure channel is no longer valid.",
	ErrorCodeBadInvalidTimestamp:                                             "The timestamp is outside the range allowed by the server.",
	ErrorCodeBadNonceInvalid:                                                 "The nonce does appear to be not a random value or it is not the correct length.",
	ErrorCodeBadSessionIdInvalid:                                             "The session id is not valid.",
	ErrorCodeBadSessionClosed:                                                "The session was closed by the client.",
	ErrorCodeBadSessionNotActivated:                                          "The session cannot be used because ActivateSession has not been called.",
	ErrorCodeBadSubscriptionIdInvalid:                                        "The subscription id is not valid.",
	ErrorCodeBadRequestHeaderInvalid:                                         "The header for the request is missing or invalid.",
	ErrorCodeBadTimestampsToReturnInvalid:                                    "The timestamps to return parameter is invalid.",
	ErrorCodeBadRequestCancelledByClient:                                     "The request was cancelled by the client.",
	ErrorCodeBadTooManyArguments:                                             "Too many arguments were provided.",
	ErrorCodeBadLicenseExpired:                                               "The server requires a license to operate in general or to perform a service or operation, but existing license is expired.",
	ErrorCodeBadLicenseLimitsExceeded:                                        "The server has limits on number of allowed operations / objects, based on installed licenses, and these limits where exceeded.",
	ErrorCodeBadLicenseNotAvailable:                                          "The server does not have a license which is required to operate in general or to perform a service or operation.",
	ErrorCodeBadServerTooBusy:                                                "The Server does not have the resources to process the request at this time.",
	ErrorCodeGoodPasswordChangeRequired:                                      "The log-on for the user succeeded but the user is required to change the password.",
	ErrorCodeGoodSubscriptionTransferred:                                     "The subscription was transferred to another session.",
	ErrorCodeGoodCompletesAsynchronously:                                     "The processing will complete asynchronously.",
	ErrorCodeGoodOverload:                                                    "Sampling has slowed down due to resource limitations.",
	ErrorCodeGoodClamped:                                                     "The value written was accepted but was clamped.",
	ErrorCodeBadNoCommunication:                                              "Communication with the data source is defined, but not established, and there is no last known value available.",
	ErrorCodeBadWaitingForInitialData:                                        "Waiting for the server to obtain values from the underlying data source.",
	ErrorCodeBadNodeIdInvalid:                                                "The syntax the node id is not valid or refers to a node that is not valid for the operation.",
	ErrorCodeBadNodeIdUnknown:                                                "The node id refers to a node that does not exist in the server address space.",
	ErrorCodeBadAttributeIdInvalid:                                           "The attribute is not supported for the specified Node.",
	ErrorCodeBadIndexRangeInvalid:                                            "The syntax of the index range parameter is invalid.",
	ErrorCodeBadIndexRangeNoData:                                             "No data exists within the range of indexes specified.",
	ErrorCodeBadIndexRangeDataMismatch:                                       "The written data does not match the IndexRange specified.",
	ErrorCodeBadDataEncodingInvalid:                                          "The data encoding is invalid.",
	ErrorCodeBadDataEncodingUnsupported:                                      "The server does not support the requested data encoding for the node.",
	ErrorCodeBadNotReadable:                                                  "The access level does not allow reading or subscribing to the Node.",
	ErrorCodeBadNotWritable:                                                  "The access level does not allow writing to the Node.",
	ErrorCodeBadOutOfRange:                                                   "The value was out of range.",
	ErrorCodeBadNotSupported:                                                 "The requested operation is not supported.",
	ErrorCodeBadNotFound:                                                     "A requested item was not found or a search operation ended without success.",
	ErrorCodeBadObjectDeleted:                                                "The object cannot be used because it has been deleted.",
	ErrorCodeBadNotImplemented:                                               "Requested operation is not implemented.",
	ErrorCodeBadMonitoringModeInvalid:                                        "The monitoring mode is invalid.",
	ErrorCodeBadMonitoredItemIdInvalid:                                       "The monitoring item id does not refer to a valid monitored item.",
	ErrorCodeBadMonitoredItemFilterInvalid:                                   "The monitored item filter parameter is not valid.",
	ErrorCodeBadMonitoredItemFilterUnsupported:                               "The server does not support the requested monitored item filter.",
	ErrorCodeBadFilterNotAllowed:                                             "A monitoring filter cannot be used in combination with the attribute specified.",
	ErrorCodeBadStructureMissing:                                             "A mandatory structured parameter was missing or null.",
	ErrorCodeBadEventFilterInvalid:                                           "The event filter is not valid.",
	ErrorCodeBadContentFilterInvalid:                                         "The content filter is not valid.",
	ErrorCodeBadFilterOperatorInvalid:                                        "An unrecognized operator was provided in a filter.",
	ErrorCodeBadFilterOperatorUnsupported:                                    "A valid operator was provided, but the server does not provide support for this filter operator.",
	ErrorCodeBadFilterOperandCountMismatch:                                   "The number of operands provided for the filter operator was less then expected for the operand provided.",
	ErrorCodeBadFilterOperandInvalid:                                         "The operand used in a content filter is not valid.",
	ErrorCodeBadFilterElementInvalid:                                         "The referenced element is not a valid element in the content filter.",
	ErrorCodeBadFilterLiteralInvalid:                                         "The referenced literal is not a valid value.",
	ErrorCodeBadContinuationPointInvalid:                                     "The continuation point provide is longer valid.",
	ErrorCodeBadNoContinuationPoints:                                         "The operation could not be processed because all continuation points have been allocated.",
	ErrorCodeBadReferenceTypeIdInvalid:                                       "The reference type id does not refer to a valid reference type node.",
	ErrorCodeBadBrowseDirectionInvalid:                                       "The browse direction is not valid.",
	ErrorCodeBadNodeNotInView:                                                "The node is not part of the view.",
	ErrorCodeBadNumericOverflow:                                              "The number was not accepted because of a numeric overflow.",
	ErrorCodeBadLocaleNotSupported:                                           "The locale in the requested write operation is not supported.",
	ErrorCodeBadNoValue:                                                      "The variable has no default value and no initial value.",
	ErrorCodeBadServerUriInvalid:                                             "The ServerUri is not a valid URI.",
	ErrorCodeBadServerNameMissing:                                            "No ServerName was specified.",
	ErrorCodeBadDiscoveryUrlMissing:                                          "No DiscoveryUrl was specified.",
	ErrorCodeBadSempahoreFileMissing:                                         "The semaphore file specified by the client is not valid.",
	ErrorCodeBadRequestTypeInvalid:                                           "The security token request type is not valid.",
	ErrorCodeBadSecurityModeRejected:                                         "The security mode does not meet the requirements set by the server.",
	ErrorCodeBadSecurityPolicyRejected:                                       "The security policy does not meet the requirements set by the server.",
	ErrorCodeBadTooManySessions:                                              "The server has reached its maximum number of sessions.",
	ErrorCodeBadUserSignatureInvalid:                                         "The user token signature is missing or invalid.",
	ErrorCodeBadApplicationSignatureInvalid:                                  "The signature generated with the client certificate is missing or invalid.",
	ErrorCodeBadNoValidCertificates:                                          "The client did not provide at least one software certificate that is valid and meets the profile requirements for the server.",
	ErrorCodeBadIdentityChangeNotSupported:                                   "The server does not support changing the user identity assigned to the session.",
	ErrorCodeBadRequestCancelledByRequest:                                    "The request was cancelled by the client with the Cancel service.",
	ErrorCodeBadParentNodeIdInvalid:                                          "The parent node id does not to refer to a valid node.",
	ErrorCodeBadReferenceNotAllowed:                                          "The reference could not be created because it violates constraints imposed by the data model.",
	ErrorCodeBadNodeIdRejected:                                               "The requested node id was reject because it was either invalid or server does not allow node ids to be specified by the client.",
	ErrorCodeBadNodeIdExists:                                                 "The requested node id is already used by another node.",
	ErrorCodeBadNodeClassInvalid:                                             "The node class is not valid.",
	ErrorCodeBadBrowseNameInvalid:                                            "The browse name is invalid.",
	ErrorCodeBadBrowseNameDuplicated:                                         "The browse name is not unique among nodes that share the same relationship with the parent.",
	ErrorCodeBadNodeAttributesInvalid:                                        "The node attributes are not valid for the node class.",
	ErrorCodeBadTypeDefinitionInvalid:                                        "The type definition node id does not reference an appropriate type node.",
	ErrorCodeBadSourceNodeIdInvalid:                                          "The source node id does not reference a valid node.",
	ErrorCodeBadTargetNodeIdInvalid:                                          "The target node id does not reference a valid node.",
	ErrorCodeBadDuplicateReferenceNotAllowed:                                 "The reference type between the nodes is already defined.",
	ErrorCodeBadInvalidSelfReference:                                         "The server does not allow this type of self reference on this node.",
	ErrorCodeBadReferenceLocalOnly:                                           "The reference type is not valid for a reference to a remote server.",
	ErrorCodeBadNoDeleteRights:                                               "The server will not allow the node to be deleted.",
	ErrorCodeUncertainReferenceNotDeleted:                                    "The server was not able to delete all target references.",
	ErrorCodeBadServerIndexInvalid:                                           "The server index is not valid.",
	ErrorCodeBadViewIdUnknown:                                                "The view id does not refer to a valid view node.",
	ErrorCodeBadViewTimestampInvalid:                                         "The view timestamp is not available or not supported.",
	ErrorCodeBadViewParameterMismatch:                                        "The view parameters are not consistent with each other.",
	ErrorCodeBadViewVersionInvalid:                                           "The view version is not available or not supported.",
	ErrorCodeUncertainNotAllNodesAvailable:                                   "The list of references may not be complete because the underlying system is not available.",
	ErrorCodeGoodResultsMayBeIncomplete:                                      "The server should have followed a reference to a node in a remote server but did not. The result set may be incomplete.",
	ErrorCodeBadNotTypeDefinition:                                            "The provided Nodeid was not a type definition nodeid.",
	ErrorCodeUncertainReferenceOutOfServer:                                   "One of the references to follow in the relative path references to a node in the address space in another server.",
	ErrorCodeBadTooManyMatches:                                               "The requested operation has too many matches to return.",
	ErrorCodeBadQueryTooComplex:                                              "The requested operation requires too many resources in the server.",
	ErrorCodeBadNoMatch:                                                      "The requested operation has no match to return.",
	ErrorCodeBadMaxAgeInvalid:                                                "The max age parameter is invalid.",
	ErrorCodeBadSecurityModeInsufficient:                                     "The operation is not permitted over the current secure channel.",
	ErrorCodeBadHistoryOperationInvalid:                                      "The history details parameter is not valid.",
	ErrorCodeBadHistoryOperationUnsupported:                                  "The server does not support the requested operation.",
	ErrorCodeBadInvalidTimestampArgument:                                     "The defined timestamp to return was invalid.",
	ErrorCodeBadWriteNotSupported:                                            "The server does not support writing the combination of value, status and timestamps provided.",
	ErrorCodeBadTypeMismatch:                                                 "The value supplied for the attribute is not of the same type as the attribute's value.",
	ErrorCodeBadMethodInvalid:                                                "The method id does not refer to a method for the specified object.",
	ErrorCodeBadArgumentsMissing:                                             "The client did not specify all of the input arguments for the method.",
	ErrorCodeBadNotExecutable:                                                "The executable attribute does not allow the execution of the method.",
	ErrorCodeBadTooManySubscriptions:                                         "The server has reached its maximum number of subscriptions.",
	ErrorCodeBadTooManyPublishRequests:                                       "The server has reached the maximum number of queued publish requests.",
	ErrorCodeBadNoSubscription:                                               "There is no subscription available for this session.",
	ErrorCodeBadSequenceNumberUnknown:                                        "The sequence number is unknown to the server.",
	ErrorCodeGoodRetransmissionQueueNotSupported:                             "The Server does not support retransmission queue and acknowledgement of sequence numbers is not available.",
	ErrorCodeBadMessageNotAvailable:                                          "The requested notification message is no longer available.",
	ErrorCodeBadInsufficientClientProfile:                                    "The client of the current session does not support one or more Profiles that are necessary for the subscription.",
	ErrorCodeBadStateNotActive:                                               "The sub-state machine is not currently active.",
	ErrorCodeBadAlreadyExists:                                                "An equivalent rule already exists.",
	ErrorCodeBadTcpServerTooBusy:                                             "The server cannot process the request because it is too busy.",
	ErrorCodeBadTcpMessageTypeInvalid:                                        "The type of the message specified in the header invalid.",
	ErrorCodeBadTcpSecureChannelUnknown:                                      "The SecureChannelId and/or TokenId are not currently in use.",
	ErrorCodeBadTcpMessageTooLarge:                                           "The size of the message chunk specified in the header is too large.",
	ErrorCodeBadTcpNotEnoughResources:                                        "There are not enough resources to process the request.",
	ErrorCodeBadTcpInternalError:                                             "An internal error occurred.",
	ErrorCodeBadTcpEndpointUrlInvalid:                                        "The server does not recognize the QueryString specified.",
	ErrorCodeBadRequestInterrupted:                                           "The request could not be sent because of a network interruption.",
	ErrorCodeBadRequestTimeout:                                               "Timeout occurred while processing the request.",
	ErrorCodeBadSecureChannelClosed:                                          "The secure channel has been closed.",
	ErrorCodeBadSecureChannelTokenUnknown:                                    "The token has expired or is not recognized.",
	ErrorCodeBadSequenceNumberInvalid:                                        "The sequence number is not valid.",
	ErrorCodeBadProtocolVersionUnsupported:                                   "The applications do not have compatible protocol versions.",
	ErrorCodeBadConfigurationError:                                           "There is a problem with the configuration that affects the usefulness of the value.",
	ErrorCodeBadNotConnected:                                                 "The variable should receive its value from another variable, but has never been configured to do so.",
	ErrorCodeBadDeviceFailure:                                                "There has been a failure in the device/data source that generates the value that has affected the value.",
	ErrorCodeBadSensorFailure:                                                "There has been a failure in the sensor from which the value is derived by the device/data source.",
	ErrorCodeBadOutOfService:                                                 "The source of the data is not operational.",
	ErrorCodeBadDeadbandFilterInvalid:                                        "The deadband filter is not valid.",
	ErrorCodeUncertainNoCommunicationLastUsableValue:                         "Communication to the data source has failed. The variable value is the last value that had a good quality.",
	ErrorCodeUncertainLastUsableValue:                                        "Whatever was updating this value has stopped doing so.",
	ErrorCodeUncertainSubstituteValue:                                        "The value is an operational value that was manually overwritten.",
	ErrorCodeUncertainInitialValue:                                           "The value is an initial value for a variable that normally receives its value from another variable.",
	ErrorCodeUncertainSensorNotAccurate:                                      "The value is at one of the sensor limits.",
	ErrorCodeUncertainEngineeringUnitsExceeded:                               "The value is outside of the range of values defined for this parameter.",
	ErrorCodeUncertainSubNormal:                                              "The data value is derived from multiple sources and has less than the required number of Good sources.",
	ErrorCodeGoodLocalOverride:                                               "The value has been overridden.",
	ErrorCodeGoodSubNormal:                                                   "The value is derived from multiple sources and has the required number of Good sources, but less than the full number of Good sources.",
	ErrorCodeBadRefreshInProgress:                                            "This Condition refresh failed, a Condition refresh operation is already in progress.",
	ErrorCodeBadConditionAlreadyDisabled:                                     "This condition has already been disabled.",
	ErrorCodeBadConditionAlreadyEnabled:                                      "This condition has already been enabled.",
	ErrorCodeBadConditionDisabled:                                            "Property not available, this condition is disabled.",
	ErrorCodeBadEventIdUnknown:                                               "The specified event id is not recognized.",
	ErrorCodeBadEventNotAcknowledgeable:                                      "The event cannot be acknowledged.",
	ErrorCodeBadDialogNotActive:                                              "The dialog condition is not active.",
	ErrorCodeBadDialogResponseInvalid:                                        "The response is not valid for the dialog.",
	ErrorCodeBadConditionBranchAlreadyAcked:                                  "The condition branch has already been acknowledged.",
	ErrorCodeBadConditionBranchAlreadyConfirmed:                              "The condition branch has already been confirmed.",
	ErrorCodeBadConditionAlreadyShelved:                                      "The condition has already been shelved.",
	ErrorCodeBadConditionNotShelved:                                          "The condition is not currently shelved.",
	ErrorCodeBadShelvingTimeOutOfRange:                                       "The shelving time not within an acceptable range.",
	ErrorCodeBadNoData:                                                       "No data exists for the requested time range or event filter.",
	ErrorCodeBadBoundNotFound:                                                "No data found to provide upper or lower bound value.",
	ErrorCodeBadBoundNotSupported:                                            "The server cannot retrieve a bound for the variable.",
	ErrorCodeBadDataLost:                                                     "Data is missing due to collection started/stopped/lost.",
	ErrorCodeBadDataUnavailable:                                              "Expected data is unavailable for the requested time range due to an un-mounted volume, an off-line archive or tape, or similar reason for temporary unavailability.",
	ErrorCodeBadEntryExists:                                                  "The data or event was not successfully inserted because a matching entry exists.",
	ErrorCodeBadNoEntryExists:                                                "The data or event was not successfully updated because no matching entry exists.",
	ErrorCodeBadTimestampNotSupported:                                        "The Client requested history using a TimestampsToReturn the Server does not support.",
	ErrorCodeGoodEntryInserted:                                               "The data or event was successfully inserted into the historical database.",
	ErrorCodeGoodEntryReplaced:                                               "The data or event field was successfully replaced in the historical database.",
	ErrorCodeUncertainDataSubNormal:                                          "The aggregate value is derived from multiple values and has less than the required number of Good values.",
	ErrorCodeGoodNoData:                                                      "No data exists for the requested time range or event filter.",
	ErrorCodeGoodMoreData:                                                    "More data is available in the time range beyond the number of values requested.",
	ErrorCodeBadAggregateListMismatch:                                        "The requested number of Aggregates does not match the requested number of NodeIds.",
	ErrorCodeBadAggregateNotSupported:                                        "The requested Aggregate is not support by the server.",
	ErrorCodeBadAggregateInvalidInputs:                                       "The aggregate value could not be derived due to invalid data inputs.",
	ErrorCodeBadAggregateConfigurationRejected:                               "The aggregate configuration is not valid for specified node.",
	ErrorCodeGoodDataIgnored:                                                 "The request specifies fields which are not valid for the EventType or cannot be saved by the historian.",
	ErrorCodeBadRequestNotAllowed:                                            "The request was rejected by the server because it did not meet the criteria set by the server.",
	ErrorCodeBadRequestNotComplete:                                           "The request has not been processed by the server yet.",
	ErrorCodeBadTransactionPending:                                           "The operation is not allowed because a transaction is in progress.",
	ErrorCodeBadTicketRequired:                                               "The device identity needs a ticket before it can be accepted.",
	ErrorCodeBadTicketInvalid:                                                "The device identity needs a ticket before it can be accepted.",
	ErrorCodeBadLocked:                                                       "The requested operation is not allowed, because the Node is locked by a different application.",
	ErrorCodeBadRequiresLock:                                                 "The requested operation is not allowed, because the Node is not locked by the application.",
	ErrorCodeGoodEdited:                                                      "The value does not come from the real source and has been edited by the server.",
	ErrorCodeGoodPostActionFailed:                                            "There was an error in execution of these post-actions.",
	ErrorCodeUncertainDominantValueChanged:                                   "The related EngineeringUnit has been changed but the Variable Value is still provided based on the previous unit.",
	ErrorCodeGoodDependentValueChanged:                                       "A dependent value has been changed but the change has not been applied to the device.",
	ErrorCodeBadDominantValueChanged:                                         "The related EngineeringUnit has been changed but this change has not been applied to the device. The Variable Value is still dependent on the previous unit but its status is currently Bad.",
	ErrorCodeUncertainDependentValueChanged:                                  "A dependent value has been changed but the change has not been applied to the device. The quality of the dominant variable is uncertain.",
	ErrorCodeBadDependentValueChanged:                                        "A dependent value has been changed but the change has not been applied to the device. The quality of the dominant variable is Bad.",
	ErrorCodeGoodEdited_DependentValueChanged:                                "It is delivered with a dominant Variable value when a dependent Variable has changed but the change has not been applied.",
	ErrorCodeGoodEdited_DominantValueChanged:                                 "It is delivered with a dependent Variable value when a dominant Variable has changed but the change has not been applied.",
	ErrorCodeGoodEdited_DominantValueChanged_DependentValueChanged:           "It is delivered with a dependent Variable value when a dominant or dependent Variable has changed but change has not been applied.",
	ErrorCodeBadEdited_OutOfRange:                                            "It is delivered with a Variable value when Variable has changed but the value is not legal.",
	ErrorCodeBadInitialValue_OutOfRange:                                      "It is delivered with a Variable value when a source Variable has changed but the value is not legal.",
	ErrorCodeBadOutOfRange_DominantValueChanged:                              "It is delivered with a dependent Variable value when a dominant Variable has changed and the value is not legal.",
	ErrorCodeBadEdited_OutOfRange_DominantValueChanged:                       "It is delivered with a dependent Variable value when a dominant Variable has changed, the value is not legal and the change has not been applied.",
	ErrorCodeBadOutOfRange_DominantValueChanged_DependentValueChanged:        "It is delivered with a dependent Variable value when a dominant or dependent Variable has changed and the value is not legal.",
	ErrorCodeBadEdited_OutOfRange_DominantValueChanged_DependentValueChanged: "It is delivered with a dependent Variable value when a dominant or dependent Variable has changed, the value is not legal and the change has not been applied.",
	ErrorCodeGoodCommunicationEvent:                                          "The communication layer has raised an event.",
	ErrorCodeGoodShutdownEvent:                                               "The system is shutting down.",
	ErrorCodeGoodCallAgain:                                                   "The operation is not finished and needs to be called again.",
	ErrorCodeGoodNonCriticalTimeout:                                          "A non-critical timeout occurred.",
	ErrorCodeBadInvalidArgument:                                              "One or more arguments are invalid.",
	ErrorCodeBadConnectionRejected:                                           "Could not establish a network connection to remote server.",
	ErrorCodeBadDisconnect:                                                   "The server has disconnected from the client.",
	ErrorCodeBadConnectionClosed:                                             "The network connection has been closed.",
	ErrorCodeBadInvalidState:                                                 "The operation cannot be completed because the object is closed, uninitialized or in some other invalid state.",
	ErrorCodeBadEndOfStream:                                                  "Cannot move beyond end of the stream.",
	ErrorCodeBadNoDataAvailable:                                              "No data is currently available for reading from a non-blocking stream.",
	ErrorCodeBadWaitingForResponse:                                           "The asynchronous operation is waiting for a response.",
	ErrorCodeBadOperationAbandoned:                                           "The asynchronous operation was abandoned by the caller.",
	ErrorCodeBadExpectedStreamToBlock:                                        "The stream did not return all data requested (possibly because it is a non-blocking stream).",
	ErrorCodeBadWouldBlock:                                                   "Non blocking behaviour is required and the operation would block.",
	ErrorCodeBadSyntaxError:                                                  "A value had an invalid syntax.",
	ErrorCodeBadMaxConnectionsReached:                                        "The operation could not be finished because all available connections are in use.",
	ErrorCodeUncertainTransducerInManual:                                     "The value may not be accurate because the transducer is in manual mode.",
	ErrorCodeUncertainSimulatedValue:                                         "The value is simulated.",
	ErrorCodeUncertainSensorCalibration:                                      "The value may not be accurate due to a sensor calibration fault.",
	ErrorCodeUncertainConfigurationError:                                     "The value may not be accurate due to a configuration issue.",
	ErrorCodeGoodCascadeInitializationAcknowledged:                           "The value source supports cascade handshaking and the value has been Initialized based on an initialization request from a cascade secondary.",
	ErrorCodeGoodCascadeInitializationRequest:                                "The value source supports cascade handshaking and is requesting initialization of a cascade primary.",
	ErrorCodeGoodCascadeNotInvited:                                           "The value source supports cascade handshaking, however, the source’s current state does not allow for cascade.",
	ErrorCodeGoodCascadeNotSelected:                                          "The value source supports cascade handshaking, however, the source has not selected the corresponding cascade primary for use.",
	ErrorCodeGoodFaultStateActive:                                            "There is a fault state condition active in the value source.",
	ErrorCodeGoodInitiateFaultState:                                          "A fault state condition is being requested of the destination.",
	ErrorCodeGoodCascade:                                                     "The value is accurate, and the signal source supports cascade handshaking.",
	ErrorCodeBadDataSetIdInvalid:                                             "The DataSet specified for the DataSetWriter creation is invalid.",
}
