package main

import (
	"log"
	"math"
)

// VECTORS

type Vector []float64

func NewVector(size interface{}) Vector {
	var floatSize float64
	switch size.(type) {
	case float64:
		floatSize = size.(float64)
	case float32:
		floatSize = float64(size.(float32))
	case int:
		floatSize = float64(size.(int))
	case uint:
		floatSize = float64(size.(uint))
	default:
		log.Panicf("Size has invalid type: %T\n", size)
	}
	return make(Vector, int(floatSize))
}

func (vector Vector) Size() float64 {
	return float64(len(vector))
}

func (vector Vector) Copy() (copiedVector Vector) {
	copiedVector = NewVector(vector.Size())
	for idx, element := range vector {
		copiedVector[idx] = element
	}
	return
}

func (vector Vector) Fill(value float64) (filledVector Vector) {
	filledVector = NewVector(vector.Size())
	for idx, _ := range vector {
		filledVector[idx] = value
	}
	return
}

func (vector Vector) Calculate(other interface{}, operator func(float64, float64) float64) (calculatedVector Vector) {
	calculatedVector = vector.Copy()
	switch other.(type) {
	case Vector:
		for idx, element := range other.(Vector) {
			calculatedVector[idx] = operator(calculatedVector[idx], element)
		}
	case float64:
		for idx := range vector {
			calculatedVector[idx] = operator(calculatedVector[idx], other.(float64))
		}
	case float32:
		for idx := range vector {
			calculatedVector[idx] = operator(calculatedVector[idx], float64(other.(float32)))
		}
	case int:
		for idx := range vector {
			calculatedVector[idx] = operator(calculatedVector[idx], float64(other.(int)))
		}
	default:
		log.Panicf("Other value has invalid type: %T\n", other)
	}
	return
}

func (vector Vector) Add(other interface{}) Vector {
	return vector.Calculate(other, func(a float64, b float64) float64 {
		return a + b
	})
}

func (vector Vector) Sub(other interface{}) Vector {
	return vector.Calculate(other, func(a float64, b float64) float64 {
		return a - b
	})
}

func (vector Vector) Mul(other interface{}) Vector {
	return vector.Calculate(other, func(a float64, b float64) float64 {
		return a * b
	})
}

func (vector Vector) Div(other interface{}) Vector {
	return vector.Calculate(other, func(a float64, b float64) float64 {
		return a / b
	})
}

func (vector Vector) Pow(other interface{}) Vector {
	return vector.Calculate(other, func(a float64, b float64) float64 {
		return math.Pow(a, b)
	})
}

func (vector Vector) Mod(other interface{}) Vector {
	return vector.Calculate(other, func(a float64, b float64) float64 {
		return math.Mod(a, b)
	})
}

func (vector Vector) Dot(other Vector) float64 {
	return vector.Mul(other).Sum()
}

func (vector Vector) Sum() (sum float64) {
	for _, element := range vector {
		sum += element
	}
	return
}

func (vector Vector) Average() float64 {
	return vector.Sum() / vector.Size()
}

func (vector Vector) Magnitude() float64 {
	return math.Sqrt(vector.Pow(2).Sum())
}

func (vector Vector) Unit() Vector {
	return vector.Div(vector.Magnitude())
}

func (vector Vector) Lerp(other Vector, t float64) (result Vector) {
	result = NewVector(vector.Size())
	for idx, element := range vector {
		result[idx] = element + ((other[idx] - element) * t)
	}
	return
}

func (vector Vector) Equals(other Vector) bool {
	if vector.Size() != other.Size() {
		return false
	}
	for idx, element := range vector {
		if element != other[idx] {
			return false
		}
	}
	return true
}

func (vector Vector) Floor() (result Vector) {
	result = NewVector(vector.Size())
	for idx, element := range vector {
		result[idx] = float64(int(element))
	}
	return
}

// MATRICES

type Matrix []Vector

func NewMatrix(size Vector) (matrix Matrix) {
	matrix = make(Matrix, int(size[1]))
	for idx, _ := range matrix {
		matrix[idx] = NewVector(size[0])
	}
	return
}

func (matrix Matrix) Size() Vector {
	if len(matrix) == 0 {
		return NewVector(2)
	}
	return Vector{matrix[0].Size(), float64(len(matrix))}
}

func (matrix Matrix) Copy() (copiedMatrix Matrix) {
	copiedMatrix = NewMatrix(matrix.Size())
	for idx, vector := range matrix {
		copiedMatrix[idx] = vector.Copy()
	}
	return
}

func (matrix Matrix) Fill(value float64) (filledMatrix Matrix) {
	filledMatrix = NewMatrix(matrix.Size())
	for idx, _ := range matrix {
		filledMatrix[idx] = filledMatrix[idx].Fill(value)
	}
	return
}

func (matrix Matrix) Calculate(other interface{}, operator func(float64, float64) float64) (calculatedMatrix Matrix) {
	calculatedMatrix = matrix.Copy()
	switch other.(type) {
	case Matrix:
		for idx, vector := range other.(Matrix) {
			calculatedMatrix[idx] = calculatedMatrix[idx].Calculate(vector, operator)
		}
	default:
		for idx, _ := range matrix {
			calculatedMatrix[idx] = calculatedMatrix[idx].Calculate(other, operator)
		}
	}
	return
}

func (matrix Matrix) Add(other interface{}) Matrix {
	return matrix.Calculate(other, func(a float64, b float64) float64 {
		return a + b
	})
}

func (matrix Matrix) Sub(other interface{}) Matrix {
	return matrix.Calculate(other, func(a float64, b float64) float64 {
		return a - b
	})
}

func (matrix Matrix) Mul(other interface{}) Matrix {
	return matrix.Calculate(other, func(a float64, b float64) float64 {
		return a * b
	})
}

func (matrix Matrix) Div(other interface{}) Matrix {
	return matrix.Calculate(other, func(a float64, b float64) float64 {
		return a / b
	})
}

func (matrix Matrix) Pow(other interface{}) Matrix {
	return matrix.Calculate(other, func(a float64, b float64) float64 {
		return math.Pow(a, b)
	})
}

func (matrix Matrix) Mod(other interface{}) Matrix {
	return matrix.Calculate(other, func(a float64, b float64) float64 {
		return math.Mod(a, b)
	})
}

func (matrix Matrix) Multiplication(other Vector) Vector {
	var multipliedMatrix = NewMatrix(matrix.Size())
	for idx, vector := range matrix {
		multipliedMatrix[idx] = vector.Mul(other)
	}
	return multipliedMatrix.Sum()
}

func (matrix Matrix) Sum() (sum Vector) {
	sum = NewVector(matrix.Size()[0])
	for _, vector := range matrix {
		sum = sum.Add(vector)
	}
	return
}

func (matrix Matrix) Average() (average Vector) {
	average = NewVector(matrix.Size()[0])
	for _, vector := range matrix {
		average = average.Add(vector)
	}
	if matrix.Size()[1] > 0 {
		average = average.Div(matrix.Size()[1])
	}
	return
}
