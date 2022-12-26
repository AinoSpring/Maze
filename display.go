package main

import (
	"image"
	"image/color"
	"image/png"
	"log"
	"os"
)

type Display struct {
	image *image.NRGBA
	size  Vector
}

func DisplayMaze(maze Maze) (display Display) {
	display = NewDisplay(maze.size)
	for x := 0; x < int(maze.size[0]); x++ {
		for y := 0; y < int(maze.size[0]); y++ {
			display.SetPixel(Vector{float64(x), float64(y)}, color.Gray{Y: uint8(maze.Get(Vector{float64(x), float64(y)}) * 255)})
		}
	}
	return
}

func NewDisplay(size Vector) (display Display) {
	display.size = size
	display.image = image.NewNRGBA(image.Rect(0, 0, int(display.size[0]), int(display.size[1])))
	for y := 0; y < int(display.size[0]); y++ {
		for x := 0; x < int(display.size[1]); x++ {
			display.image.Set(x, y, color.NRGBA{
				R: 0,
				G: 0,
				B: 255,
				A: 255,
			})
		}
	}
	return
}

func (display *Display) Save(path string) {
	f, err := os.Create(path)
	if err != nil {
		log.Fatal(err)
	}

	if err := png.Encode(f, display.image); err != nil {
		err := f.Close()
		if err != nil {
			log.Fatal(err)
		}
		log.Fatal(err)
	}

	if err := f.Close(); err != nil {
		log.Fatal(err)
	}
}

func (display *Display) SetPixel(position Vector, color color.Color) {
	display.image.Set(int(position[0]), int(position[1]), color)
}

func (display *Display) Matrix(matrix Matrix, color color.Color) {
	for _, vector := range matrix {
		display.SetPixel(vector, color)
	}
}
