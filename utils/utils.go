package utils

import (
	"crypto/sha256"
	"encoding/hex"
	"encoding/json"
	"ledggo/domain"
	"os"

	"github.com/google/uuid"
)

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

func GenerateUUID() string {
	id := uuid.New()
	return id.String()
}

func GetSha256Hash(data string) string {
	x := sha256.Sum256([]byte(data))
	return hex.EncodeToString(x[:])
}
