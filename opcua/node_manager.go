package opcua

import "opcua-go/opcua/ua"

type NodeManager struct {
}

func NewNodeManager() *NodeManager {
	return &NodeManager{}
}

func (nm *NodeManager) AddNode(node ua.Node) {
}

func (nm *NodeManager) DeleteNode(node ua.Node) {
}

func (nm *NodeManager) GetNode(nodeId string) {
}
