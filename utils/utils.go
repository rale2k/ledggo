package utils

import (
	"encoding/json"
	"fmt"
	"ledggo/domain"
	"net"
	"os"
	"path/filepath"

	"github.com/google/uuid"
)

func GetFilePathInWorkingDir(fileName string) (string, error) {
	cwd, err := os.Getwd()
	if err != nil {
		return "", err
	}

	blockFilePath := filepath.Join(cwd, fileName)
	return blockFilePath, nil
}

func WriteBlockDataToFile(blocks []domain.Block) error {
	blockFilePath, err := GetFilePathInWorkingDir(BlockFileName)
	if err != nil {
		return err
	}

	data, err := json.Marshal(blocks)
	if err != nil {
		return err
	}
	return os.WriteFile(blockFilePath, data, 0644)
}

func ReadBlockDataFromFile() (data []byte, err error) {
	blockFilePath, err := GetFilePathInWorkingDir(BlockFileName)
	if err != nil {
		return nil, err
	}

	data, err = os.ReadFile(blockFilePath)
	if err != nil {
		return nil, err
	}

	return data, nil
}

func ReadConfig() (nodes []domain.Node, err error) {
	configFilePath, err := GetFilePathInWorkingDir(ConfigFileName)
	if err != nil {
		return nil, err
	}

	data, err := os.ReadFile(configFilePath)
	if err != nil {
		return nil, err
	}

	var config domain.Config
	if err := json.Unmarshal(data, &config); err != nil {
		return nil, err
	}

	return config.Nodes, nil
}

func GetOpenPort(port *int) (err error) {
	if *port < 0 || *port > 65535 {
		return fmt.Errorf("invalid port number: %d", *port)
	}

	address := fmt.Sprintf("localhost:%d", *port)

	var tcpAddr *net.TCPAddr
	if tcpAddr, err = net.ResolveTCPAddr("tcp", address); err != nil {
		return err
	}
	var listener *net.TCPListener
	listener, err = net.ListenTCP("tcp", tcpAddr)
	if err == nil {
		listener.Close()
		*port = listener.Addr().(*net.TCPAddr).Port
		return nil
	} else if *port != 0 {
		listener.Close()
		*port = 0
		return GetOpenPort(port)
	}
	return
}

func CreateBlockFileIfNotExists() (err error) {
	blockFilePath, err := GetFilePathInWorkingDir(BlockFileName)
	if err != nil {
		return err
	}

	_, err = os.Stat(blockFilePath)
	if os.IsNotExist(err) {
		file, err := os.Create(blockFilePath)
		if err != nil {
			return err
		}
		defer file.Close()
	}

	return
}

func GenerateUUID() string {
	id := uuid.New()
	return id.String()
}
