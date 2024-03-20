package p2p

import (
	"encoding/json"
	"fmt"
	"ledggo/domain"
	"ledggo/utils"
)

func DistributeNewBlock(block domain.Block, remoteAddr string) {
	for _, node := range utils.Nodes {
		// Don't send the block to the node that sent it to us
		if remoteAddr == fmt.Sprintf("%s:%s", node.Ip, node.Port) {
			continue
		}

		blockJson, err := json.Marshal(block)
		if err != nil {
			fmt.Printf("Error converting block to JSON: %v\n", err)
			return
		}

		reqBody := []byte(blockJson)
		// TODO: should use a thread pool or something, this may get out of hand with many nodes
		go utils.DoRequestToNode("POST", "/blocks", reqBody, node)
	}
}

func GetNodesFromNode(node domain.Node) ([]domain.Node, error) {
	resp, err := utils.DoRequestToNode("GET", "/nodes", nil, node)
	if err != nil {
		return nil, err
	}

	var nodes []domain.Node
	err = json.NewDecoder(resp.Body).Decode(&nodes)
	if err != nil {
		return nil, err
	}

	return nodes, nil
}

func GetNodesFromKnownNodes() {
	for _, node := range utils.Nodes {
		newNodes, err := GetNodesFromNode(node)

		if err != nil {
			continue
		}

		utils.AppendNodes(newNodes)
	}
}
