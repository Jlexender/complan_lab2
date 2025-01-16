package main

import (
	"fmt"
	"Jlexender/complan/internal/mapping"
	"Jlexender/complan/internal/image"
)

func main() {
	fmt.Println("Hello, world!")

	set := mapping.GenerateRectangleSet(complex(-1, -1), complex(1, 1), 0.1)

	err := image.PlotComplexNumbers(set, "input.png")

	if err != nil {
		fmt.Println("Error:", err)
	}

	fmt.Println("Done!")
	
}