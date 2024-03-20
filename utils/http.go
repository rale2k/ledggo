package utils

import (
	"bytes"
	"fmt"
	"ledggo/domain"
	"net/http"
)

func DoRequestToNode(method string, path string, reqBody []byte, node domain.Node) (*http.Response, error) {
	url := fmt.Sprintf("http://%s:%s%s", node.Ip, node.Port, path)

	req, err := http.NewRequest(method, url, bytes.NewBuffer(reqBody))

	if err != nil {
		return nil, fmt.Errorf("error creating request: %v", err)
	}

	// TODO: needs to be un-hardcoded from local loopback to get the actual IP
	req.Header.Set("node-ip", fmt.Sprintf("%s:%d", "127.0.0.1", RunningPort))

	response, err := http.DefaultClient.Do(req)
	if err != nil {
		fmt.Printf("Node %s unresponsive, removing...%s\n", node.Ip, err)
		RemoveNode(node)
	}

	return response, err
}
