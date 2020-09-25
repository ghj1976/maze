package mazemap

import (
	"image/color"

	"github.com/hajimehoshi/ebiten"
)

const (
	// WallWidth 墙的厚度
	WallWidth = 1
	// RoomSize 房间的长度大小，假设房间是个正方形, 不包括墙的厚度
	RoomSize = 15
)

// Room 一个具体的房间
type Room struct {
	Col        int  // 列数，第一个是0，数组下标, 房间在总体迷宫的位置。对应x轴
	Row        int  // 行数，第一个是0，数组下标, 房间在总体迷宫的位置。对应y轴
	LeftWall   bool // 房间左侧是否有墙, true 有， false 无
	RightWall  bool // 房间右侧是否有墙, true 有， false 无
	UpWall     bool // 房间上面是否有墙, true 有， false 无
	DownWall   bool // 房间下面是否有墙, true 有， false 无
	Visited    bool // 算法中，是否被访问过，默认false
	IsOpenDoor bool // 是否是入口
	IsExitDoor bool // 是否是出口
}

// NewRoom 构造room
func NewRoom(col, row int) *Room {
	return &Room{Col: col, Row: row}
}

// Draw 画房间
func (r *Room) Draw(mazeImage *ebiten.Image) {
	// 房间要绘画的地点确认
	op := &ebiten.DrawImageOptions{}
	openWallx := 0
	openWallw := 0
	if r.IsOpenDoor { // 画入口
		openWallx = 0 - WallWidth
		openWallw = WallWidth
	}
	xx := WallWidth + (2*WallWidth+RoomSize)*r.Col + WallWidth*getBoolVal(r.LeftWall) + openWallx
	yy := WallWidth + (2*WallWidth+RoomSize)*r.Row + WallWidth*getBoolVal(r.UpWall)
	op.GeoM.Translate(float64(xx), float64(yy))

	// 房间情况分析
	exitWall := 0
	if r.IsExitDoor {
		exitWall = WallWidth
	}
	roomWidth := WallWidth*getBoolVal(!r.LeftWall) + WallWidth*getBoolVal(!r.RightWall) + RoomSize + openWallw + exitWall
	roomHeight := WallWidth*getBoolVal(!r.UpWall) + WallWidth*getBoolVal(!r.DownWall) + RoomSize
	roomImage, _ := ebiten.NewImage(roomWidth, roomHeight, ebiten.FilterDefault)
	roomImage.Fill(color.White)

	// 绘制房间
	mazeImage.DrawImage(roomImage, op)
}

func getBoolVal(b bool) int {
	if b {
		return 1
	}
	return 0
}
