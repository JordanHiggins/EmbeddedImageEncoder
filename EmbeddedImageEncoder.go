package main

import (
	"flag"
	"fmt"
	"image"
	_ "image/gif"
	_ "image/jpeg"
	_ "image/png"
	"os"
)

func main() {
	bitmapMode := flag.Bool("bitmap", false, "Bitmap Image mode converts the image to a 12-bit format, where each pixel is padded to 16 bits.")
	nativeMode := flag.Bool("native", false, "Native Image mode converts the image to the 12-bit format used by the Nokia 6100 display.")
	templateMode := flag.Bool("template", false, "Template Image mode converts the image to a 1-bit format that can be used to display the image using any color.")
	flag.Parse()

	if flag.NArg() != 2 {
		printUsage()
		return
	}

	imagePath := flag.Arg(0)
	outputPath := flag.Arg(1)

	modeCount := uint(0)

	if *bitmapMode {
		modeCount++
	}
	if *nativeMode {
		modeCount++
	}
	if *templateMode {
		modeCount++
	}

	if modeCount > 1 {
		fmt.Println("Only one mode can be specified at a time.")
		return
	}

	image, imageFormat, err := readImage(imagePath)
	if err != nil {
		fmt.Printf("Failed to load <%v>. %v\n", imagePath, err.Error())
		return
	}

	outputFile, err := os.Create(outputPath)
	if err != nil {
		fmt.Printf("Failed to open <%v> for writing. %v\n", outputPath, err.Error())
		return
	}
	defer outputFile.Close()

	imageBounds := image.Bounds()
	imageWidth := imageBounds.Max.X - imageBounds.Min.X
	imageHeight := imageBounds.Max.Y - imageBounds.Min.Y

	fmt.Printf("Opened %v file <%v> (%v x %v pixel(s)).\n", imageFormat, imagePath, imageWidth, imageHeight)

	if *bitmapMode {
		EncodeBitmapImage(image, outputFile)
	} else if *templateMode {
		EncodeTemplateImage(image, outputFile)
	} else {
		EncodeNativeImage(image, outputFile)
	}
}

func printUsage() {
	fmt.Printf("Usage: %v [option [...]] <infile> <outfile>\n", os.Args[0])
	fmt.Println()
	fmt.Println("Options:")
	flag.VisitAll(func(flag *flag.Flag) {
		fmt.Printf("  -%v (%v): %v\n", flag.Name, flag.DefValue, flag.Usage)
	})
	fmt.Println()
	fmt.Println("Unless otherwise specified, the Native Image mode will be used.")
}

func readImage(path string) (image.Image, string, error) {
	file, err := os.Open(path)
	if err != nil {
		return nil, "", err
	}
	defer file.Close()

	return image.Decode(file)
}
