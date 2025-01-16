package main

import (
	"Jlexender/complan/internal/image"
	"Jlexender/complan/internal/mapping"
	"fmt"
	"math"
	"math/cmplx"
)

func main() {

	fmt.Println("Hello, world!")

	set := mapping.GenerateRectangleSet(complex(0, 0), complex(10, math.Pi), 0.1)


	err := image.PlotComplexNumbers(set, "input.png", 5, 5, 5)

	if err != nil {
		fmt.Println("Error:", err)
	}

	mapping.TransformInPlace(set, func(z complex128) complex128 {
		return cmplx.Exp(complex(real(z), imag(z)/2))
	})

	err = image.PlotComplexNumbers(set, "output.png", 5, 5, 5)

	if err != nil {
		fmt.Println("Error:", err)
	}

	fmt.Println("Done!")
}