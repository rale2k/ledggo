package utils

import (
	"os"
	"path/filepath"
)

func GetFilePathInWorkingDir(fileName string) (string, error) {
	cwd, err := os.Getwd()
	if err != nil {
		return "", err
	}

	blockFilePath := filepath.Join(cwd, fileName)
	return blockFilePath, nil
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
