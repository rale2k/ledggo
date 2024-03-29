package api

import (
	"fmt"
	"ledggo/domain"
	"ledggo/ledger"
	"ledggo/p2p"
	"ledggo/utils"
	"net/http"

	"github.com/gin-gonic/gin"
)

func GetKnownNodes(c *gin.Context) {
	c.JSON(http.StatusOK, utils.Nodes)
}

func GetLedgerLength(c *gin.Context) {
	c.JSON(http.StatusOK, ledger.GetLedgerLength())
}

func GetLastBlock(c *gin.Context) {
	lastBlock, err := ledger.GetLastBlock()

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, lastBlock)
}

func GetBlock(c *gin.Context) {
	var hash = c.Query("hash")

	if hash == "" {
		c.JSON(http.StatusOK, utils.Blocks)
		return
	}

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

	if block.Hash == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Block hash is required"})
		return
	}

	if block.Data == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Block data is required"})
		return
	}

	if _, err := ledger.GetBlockWithHash(block.Hash); err == nil {
		c.Status(http.StatusOK)
		return
	}

	if err := ledger.AddNewBlock(block); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		fmt.Println(err.Error())
		return
	}

	go p2p.DistributeNewBlock(block, c.Request.RemoteAddr)

	c.Status(http.StatusOK)
}
