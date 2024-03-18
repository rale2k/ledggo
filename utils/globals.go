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

func AppendNodes(nodes []string) {
	if len(nodes) == 0 {
		return
	}
	parts := strings.Split(nodes[0], " ")
	for _, node := range parts {
		parts := strings.Split(node, ":")
		if len(parts) == 2 {
			ip := parts[0]
			port := parts[1]
			newNode := domain.Node{
				Ip:   ip,
				Port: port,
			}
			Nodes = append(Nodes, newNode)
		}
	}
}

func RemoveNode(nodeToRemove domain.Node) {
	for i, node := range Nodes {
		if node.Ip == nodeToRemove.Ip && node.Port == nodeToRemove.Port {
			Nodes = append(Nodes[:i], Nodes[i+1:]...)
			break
		}
	}
}
