package main

import (
	"arena/server"
	"arena/utils"
)

func main() {
	utils.Rand.InitSeed(utils.Rand{})
	server.Init()
}
