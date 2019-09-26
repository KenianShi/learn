package main

import (
	"fmt"
	"os"
)

func readMaze(path string) [][]int {
	file, err := os.Open(path)
	if err != nil {
		panic(err)
	}
	defer file.Close()
	var row, col int
	fmt.Fscanf(file, "%d %d", &row, &col)
	maze := make([][]int, row)
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
	if p.j < 0 || p.j >= len(grid[0]) {
		return 0, false
	}
	return grid[p.i][p.j], true

}

var dirs []point = []point{{-1, 0}, {0, -1}, {1, 0}, {0, 1}}

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
			if !ok || val != 0 {
				continue
			}
			val, ok = next.at(steps)
			if !ok || val != 0 {
				continue
			}
			if next == start {
				continue
			}

			curStep, _ := cur.at(steps)

			steps[next.i][next.j] = curStep + 1
			Q = append(Q, next)
		}
	}
	return steps
}

func main() {
	mazes := readMaze("maze3/maze.in")
	for i := range mazes {
		for j := range mazes[i] {
			fmt.Printf("%4d", mazes[i][j])
		}
		fmt.Println()
	}
	fmt.Println("*******************************************************")
	steps := walk(mazes, point{0, 0}, point{len(mazes) - 1, len(mazes[0]) - 1})
	for i := range steps {
		for j := range steps[i] {
			fmt.Printf("%4d", steps[i][j])
		}
		fmt.Println()
	}

}
