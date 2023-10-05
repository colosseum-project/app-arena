package main

import (
	"github.com/colosseum-project/app-arena/server"
	"github.com/colosseum-project/app-arena/utils"
)

func main() {
	utils.Rand.InitSeed(utils.Rand{})
	server.Init()
}
