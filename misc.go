package main

func MatrixIncludes(matrix Matrix, other Vector) bool {
	for _, vector := range matrix {
		if vector.Equals(other) {
			return true
		}
	}
	return false
}
