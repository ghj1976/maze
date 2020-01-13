package main

import (
	"image/color"

	"github.com/hajimehoshi/ebiten"
)

// Maze 迷宫类
type Maze struct {
	width  int
	height int
	RoomS  [][]*Room
}

// NewMaze 新建一个迷宫类
func NewMaze(width, height int) (*Maze, error) {

	m := &Maze{width: width, height: height}

	// 分配房间
	for i := 0; i < width; i++ {
		ri := make([]*Room, 0, width)
		for j := 0; j < height; j++ {
			// RoomS[i][j] =
			rj := NewRoom(i, j)
			ri = append(ri, rj)
		}
		m.RoomS = append(m.RoomS, ri)
	}

	return m, nil
}

var (
	mazeImage *ebiten.Image
)

// Draw 画房间
func (m *Maze) Draw(screen *ebiten.Image) {

	if mazeImage == nil {
		w, h := m.Size()
		mazeImage, _ = ebiten.NewImage(w, h, ebiten.FilterDefault)
		mazeImage.Fill(color.RGBA{0xfa, 0xd8, 0xef, 0xff})
	}

	op := &ebiten.DrawImageOptions{}
	sw, sh := screen.Size()
	bw, bh := mazeImage.Size()
	x := (sw - bw) / 2
	y := (sh - bh) / 2
	op.GeoM.Translate(float64(x), float64(y))

	for i := 0; i < m.width; i++ {
		for j := 0; j < m.height; j++ {
			m.RoomS[i][j].Draw(mazeImage)
		}
	}
	screen.DrawImage(mazeImage, op)

}

// Size 迷宫的大小
func (m *Maze) Size() (int, int) {
	x := (WallWidth+1)*m.width + RoomSize*m.width
	y := (WallWidth+1)*m.height + RoomSize*m.height
	return x, y
}
