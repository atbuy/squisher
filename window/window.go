package window

import (
	"image/color"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/canvas"

	"github.com/atbuy/squisher/lib"
)

func DisplayImage(sqsImage lib.SquishImage) {
	a := app.New()
	w := a.NewWindow("Image")

	img := canvas.NewRasterWithPixels(
		func(x, y, _, _ int) color.Color {
			pixel := sqsImage.GetPixel(x, y)
			return color.RGBA{
				R: pixel.Red,
				G: pixel.Green,
				B: pixel.Blue,
				A: pixel.Alpha,
			}
		},
	)

	w.SetContent(img)
	w.Resize(fyne.NewSize(float32(sqsImage.Width), float32(sqsImage.Height)))
	w.ShowAndRun()
}
