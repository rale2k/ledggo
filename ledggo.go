package main

import (
	"fmt"
	"ledggo/api"
	"ledggo/utils"
	"os"
	"strconv"

	"github.com/gin-gonic/gin"
)

func main() {
	var port = utils.DefaultPort

	if err := utils.GetOpenPort(&port); err != nil {
		fmt.Println(err.Error())
		os.Exit(0)
	}

	nodes, err := utils.ReadConfig()
	if err != nil {
		fmt.Println(err.Error())
		os.Exit(0)
	}

	utils.Nodes = nodes

	if err := utils.CreateBlockFileIfNotExists(); err != nil {
		fmt.Println(err.Error())
		os.Exit(0)
	}

	router := gin.Default()

	router.GET("/nodes", api.GetKnownNodes)
	router.GET("/blocks/:hash", api.GetBlock)
	router.GET("/blocks", api.GetBlocks)
	router.POST("/block", api.PostBlock)

	router.Run("localhost:" + strconv.Itoa(port))
}
