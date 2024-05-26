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
	if utils.State.Role == domain.NODE {
		ledger, err := GetLedgerFromNode(utils.State.Coordinator)
		if err != nil {
			return err
		}

		utils.Blocks = ledger
	}

	var longestLedgerLength int
	var longestLedgerNode domain.Node

	for _, node := range utils.State.Nodes {
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
	for _, node := range utils.State.Nodes {
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
		response, _ := utils.DoRequestToNode("POST", "/blocks", reqBody, node)

		if 400 <= response.StatusCode && response.StatusCode < 500 {
			CancelBlocksInTx()
			b, _ := io.ReadAll(response.Body)
			fmt.Printf("Error distributing block tx: %v\n", string(b))
			return
		}
	}

	CommitBlocksInTx()
}

func CommitBlocksInTx() {
	for _, node := range utils.State.Nodes {
		go utils.DoRequestToNode("POST", "/commit", nil, node)
	}
}

func CancelBlocksInTx() {
	for _, node := range utils.State.Nodes {
		go utils.DoRequestToNode("POST", "/cancel", nil, node)
	}
}

func GetNodeInfo(node domain.Node) (domain.State, error) {
	resp, err := utils.DoRequestToNode("GET", "/info", nil, node)
	if err != nil {
		return domain.State{}, err
	}

	var state domain.State
	err = json.NewDecoder(resp.Body).Decode(&state)
	if err != nil {
		return domain.State{}, err
	}

	return state, nil
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
	for _, node := range utils.State.Nodes {
		state, err := GetNodeInfo(node)

		if state.Role == domain.COORDINATOR {
			utils.SetCoordinator(node)
		}

		if state.Role == domain.NODE {
			utils.SetCoordinator(state.Coordinator)
		}

		if err != nil {
			continue
		}

		utils.AppendNodes(state.Nodes)
	}
}
