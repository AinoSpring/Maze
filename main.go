package main

import (
	"image/color"
	"math/rand"
	"time"
)

func main() {
	rand.Seed(time.Now().UnixNano())
	var maze = NewMaze(NewVector(2).Fill(50))
	var display = DisplayMaze(maze)
	var solvedPath = Solve(maze, maze.start, NewVector(0))
	display.Matrix(solvedPath, color.RGBA{
		R: 0,
		G: 0,
		B: 255,
		A: 255,
	})
	display.SetPixel(maze.start, color.RGBA{
		R: 255,
		G: 0,
		B: 0,
		A: 255,
	})
	display.SetPixel(maze.finish, color.RGBA{
		R: 0,
		G: 255,
		B: 0,
		A: 255,
	})
	display.Save("maze.png")
}
