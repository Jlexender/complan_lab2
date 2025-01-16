package image

import (
    "image"
    "image/color"
    "image/png"
    "os"
)

// PlotComplexNumbers plots the given complex numbers on a complex plane.
func PlotComplexNumbers(numbers []complex128, filename string, radius int, maxX, maxY float64) error {
    const width, height = 1600, 1600
    img := image.NewRGBA(image.Rect(0, 0, width, height))
    black := color.RGBA{0, 0, 0, 255}
    white := color.RGBA{255, 255, 255, 255}

    for x := 0; x < width; x++ {
        for y := 0; y < height; y++ {
            img.Set(x, y, white)
        }
    }

    generateAxis(img, width, height)

    for _, num := range numbers {
        x := int((real(num)/maxX + 1) * float64(width) / 2)
        y := int((1 - imag(num)/maxY) * float64(height) / 2)
        drawCircle(img, x, y, radius, black)
    }

    file, err := os.Create(filename)
	if err != nil {
		return err
    }
    defer file.Close()

    return png.Encode(file, img)
}

func generateAxis(img *image.RGBA, width, height int) {
    black := color.RGBA{0, 0, 0, 255}
    for x := 0; x < width; x++ {
        img.Set(x, height/2, black)
    }
    for y := 0; y < height; y++ {
        img.Set(width/2, y, black)
    }
}

func drawCircle(img *image.RGBA, x0, y0, radius int, col color.Color) {
    for y := -radius; y <= radius; y++ {
        for x := -radius; x <= radius; x++ {
            if x*x+y*y <= radius*radius {
                img.Set(x0+x, y0+y, col)
            }
        }
    }
}