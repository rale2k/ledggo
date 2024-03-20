package p2p

import (
	"encoding/json"
	"fmt"
	"io"
	"ledggo/domain"
	"ledggo/utils"
	"strconv"
)

func QueryLedgerFromNodes() error {
	var longestLedgerLength int
	var longestLedgerNode domain.Node

	for _, node := range utils.Nodes {
		ledgerLength, err := GetLedgerLengthFromNode(node)
		if err != nil {
			continue
		}

		if ledgerLength > longestLedgerLength {
			longestLedgerLength = ledgerLength
			longestLedgerNode = node
		}
	}

	if longestLedgerNode.Ip == "" {
		return fmt.Errorf("no nodes with available ledgers")
	}

	ledger, err := GetLedgerFromNode(longestLedgerNode)
	if err != nil {
		return err
	}

	utils.Blocks = ledger

	return nil
}

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

func GetLedgerFromNode(node domain.Node) ([]domain.Block, error) {
	resp, err := utils.DoRequestToNode("GET", "/blocks", nil, node)
	if err != nil {
		return nil, err
	}

	var blocks []domain.Block
	err = json.NewDecoder(resp.Body).Decode(&blocks)
	if err != nil {
		return nil, err
	}

	return blocks, nil
}

func GetLedgerLengthFromNode(node domain.Node) (int, error) {
	resp, err := utils.DoRequestToNode("GET", "/blocks/count", nil, node)
	if err != nil {
		return 0, err
	}

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return 0, err
	}

	respString := string(body)

	return strconv.Atoi(respString)
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
