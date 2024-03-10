package api

import (
	"ledggo/domain"
	"ledggo/ledger"
	"ledggo/utils"
	"net/http"

	"github.com/gin-gonic/gin"
)

func GetKnownNodes(c *gin.Context) {
	c.IndentedJSON(http.StatusOK, utils.Nodes)
}

func GetBlocks(c *gin.Context) {
	var blocks []domain.Block

	if err := ledger.GetBlocks(&blocks); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, blocks)
}

func GetBlock(c *gin.Context) {
	var hash = c.Param("hash")

	block, err := ledger.GetBlockWithHash(hash)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, block)
}

func PostBlock(c *gin.Context) {
	var block domain.Block
	if err := c.ShouldBindJSON(&block); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	ledger.AddNewBlock(block)

	c.Status(http.StatusOK)
}
