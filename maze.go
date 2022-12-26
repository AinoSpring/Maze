package main

import (
	"math"
	"math/rand"
	"strconv"
)

const (
	PATH float64 = 0
	WALL float64 = 1
)

type Maze struct {
	grid   Matrix
	size   Vector
	start  Vector
	finish Vector
}

func NewMaze(size Vector) (maze Maze) {
	maze.grid = NewMatrix(size).Fill(WALL)
	maze.size = size
	maze.start = maze.RandomBorderPosition()
	maze.Generate(maze.start, -1)
	maze.CalculateFinish()
	return
}

func (maze *Maze) Get(position Vector) float64 {
	return maze.grid[int(position[0])][int(position[1])]
}

func (maze *Maze) Set(position Vector, value float64) {
	maze.grid[int(position[0])][int(position[1])] = value
}

func (maze *Maze) RandomBorderPosition() Vector {
	var randomPosition = Vector{rand.Float64(), rand.Float64()}
	var axis = rand.Intn(2)
	randomPosition[axis] = math.Round(randomPosition[axis]) * (maze.size[axis] - 1)
	randomPosition[1-axis] = randomPosition[1-axis] * (maze.size[axis] - 1)
	return randomPosition.Floor()
}

func (maze *Maze) InBounds(position Vector) bool {
	return position[0] >= 0 && position[1] >= 0 && position[0] < maze.size[0] && position[1] < maze.size[1]
}

func (maze *Maze) AvailableDirections(position Vector) (availableDirections Matrix) {
	availableDirections = make(Matrix, 0)
	for x := -1; x < 2; x += 2 {
		var direction = Vector{float64(x), 0}
		if maze.InBounds(position.Add(direction)) {
			availableDirections = append(availableDirections, direction)
		}
	}
	for y := -1; y < 2; y += 2 {
		var direction = Vector{0, float64(y)}
		if maze.InBounds(position.Add(direction)) {
			availableDirections = append(availableDirections, direction)
		}
	}
	return
}

func (maze *Maze) CountBounds(position Vector) (count int) {
	for _, direction := range maze.AvailableDirections(position) {
		if addedPosition := position.Add(direction); !maze.InBounds(addedPosition) || maze.Get(addedPosition) == PATH {
			count++
		}
	}
	return count
}

func (maze *Maze) NotBoundingDirections(position Vector, toleratedBunds int) (notBoundingDirections Matrix) {
	notBoundingDirections = make(Matrix, 0)
	for _, direction := range maze.AvailableDirections(position) {
		if !maze.InBounds(position.Add(direction)) {
			continue
		}
		if maze.CountBounds(position.Add(direction)) <= toleratedBunds {
			notBoundingDirections = append(notBoundingDirections, direction)
		}
	}
	return
}

func (maze *Maze) Generate(position Vector, i int) {
	if i == 0 {
		return
	}
	maze.Set(position, PATH)
	var notBoundingDirections = maze.NotBoundingDirections(position, 1)
	if notBoundingDirections.Size()[1] == 0 {
		return
	}
	rand.Shuffle(int(notBoundingDirections.Size()[1]), func(i int, j int) {
		notBoundingDirections[i], notBoundingDirections[j] = notBoundingDirections[j], notBoundingDirections[i]
	})
	for _, direction := range notBoundingDirections {
		if addedDirection := position.Add(direction); maze.CountBounds(addedDirection) == 1 && maze.Get(addedDirection) != PATH {
			maze.Generate(position.Add(direction), i-1)
		}
	}
}

func (maze *Maze) CalculateFinish() {
	var max = Vector{-1}
	for x, vector := range maze.grid {
		for y, value := range vector {
			if value != PATH {
				continue
			}
			var currentPosition = Vector{float64(x), float64(y)}
			if max[0] == -1 {
				max = currentPosition
				continue
			}
			if max.Sub(maze.start).Magnitude() < currentPosition.Sub(maze.start).Magnitude() {
				max = currentPosition
			}
		}
	}
	maze.finish = max
}

func (maze *Maze) String() (mazeString string) {
	for x := 0; x < int(maze.size[0]); x++ {
		for y := 0; y < int(maze.size[0]); y++ {
			mazeString += strconv.Itoa(int(maze.Get(Vector{float64(x), float64(y)})))
		}
		mazeString += "\n"
	}
	return
}
