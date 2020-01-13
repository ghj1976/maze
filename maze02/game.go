package main

import (
	"image/color"

	"github.com/hajimehoshi/ebiten"
)

const (
	ScreenWidth  = 420
	ScreenHeight = 600
)

// Game 游戏逻辑
type Game struct {
	maze *Maze
}

// NewGame generates a new Game object.
func NewGame() (*Game, error) {
	g := &Game{}

	g.maze, _ = NewMaze(5, 6)

	return g, nil
}

// Draw 绘图
func (g *Game) Draw(screen *ebiten.Image) {
	screen.Fill(color.RGBA{0xfa, 0xf8, 0xef, 0xff})

	g.maze.Draw(screen)
}
