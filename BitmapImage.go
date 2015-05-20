package main

import (
	"fmt"
	"image"
	_ "image/gif"
	_ "image/jpeg"
	_ "image/png"
	"io"
)

func EncodeBitmapImage(image image.Image, output io.Writer) {
	fmt.Println("Encoding bitmap image...")

	imageBounds := image.Bounds()
	imageWidth := imageBounds.Max.X - imageBounds.Min.X
	imageHeight := imageBounds.Max.Y - imageBounds.Min.Y

	fmt.Fprintf(output, "        DC8 %v, %v", imageWidth, imageHeight)

	for y := 0; y < imageHeight; y++ {
		fmt.Fprint(output, "\n        DC16 ")
		for x := 0; x < imageWidth; x++ {
			if x > 0 {
				fmt.Fprint(output, ", ")
			}

			pixel := image.At(x, y)
			red, green, blue, _ := pixel.RGBA()

			fmt.Fprintf(output, "0x%x%x%x", red>>12, green>>12, blue>>12)
		}
	}

	fmt.Println("Encoding complete.")
}
