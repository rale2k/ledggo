package ledger

import (
	"fmt"
	"ledggo/domain"
	"ledggo/utils"
)

func AddNewBlock(block domain.Block) error {
	if err := validateBlockHash(&block, utils.Blocks); err != nil {
		return err
	}

	utils.Blocks = append(utils.Blocks, block)
	return nil
}

func GetBlockWithHash(hash string) (block domain.Block, err error) {
	for _, block := range utils.Blocks {
		if block.Hash == hash {
			return block, nil
		}
	}

	return domain.Block{}, fmt.Errorf("could not find block with hash: %s", hash)

}

func BlockExists(hash string) bool {
	_, err := GetBlockWithHash(hash)
	return err == nil
}

func validateBlockHash(newBlock *domain.Block, previousBlocks []domain.Block) error {
	var hash string

	if len(previousBlocks) > 0 {
		lastBlock := previousBlocks[len(previousBlocks)-1]
		hash = utils.GetSha256Hash(newBlock.Data + lastBlock.Hash)
	} else {
		hash = utils.GetSha256Hash(newBlock.Data)
	}

	if newBlock.Hash != hash {
		return fmt.Errorf("block hash is invalid")
	}

	return nil
}
