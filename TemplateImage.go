package main

import (
	"fmt"
	"image"
	"image/color"
	_ "image/gif"
	_ "image/jpeg"
	_ "image/png"
	"io"
)

func EncodeTemplateImage(image image.Image, output io.Writer) {
	fmt.Println("Encoding template image...")

	imageBounds := image.Bounds()
	imageWidth := imageBounds.Max.X - imageBounds.Min.X
	imageHeight := imageBounds.Max.Y - imageBounds.Min.Y
	pixelCount := imageWidth * imageHeight

	data := make([]byte, (pixelCount+7)/8)
	for y := 0; y < imageHeight; y++ {
		for x := 0; x < imageWidth; x++ {
			pixel := image.At(x, y)
			grayPixel := color.GrayModel.Convert(pixel).(color.Gray)

			if grayPixel.Y < 128 {
				pixelIndex := uint(x + (y * imageWidth))
				byteIndex := uint(pixelIndex / 8)
				bitIndex := 7 - uint(pixelIndex%8)

				data[byteIndex] |= (1 << bitIndex)
			}
		}
	}

	fmt.Fprintf(output, "        DC8 %v, %v\n", imageWidth, imageHeight)
	fmt.Fprint(output, "        DC8 ")

	for i, v := range data {
		if i > 0 {
			fmt.Fprint(output, ", ")
		}

		fmt.Fprintf(output, "0x%02x", v)
	}

	fmt.Println("Encoding completed.")
}
