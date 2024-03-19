package utils

import (
	"ledggo/domain"
	"strings"
)

// Globals
const DefaultPort = 8080
const BlockFileName = "blocks.json"
const ConfigFileName = "config.json"

var Nodes = []domain.Node{}
var Blocks = []domain.Block{}

func AppendNodesFromString(nodesCsv string) {
	if len(nodesCsv) == 0 {
		return
	}
	parts := strings.Split(nodesCsv, ";")
	for _, node := range parts {
		parts := strings.Split(node, ":")
		if len(parts) == 2 {
			ip := parts[0]
			port := parts[1]
			newNode := domain.Node{
				Ip:   ip,
				Port: port,
			}

			if !IsNodeDuplicate(newNode) {
				Nodes = append(Nodes, newNode)
			}
		}
	}
}

func AppendNodes(nodes []domain.Node) {
	for _, node := range nodes {
		if !IsNodeDuplicate(node) {
			Nodes = append(Nodes, node)
		}
	}
}

func IsNodeDuplicate(node domain.Node) bool {
	for _, n := range Nodes {
		if n.Ip == node.Ip && n.Port == node.Port {
			return true
		}
	}
	return false
}

func RemoveNode(nodeToRemove domain.Node) {
	for i, node := range Nodes {
		if node.Ip == nodeToRemove.Ip && node.Port == nodeToRemove.Port {
			Nodes = append(Nodes[:i], Nodes[i+1:]...)
			break
		}
	}
}
