package main

import (
	"flag"
	"log"

	"github.com/ghj1976/maze/maze03/mazemap"
	"github.com/hajimehoshi/ebiten"
)

var (
	game *mazemap.Game

	row = flag.Int("row", 5, "请输入Rows，Y轴的行数。")
	col = flag.Int("col", 6, "请输入Columns，X轴的列数。")
)

func update(screen *ebiten.Image) error {
	if ebiten.IsDrawingSkipped() {
		return nil
	}
	game.Draw(screen)

	return nil

}

func main() {

	// 读取命令参数
	flag.Parse()

	var err error
	game, err = mazemap.NewGame(*col, *row)
	if err != nil {
		log.Fatal(err)
	}

	w, h := game.GetScreenSize()
	if err := ebiten.Run(update, w, h, 1, "迷宫测试"); err != nil {
		log.Fatal(err)
	}

}
