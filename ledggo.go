package main

import (
	"flag"
	"ledggo/api"
	"ledggo/utils"
	"strconv"

	"github.com/gin-gonic/gin"
)

func main() {
	var port = flag.Int("port", 8080, "Port to run the node on")
	flag.Parse()

	utils.AppendNodes(flag.Args())

	router := gin.Default()

	router.GET("/nodes", api.GetKnownNodes)
	router.GET("/blocks/:hash", api.GetBlock)
	router.GET("/blocks", api.GetBlocks)
	router.POST("/block", api.PostBlock)

	router.Run("localhost:" + strconv.Itoa(*port))
}
