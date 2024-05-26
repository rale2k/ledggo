package utils

import (
	"ledggo/domain"
	"strconv"
	"strings"
)

// Globals
var RunningPort = 8080

var Blocks = []domain.Block{}
var BlocksInTx = []domain.Block{}

var State = domain.State{
	Role: domain.COORDINATOR,
	Coordinator: domain.Node{
		Ip:   "127.0.0.1",
		Port: "8080",
	},
	Nodes: []domain.Node{},
}

func AppendNodesFromIPStringCSV(nodesCsv string) {
	if len(nodesCsv) == 0 {
		return
	}
	parts := strings.Split(nodesCsv, ";")
	for _, node := range parts {
		AppendNodesFromIPString(node)
	}
}

func AppendNodes(nodes []domain.Node) {
	for _, node := range nodes {
		if IsNodeValid(node) {
			State.Nodes = append(State.Nodes, node)
		}
	}
}

func SetCoordinator(node domain.Node) {
	State.Coordinator = node
	State.Role = domain.NODE
}

func AppendNodesFromIPString(node string) {
	parts := strings.Split(node, ":")
	if len(parts) == 2 {
		ip := parts[0]
		port := parts[1]
		newNode := domain.Node{
			Ip:   ip,
			Port: port,
		}

		if IsNodeValid(newNode) {
			State.Nodes = append(State.Nodes, newNode)
		}
	}
}

func IsNodeValid(node domain.Node) bool {
	if node.Ip == "127.0.0.1" && node.Port == strconv.Itoa(RunningPort) {
		return false
	}

	for _, n := range State.Nodes {
		if n.Ip == node.Ip && n.Port == node.Port {
			return false
		}
	}
	return true
}

func RemoveNode(nodeToRemove domain.Node) {
	for i, node := range State.Nodes {
		if node.Ip == nodeToRemove.Ip && node.Port == nodeToRemove.Port {
			State.Nodes = append(State.Nodes[:i], State.Nodes[i+1:]...)
			break
		}
	}
}
