package main

import (
	"ggengine/engine"
	"ggengine/game"
)

func main() {

	game := game.Create()
	engine := engine.Init()
	engine.Run(game)
}
