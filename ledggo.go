package main

import (
	"flag"
	"ledggo/api"
	"ledggo/p2p"
	"ledggo/utils"
	"strconv"

	"github.com/gin-gonic/gin"
)

func main() {
	var port = flag.Int("port", utils.RunningPort, "Port to run the node on")
	var nodes = flag.String("nodes", "", "{ip}:{port} of known nodes to connect to separated by semicolons")
	flag.Parse()

	utils.RunningPort = *port

	utils.AppendNodesFromIPStringCSV(*nodes)
	p2p.GetNodesFromKnownNodes()
	p2p.QueryLedgerFromNodes()

	router := gin.Default()

	router.Use(utils.SaveNodeRequestIp)

	router.GET("/info", api.GetNodeInfo)
	router.GET("/blocks", api.GetBlock)
	router.GET("/blocks/last", api.GetLastBlock)
	router.GET("/blocks/count", api.GetLedgerLength)
	router.POST("/blocks", api.PostBlock)
	router.POST("/commit", api.CommitBlocksInTx)
	router.POST("/cancel", api.CancelBlocksInTx)

	router.Run("localhost:" + strconv.Itoa(*port))
}
