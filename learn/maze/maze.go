package main

import (
	"fmt"
	"os"
)

func readMaze(fileName string) [][]int {
	file, err := os.Open(fileName)
	if err != nil {
		panic(err)
	}
	defer file.Close()
	var row, col int
	fmt.Fscanf(file, "%d %d ", &row, &col)

	maze := make([][]int, row)

	// for range 的使用
	// 当range后面接的是数组时，返回一个参数是index，返回两个参数是index，value
	// 当range后面接的是map时， 返回一个参数是key，  返回两个参数是key，value
	for i := range maze {
		maze[i] = make([]int, col)
		for j := range maze[i] {
			fmt.Fscanf(file, "%d", &maze[i][j])
		}
	}
	return maze
}

type point struct {
	i, j int
}

func (p point) add(po point) point {
	return point{p.i + po.i, p.j + po.j}
}

func (p point) at(grid [][]int) (int, bool) {
	if p.i < 0 || p.i >= len(grid) {
		return 0, false
	}
	if p.j < 0 || p.j >= len(grid[p.i]) {
		return 0, false
	}
	return grid[p.i][p.j], true

}

var dirs = []point{{-1, 0}, {0, -1}, {1, 0}, {0, 1}}

func walk(maze [][]int, start, end point) [][]int {
	steps := make([][]int, len(maze))
	for i := range steps {
		steps[i] = make([]int, len(maze[0]))
	}
	Q := []point{start}
	for len(Q) > 0 {
		cur := Q[0]
		Q = Q[1:]
		if cur == end {
			break
		}
		for _, dir := range dirs {
			next := cur.add(dir)
			val, ok := next.at(maze)
			if !ok || val == 1 {
				continue
			}
			val, ok = next.at(steps)
			if !ok || val != 0 {
				continue
			}
			if next == start {
				continue
			}
			curSteps, _ := cur.at(steps)
			steps[next.i][next.j] = curSteps + 1
			Q = append(Q, next)
		}
	}
	return steps
}

func main() {
	maze := readMaze("maze/maze.in")
	steps := walk(maze, point{0, 0}, point{len(maze), len(maze[0])})

	for _, v := range steps {
		for _, j := range v {
			fmt.Printf("%4d ", j)
		}
		fmt.Println()
	}
}
