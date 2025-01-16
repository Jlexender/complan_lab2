package main

import (
	"Jlexender/complan/internal/image"
	"Jlexender/complan/internal/mapping"
	"fmt"
	"math"
	"math/cmplx"
)

func main() {
	const (
		width, height = 800, 800
		radius        = 3
		limit         = 10
		basePath      = "../../out/"
	)

	set := mapping.GenerateRectangleSet(complex(0, 0), complex(10, math.Pi), 0.1)

	err := image.PlotComplexNumbers(set, basePath+"input.png", radius, limit, limit)
	if err != nil {
		fmt.Println("Error:", err)
	}

	mapping.TransformInPlace(set, func(z complex128) complex128 {
		return cmplx.Exp(complex(real(z), imag(z)/2))
	})

	err = image.PlotComplexNumbers(set, basePath+"output.png", radius, limit, limit)
	if err != nil {
		fmt.Println("Error:", err)
	}

	fmt.Println("Done!")
}
