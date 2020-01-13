package main

import (
	"log"

	"github.com/hajimehoshi/ebiten"
)

var (
	game *Game
)

func update(screen *ebiten.Image) error {
	if ebiten.IsDrawingSkipped() {
		return nil
	}
	game.Draw(screen)

	return nil

}

func main() {

	var err error
	game, err = NewGame()
	if err != nil {
		log.Fatal(err)
	}

	if err := ebiten.Run(update, ScreenWidth, ScreenHeight, 1, "迷宫测试"); err != nil {
		log.Fatal(err)
	}

}
