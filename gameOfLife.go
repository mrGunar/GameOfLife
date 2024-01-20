package main

import (
	"fmt"
	"time"
)

const (
	ROWS    = 10
	COLUMNS = 10
)

type Coords struct {
	x int
	y int
}

type Grid struct {
	rows int
	cols int
	grid [][]int
}

func (g *Grid) show() {
	for i := 0; i < g.rows; i++ {
		for j := 0; j < g.cols; j++ {
			fmt.Printf("%d ", g.grid[i][j])
		}
		fmt.Println()
	}
}

func (g *Grid) init() {
	g.grid[1][0] = 1
	g.grid[2][1] = 1
	g.grid[0][2] = 1
	g.grid[1][2] = 1
	g.grid[2][2] = 1
}

func (g *Grid) find(c Coords) (number int) {
	nei := [][]int{{0, 1}, {1, 1}, {-1, 1}, {1, 0}, {-1, 0}, {0, -1}, {-1, -1}, {1, -1}}
	for _, k := range nei {
		x, y := c.x+k[0], c.y+k[1]
		if x >= 0 && x < g.rows && y >= 0 && y < g.cols {
			if g.grid[x][y] == 1 {
				number++
			}
		}
	}
	return
}

func apply_rule(number_of_nei int, value int) (res int) {
	if value == 1 {
		if number_of_nei == 3 || number_of_nei == 2 {
			return 1
		} else {
			return 0
		}
	} else {
		if number_of_nei == 3 {
			return 1
		}
	}
	return
}

func create_grid(N, M int) Grid {
	g := Grid{rows: N, cols: M}
	g.grid = make([][]int, N)
	for i := range g.grid {
		g.grid[i] = make([]int, M)
	}
	return g
}

func next_generation(g *Grid) {
	new_grid := create_grid(ROWS, COLUMNS)
	for i := 0; i < g.rows; i++ {
		for j := 0; j < g.cols; j++ {
			cords := Coords{i, j}
			number_of_nei := g.find(cords)
			new_grid.grid[i][j] = apply_rule(number_of_nei, g.grid[i][j])
		}
	}
	g.grid = new_grid.grid
	return
}

func main() {
	fmt.Println("Conway's Game Of Life!")
	grid := create_grid(ROWS, COLUMNS)
	grid.init()
	grid.show()

	for {
		fmt.Println("======================")
		next_generation(&grid)
		grid.show()
		time.Sleep(time.Second * 2)

	}

}
