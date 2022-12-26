package main

func GetPathDirections(maze Maze, position Vector) (directions Matrix) {
	directions = make(Matrix, 0)
	for x := -1; x < 2; x += 2 {
		var direction = Vector{float64(x), 0}
		if addedPosition := position.Add(direction); maze.InBounds(addedPosition) && maze.Get(addedPosition) == PATH {
			directions = append(directions, direction)
		}
	}
	for y := -1; y < 2; y += 2 {
		var direction = Vector{0, float64(y)}
		if addedPosition := position.Add(direction); maze.InBounds(addedPosition) && maze.Get(addedPosition) == PATH {
			directions = append(directions, direction)
		}
	}
	return
}

func Solve(maze Maze, position Vector, lastPosition Vector) Matrix {
	if position.Equals(maze.finish) {
		return Matrix{position}
	}
	for _, direction := range GetPathDirections(maze, position) {
		if addedPosition := position.Add(direction); !lastPosition.Equals(addedPosition) {
			if solvedPath := Solve(maze, addedPosition, position); solvedPath.Size()[1] != 0 {
				return append(solvedPath, position)
			}
		}
	}
	return Matrix{}
}
