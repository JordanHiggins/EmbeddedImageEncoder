package main

import (
	"flag"
	"fmt"
	"image"
	_ "image/gif"
	_ "image/jpeg"
	_ "image/png"
	"io"
)

var maxLineCount = flag.Uint("maxline", 32000, "The maximum length of a single line of code, in characters. Used in Native Image mode.")

func writeByte(value byte, output io.Writer, lineCount *uint) {
	literal := fmt.Sprintf("0x%02x", value)
	if (*lineCount + 2 + uint(len(literal))) > *maxLineCount {
		*lineCount = 12 + uint(len(literal))
		fmt.Fprintf(output, "\n        DC8 %v", literal)
	} else {
		*lineCount += 2 + uint(len(literal))
		fmt.Fprintf(output, ", %v", literal)
	}
}

func EncodeNativeImage(image image.Image, output io.Writer) {
	fmt.Println("Encoding native image...")

	imageBounds := image.Bounds()
	imageWidth := imageBounds.Max.X - imageBounds.Min.X
	imageHeight := imageBounds.Max.Y - imageBounds.Min.Y
	pixelCount := imageWidth * imageHeight

	fmt.Fprintf(output, "        DC8 %v, %v", imageWidth, imageHeight)

	lineCount := *maxLineCount
	for i := 0; i < pixelCount; i += 2 {
		pixelA := image.At(i%imageWidth, i/imageWidth)
		redA, greenA, blueA, _ := pixelA.RGBA()

		pixelB := image.At((i+1)%imageWidth, (i+1)/imageWidth)
		redB, greenB, blueB, _ := pixelB.RGBA()

		byteA := byte((redA>>12)<<4 | (greenA>>12)<<0)
		writeByte(byteA, output, &lineCount)

		byteB := byte((blueA>>12)<<4 | (redB>>12)<<0)
		writeByte(byteB, output, &lineCount)

		byteC := byte((greenB>>12)<<4 | (blueB>>12)<<0)
		writeByte(byteC, output, &lineCount)
	}

	fmt.Println("Encoding complete.")
}
