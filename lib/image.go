package lib

type Pixel struct {
	Red   uint8
	Green uint8
	Blue  uint8
	Alpha uint8
}

type SquishImage struct {
	Version int
	Width   int
	Height  int
	Data    [][]Pixel
}

func (sqs *SquishImage) GetPixel(x, y int) Pixel {
	heightLower := max(0, y)
	widthLower := max(0, x)

	if heightLower >= sqs.Height {
		return Pixel{0, 0, 0, 255}
	}

	if widthLower >= sqs.Width {
		return Pixel{0, 0, 0, 255}
	}

	height := min(heightLower, sqs.Height-1)
	width := min(widthLower, sqs.Width-1)

	return sqs.Data[height][width]
}
