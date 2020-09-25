package mazemap

import (
	"image/color"

	"github.com/hajimehoshi/ebiten"
)

// Maze 迷宫类
type Maze struct {
	ColNum  int // 迷宫列数，对应X轴
	RowNum  int // 迷宫行数, 对应Y轴
	RoomArr [][]*Room
}

// NewMaze 新建一个迷宫类
func NewMaze(col, row int) (*Maze, error) {

	m := &Maze{ColNum: col, RowNum: row}

	// 分配房间
	for i := 0; i < col; i++ {
		ri := make([]*Room, 0, col)
		for j := 0; j < row; j++ {
			// RoomS[i][j] =
			rj := NewRoom(i, j)
			rj.LeftWall = true
			rj.RightWall = true
			rj.UpWall = true
			rj.DownWall = true
			ri = append(ri, rj)
		}
		m.RoomArr = append(m.RoomArr, ri)
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
		mazeImage.Fill(color.RGBA{34, 49, 255, 255})
		//color.RGBA{0x00, 0x00, 0xff, 0xff})
		//color.RGBA{34, 49, 63, 255})
		//

	}

	op := &ebiten.DrawImageOptions{}
	sw, sh := screen.Size()
	bw, bh := mazeImage.Size()
	x := (sw - bw) / 2
	y := (sh - bh) / 2
	op.GeoM.Translate(float64(x), float64(y))

	for i := 0; i < m.ColNum; i++ {
		for j := 0; j < m.RowNum; j++ {
			m.RoomArr[i][j].Draw(mazeImage)
		}
	}
	screen.DrawImage(mazeImage, op)

}

// Size 迷宫的大小
func (m *Maze) Size() (int, int) {
	x := (2*WallWidth+RoomSize)*m.ColNum + 2*WallWidth
	y := (2*WallWidth+RoomSize)*m.RowNum + 2*WallWidth
	return x, y
}

// OpenDoor 打开入口和出口
func (m *Maze) OpenDoor() {
	m.RoomArr[0][0].LeftWall = false
	m.RoomArr[0][0].IsOpenDoor = true
	m.RoomArr[m.ColNum-1][m.RowNum-1].RightWall = false
	m.RoomArr[m.ColNum-1][m.RowNum-1].IsExitDoor = true
}
