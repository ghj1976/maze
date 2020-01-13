package main

import (
	"image/color"
	"log"

	"github.com/hajimehoshi/ebiten"
	"github.com/hajimehoshi/ebiten/ebitenutil"
)

const (
	ScreenWidth  = 420
	ScreenHeight = 600
	boardSize    = 4
)

func update(screen *ebiten.Image) error {
	if ebiten.IsDrawingSkipped() {
		return nil
	}

	screen.Fill(color.RGBA{0xfa, 0xf8, 0xef, 0xff})

	square, _ := ebiten.NewImage(16, 16, ebiten.FilterNearest)
	square.Fill(color.Black)
	opts := &ebiten.DrawImageOptions{}
	opts.GeoM.Translate(64, 64)
	screen.DrawImage(square, opts)

	ebitenutil.DrawLine(screen, 20, 20, 60, 70, color.Black)

	return nil
}

func main() {
	if err := ebiten.Run(update, ScreenWidth, ScreenHeight, 1, "测试"); err != nil {
		log.Fatal(err)
	}
}
