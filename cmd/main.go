package main

import (
	"log"

	"github.com/crislainesc/space_game/internal/game"
	"github.com/hajimehoshi/ebiten/v2"
)

func main() {
	g := game.NewGame()

	ebiten.SetFullscreen(true)
	ebiten.SetScreenClearedEveryFrame(true)
	if err := ebiten.RunGame(g); err != nil {
		log.Fatal(err)
	}
}
