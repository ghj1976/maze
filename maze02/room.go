package main

import (
	"image/color"

	"github.com/hajimehoshi/ebiten"
)

const (
	// WallWidth 墙的厚度
	WallWidth = 4
	// RoomSize 房间的长度大小，假设房间是个正方形
	RoomSize = 60
)

// Room 一个具体的房间
type Room struct {
	x int
	y int
}

// NewRoom 构造room
func NewRoom(x, y int) *Room {
	return &Room{x: x, y: y}
}

var (
	roomImage *ebiten.Image
)

func init() {
	roomImage, _ = ebiten.NewImage(RoomSize, RoomSize, ebiten.FilterDefault)
	roomImage.Fill(color.White)
}

// Draw 画房间
func (r *Room) Draw(mazeImage *ebiten.Image) {
	op := &ebiten.DrawImageOptions{}
	xx := WallWidth + (WallWidth+RoomSize)*r.x
	yy := WallWidth + (WallWidth+RoomSize)*r.y
	op.GeoM.Translate(float64(xx), float64(yy))
	mazeImage.DrawImage(roomImage, op)
}
