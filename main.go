package main

import (
	"math/rand"
	"time"
)

func main() {
	rand.Seed(time.Now().UnixNano())
	var maze = NewMaze(NewVector(2).Fill(50))
	var display = DisplayMaze(maze)
	display.Save("maze.png")
}
