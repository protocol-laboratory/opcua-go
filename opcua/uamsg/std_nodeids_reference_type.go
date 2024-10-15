package uamsg

// https://github.com/OPCFoundation/UA-Nodeset/blob/latest/Schema/NodeIds.csv
var (
	ReferenceTypeReferences                          NodeId = NodeId{TwoByte, 0, uint16(31)}
	ReferenceTypeNonHierarchicalReferences           NodeId = NodeId{TwoByte, 0, uint16(32)}
	ReferenceTypeHierarchicalReferences              NodeId = NodeId{TwoByte, 0, uint16(33)}
	ReferenceTypeHasChild                            NodeId = NodeId{TwoByte, 0, uint16(34)}
	ReferenceTypeOrganizes                           NodeId = NodeId{TwoByte, 0, uint16(35)}
	ReferenceTypeHasEventSource                      NodeId = NodeId{TwoByte, 0, uint16(36)}
	ReferenceTypeHasModellingRule                    NodeId = NodeId{TwoByte, 0, uint16(37)}
	ReferenceTypeHasEncoding                         NodeId = NodeId{TwoByte, 0, uint16(38)}
	ReferenceTypeHasDescription                      NodeId = NodeId{TwoByte, 0, uint16(39)}
	ReferenceTypeHasTypeDefinition                   NodeId = NodeId{TwoByte, 0, uint16(40)}
	ReferenceTypeGeneratesEvent                      NodeId = NodeId{TwoByte, 0, uint16(41)}
	ReferenceTypeAggregates                          NodeId = NodeId{TwoByte, 0, uint16(44)}
	ReferenceTypeHasSubtype                          NodeId = NodeId{TwoByte, 0, uint16(45)}
	ReferenceTypeHasProperty                         NodeId = NodeId{TwoByte, 0, uint16(46)}
	ReferenceTypeHasComponent                        NodeId = NodeId{TwoByte, 0, uint16(47)}
	ReferenceTypeHasNotifier                         NodeId = NodeId{TwoByte, 0, uint16(48)}
	ReferenceTypeHasOrderedComponent                 NodeId = NodeId{TwoByte, 0, uint16(49)}
	ReferenceTypeFromState                           NodeId = NodeId{TwoByte, 0, uint16(51)}
	ReferenceTypeToState                             NodeId = NodeId{TwoByte, 0, uint16(52)}
	ReferenceTypeHasCause                            NodeId = NodeId{TwoByte, 0, uint16(53)}
	ReferenceTypeHasEffect                           NodeId = NodeId{TwoByte, 0, uint16(54)}
	ReferenceTypeHasHistoricalConfiguration          NodeId = NodeId{TwoByte, 0, uint16(56)}
	ReferenceTypeHasSubStateMachine                  NodeId = NodeId{TwoByte, 0, uint16(117)}
	ReferenceTypeHasArgumentDescription              NodeId = NodeId{TwoByte, 0, uint16(129)}
	ReferenceTypeHasOptionalInputArgumentDescription NodeId = NodeId{TwoByte, 0, uint16(131)}
	ReferenceTypeAlwaysGeneratesEvent                NodeId = NodeId{FourByte, 0, uint16(3065)}
	ReferenceTypeHasTrueSubState                     NodeId = NodeId{FourByte, 0, uint16(9004)}
	ReferenceTypeHasFalseSubState                    NodeId = NodeId{FourByte, 0, uint16(9005)}
	ReferenceTypeHasCondition                        NodeId = NodeId{FourByte, 0, uint16(9006)}
	ReferenceTypeHasPubSubConnection                 NodeId = NodeId{FourByte, 0, uint16(14476)}
	ReferenceTypeDataSetToWriter                     NodeId = NodeId{FourByte, 0, uint16(14936)}
	ReferenceTypeHasGuard                            NodeId = NodeId{FourByte, 0, uint16(15112)}
	ReferenceTypeHasDataSetWriter                    NodeId = NodeId{FourByte, 0, uint16(15296)}
	ReferenceTypeHasDataSetReader                    NodeId = NodeId{FourByte, 0, uint16(15297)}
	ReferenceTypeHasAlarmSuppressionGroup            NodeId = NodeId{FourByte, 0, uint16(16361)}
	ReferenceTypeAlarmGroupMember                    NodeId = NodeId{FourByte, 0, uint16(16362)}
	ReferenceTypeHasEffectDisable                    NodeId = NodeId{FourByte, 0, uint16(17276)}
	ReferenceTypeHasDictionaryEntry                  NodeId = NodeId{FourByte, 0, uint16(17597)}
	ReferenceTypeHasInterface                        NodeId = NodeId{FourByte, 0, uint16(17603)}
	ReferenceTypeHasAddIn                            NodeId = NodeId{FourByte, 0, uint16(17604)}
	ReferenceTypeHasEffectEnable                     NodeId = NodeId{FourByte, 0, uint16(17983)}
	ReferenceTypeHasEffectSuppressed                 NodeId = NodeId{FourByte, 0, uint16(17984)}
	ReferenceTypeHasEffectUnsuppressed               NodeId = NodeId{FourByte, 0, uint16(17985)}
	ReferenceTypeHasWriterGroup                      NodeId = NodeId{FourByte, 0, uint16(18804)}
	ReferenceTypeHasReaderGroup                      NodeId = NodeId{FourByte, 0, uint16(18805)}
	ReferenceTypeAliasFor                            NodeId = NodeId{FourByte, 0, uint16(23469)}
	ReferenceTypeIsDeprecated                        NodeId = NodeId{FourByte, 0, uint16(23562)}
	ReferenceTypeHasStructuredComponent              NodeId = NodeId{FourByte, 0, uint16(24136)}
	ReferenceTypeAssociatedWith                      NodeId = NodeId{FourByte, 0, uint16(24137)}
	ReferenceTypeUsesPriorityMappingTable            NodeId = NodeId{FourByte, 0, uint16(25237)}
	ReferenceTypeHasLowerLayerInterface              NodeId = NodeId{FourByte, 0, uint16(25238)}
	ReferenceTypeIsExecutableOn                      NodeId = NodeId{FourByte, 0, uint16(25253)}
	ReferenceTypeControls                            NodeId = NodeId{FourByte, 0, uint16(25254)}
	ReferenceTypeUtilizes                            NodeId = NodeId{FourByte, 0, uint16(25255)}
	ReferenceTypeRequires                            NodeId = NodeId{FourByte, 0, uint16(25256)}
	ReferenceTypeIsPhysicallyConnectedTo             NodeId = NodeId{FourByte, 0, uint16(25257)}
	ReferenceTypeRepresentsSameEntityAs              NodeId = NodeId{FourByte, 0, uint16(25258)}
	ReferenceTypeRepresentsSameHardwareAs            NodeId = NodeId{FourByte, 0, uint16(25259)}
	ReferenceTypeRepresentsSameFunctionalityAs       NodeId = NodeId{FourByte, 0, uint16(25260)}
	ReferenceTypeIsHostedBy                          NodeId = NodeId{FourByte, 0, uint16(25261)}
	ReferenceTypeHasPhysicalComponent                NodeId = NodeId{FourByte, 0, uint16(25262)}
	ReferenceTypeHasContainedComponent               NodeId = NodeId{FourByte, 0, uint16(25263)}
	ReferenceTypeHasAttachedComponent                NodeId = NodeId{FourByte, 0, uint16(25264)}
	ReferenceTypeIsExecutingOn                       NodeId = NodeId{FourByte, 0, uint16(25265)}
	ReferenceTypeHasPushedSecurityGroup              NodeId = NodeId{FourByte, 0, uint16(25345)}
	ReferenceTypeAlarmSuppressionGroupMember         NodeId = NodeId{FourByte, 0, uint16(32059)}
	ReferenceTypeHasKeyValueDescription              NodeId = NodeId{FourByte, 0, uint16(32407)}
	ReferenceTypeHasEngineeringUnitDetails           NodeId = NodeId{FourByte, 0, uint16(32558)}
	ReferenceTypeHasQuantity                         NodeId = NodeId{FourByte, 0, uint16(32559)}
	ReferenceTypeHasCurrentData                      NodeId = NodeId{FourByte, 0, uint16(32633)}
	ReferenceTypeHasCurrentEvent                     NodeId = NodeId{FourByte, 0, uint16(32634)}
	ReferenceTypeHasReferenceDescription             NodeId = NodeId{FourByte, 0, uint16(32679)}
)
