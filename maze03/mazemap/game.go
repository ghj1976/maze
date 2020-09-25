package mazemap

import (
	"image/color"
	"math/rand"
	"time"

	"github.com/hajimehoshi/ebiten"
)

func init() {
	rand.Seed(time.Now().UnixNano())
}

// Game 游戏逻辑
type Game struct {
	maze *Maze
}

// NewGame generates a new Game object.
func NewGame(col, row int) (*Game, error) {
	g := &Game{}

	// 默认初始化 全部墙存在
	g.maze, _ = NewMaze(col, row)

	// 算法，计算那些墙打开。
	g.maze.BuildRB()

	// 开进口和出口
	g.maze.OpenDoor()

	return g, nil
}

// GetScreenSize 得到游戏区域 屏幕得大小
func (g *Game) GetScreenSize() (int, int) {
	if g == nil {
		return 420, 600
	}
	w, h := g.maze.Size()
	return w + 80, h + 60
}

// Draw 绘图
func (g *Game) Draw(screen *ebiten.Image) {
	screen.Fill(color.RGBA{0xfa, 0xf8, 0xef, 0xff})

	g.maze.Draw(screen)
}
