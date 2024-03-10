package ledger

import (
	"crypto/sha256"
	"encoding/hex"
	"encoding/json"
	"fmt"
	"ledggo/domain"
	"ledggo/utils"
)

func AddNewBlock(block domain.Block) error {
	data, err := utils.ReadBlockDataFromFile()
	if err != nil {
		return err
	}

	var blocks []domain.Block
	if err := json.Unmarshal(data, &blocks); err != nil {
		return err
	}

	setBlockHash(&block, blocks)

	blocks = append(blocks, block)

	return utils.WriteBlockDataToFile(blocks)
}

func GetBlockWithHash(hash string) (block domain.Block, err error) {
	data, err := utils.ReadBlockDataFromFile()
	if err != nil {
		return domain.Block{}, err
	}

	var blocks []domain.Block
	if err := json.Unmarshal(data, &blocks); err != nil {
		return domain.Block{}, err
	}

	for _, block := range blocks {
		if block.Hash == hash {
			return block, nil
		}
	}

	return domain.Block{}, fmt.Errorf("could not find block with hash: %s", hash)

}

func GetBlocks(blocks *[]domain.Block) error {
	data, err := utils.ReadBlockDataFromFile()
	if err != nil {
		return err
	}

	err = json.Unmarshal(data, &blocks)
	if err != nil {
		return err
	}

	return nil
}

func setBlockHash(newBlock *domain.Block, previousBlocks []domain.Block) {
	if len(previousBlocks) > 0 {
		lastBlock := previousBlocks[len(previousBlocks)-1]
		x := sha256.Sum256([]byte(newBlock.Data + lastBlock.Hash))
		newBlock.Hash = hex.EncodeToString(x[:])
	} else {
		x := sha256.Sum256([]byte(newBlock.Data))
		newBlock.Hash = hex.EncodeToString(x[:])
	}
}
