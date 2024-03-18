package p2p

import (
	"bytes"
	"encoding/json"
	"fmt"
	"ledggo/domain"
	"ledggo/utils"
	"net/http"
)

func DistributeNewBlock(block domain.Block, remoteAddr string) {
	for _, node := range utils.Nodes {
		url := fmt.Sprintf("http://%s:%s/block", node.Ip, node.Port)

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
		_, err = http.Post(url, "application/json", bytes.NewBuffer(reqBody))

		if err != nil {
			fmt.Printf("Node %s unresponsive, removing...\n", node.Ip)
			utils.RemoveNode(node)
		}
	}
}
