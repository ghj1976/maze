package mazemap

import (
	"container/list"
	"fmt"
	"math/rand"
)

// RandomRune 从给定的字符中随机出一个
func RandomRune(letterRunes []rune) rune {
	return letterRunes[rand.Intn(len(letterRunes))]
}

// BuildRB 使用 递归回溯 算法产生迷宫
func (m *Maze) BuildRB() {
	stack := list.New() // 维护的后进先出队列

	x := 0
	y := 0

	for {

		m.RoomArr[x][y].Visited = true // 当前这个room已经被访问了。

		// 当前位置room，还有那些可选择的方向
		var check []rune
		if x > 0 && !m.RoomArr[x-1][y].Visited {
			check = append(check, 'L')
		}
		if y > 0 && !m.RoomArr[x][y-1].Visited {
			check = append(check, 'U')
		}
		if x < m.ColNum-1 && !m.RoomArr[x+1][y].Visited {
			check = append(check, 'R')
		}
		if y < m.RowNum-1 && !m.RoomArr[x][y+1].Visited {
			check = append(check, 'D')
		}

		if len(check) > 0 {
			stack.PushBack([]int{x, y}) // 还有其他候选方向，压入栈

			switch RandomRune(check) {
			case 'L':
				m.RoomArr[x][y].LeftWall = false
				x--
				m.RoomArr[x][y].RightWall = false
			case 'R':
				m.RoomArr[x][y].RightWall = false
				x++
				m.RoomArr[x][y].LeftWall = false
			case 'U':
				m.RoomArr[x][y].UpWall = false
				y--
				m.RoomArr[x][y].DownWall = false
			case 'D':
				m.RoomArr[x][y].DownWall = false
				y++
				m.RoomArr[x][y].UpWall = false
			default:
				fmt.Println("Default 不应该出现！")
			}

		} else {
			ele := stack.Back()
			val := ele.Value.([]int)
			stack.Remove(ele) //将最后一个元素删除，出栈
			x = val[0]
			y = val[1]
		}

		if stack.Len() == 0 { // 全部处理完
			break
		}
	}
}
