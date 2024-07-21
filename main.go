package main

import (
	"fmt"
	"image"
	"image/png"
	"io"
	"os"

	"github.com/atbuy/squisher/lib"
	"github.com/atbuy/squisher/sqsio"
	"github.com/atbuy/squisher/window"
)

func main() {
	image.RegisterFormat("png", "png", png.Decode, png.DecodeConfig)

	// Read pixels from input image
	pixels, width, height := loadImageData("lena.png")

	// Create output image file
	fileOut, err := os.Create("image.sqs")
	if err != nil {
		panic(err)
	}

	defer fileOut.Close()

	// Create squish image struct to display
	image := lib.SquishImage{
		Version: 1,
		Width:   width,
		Height:  height,
		Data:    pixels,
	}

	window.DisplayImage(image)

	metadata := sqsio.ToBin(1, 1, 1)
	data := sqsio.ImageToBin(image)
	imageData := sqsio.JoinStrBin(metadata, data)

	// Compress the image data using Run Length Encoding
	compressed1 := lib.RLECompression(imageData)
	runLength := sqsio.StrToBytes(compressed1)

	// byteData := sqsio.BinToBytes(imageData)

	if _, err := fileOut.Write(runLength); err != nil {
		panic(err)
	}
}

func loadImageData(path string) ([][]lib.Pixel, int, int) {
	file, err := os.Open(path)
	if err != nil {
		fmt.Println("Error: Image could not be opened")
		os.Exit(1)
	}
	defer file.Close()

	pixels, width, height, err := getPixels(file)
	if err != nil {
		fmt.Println("Error: Image could not be decoded")
		os.Exit(1)
	}

	return pixels, width, height
}

func getPixels(file io.Reader) ([][]lib.Pixel, int, int, error) {
	img, _, err := image.Decode(file)
	if err != nil {
		return nil, 0, 0, err
	}

	bounds := img.Bounds()
	width, height := bounds.Max.X, bounds.Max.Y

	var pixels [][]lib.Pixel
	for y := 0; y < height; y++ {
		var row []lib.Pixel
		for x := 0; x < width; x++ {
			row = append(row, rgbaToPixel(img.At(x, y).RGBA()))
		}

		pixels = append(pixels, row)
	}

	return pixels, width, height, nil
}

func rgbaToPixel(r uint32, g uint32, b uint32, a uint32) lib.Pixel {
	return lib.Pixel{
		Red:   uint8(r / 257),
		Green: uint8(g / 257),
		Blue:  uint8(b / 257),
		Alpha: uint8(a / 257),
	}
}
