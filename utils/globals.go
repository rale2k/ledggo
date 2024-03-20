package utils

import (
	"ledggo/domain"
	"strconv"
	"strings"
)

// Globals
var RunningPort = 8080

var Nodes = []domain.Node{}
var Blocks = []domain.Block{}

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
			Nodes = append(Nodes, node)
		}
	}
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
			Nodes = append(Nodes, newNode)
		}
	}
}

func IsNodeValid(node domain.Node) bool {
	if node.Ip == "127.0.0.1" && node.Port == strconv.Itoa(RunningPort) {
		return false
	}

	for _, n := range Nodes {
		if n.Ip == node.Ip && n.Port == node.Port {
			return false
		}
	}
	return true
}

func RemoveNode(nodeToRemove domain.Node) {
	for i, node := range Nodes {
		if node.Ip == nodeToRemove.Ip && node.Port == nodeToRemove.Port {
			Nodes = append(Nodes[:i], Nodes[i+1:]...)
			break
		}
	}
}
