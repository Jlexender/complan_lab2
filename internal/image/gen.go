package image

import (
	"image"
    "image/color"
	"image/png"
	"os"
)

// PlotComplexNumbers plots the given complex numbers on a complex plane.
func PlotComplexNumbers(numbers []complex128, filename string) error {
	const width, height = 800, 800
	img := image.NewRGBA(image.Rect(0, 0, width, height))
	black := color.RGBA{0, 0, 0, 255}
	white := color.RGBA{255, 255, 255, 255}

	// Fill the image with white color
	for x := 0; x < width; x++ {
		for y := 0; y < height; y++ {
			img.Set(x, y, white)
		}
	}

	// Plot each complex number
	for _, num := range numbers {
		x := int(real(num)*float64(width)/2 + float64(width)/2)
		y := int(imag(num)*float64(height)/2 + float64(height)/2)
		if x >= 0 && x < width && y >= 0 && y < height {
			img.Set(x, y, black)
		}
	}

	// Save the image to a file
	file, err := os.Create(filename)
	if err != nil {
		return err
	}
	defer file.Close()

	err = png.Encode(file, img)
	if err != nil {
		return err
	}

	return nil
}
