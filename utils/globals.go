package utils

import "ledggo/domain"

// Globals
const DefaultPort = 8080
const BlockFileName = "blocks.json"
const ConfigFileName = "config.json"

var Nodes = []domain.Node{}
var Blocks = []domain.Block{}
