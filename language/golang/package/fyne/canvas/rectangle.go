package main

import (
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/canvas"
	"image/color"
)

func main() {
	a := app.New()
	w := a.NewWindow("Rectangle")
	w.Resize(fyne.NewSize(400, 400))

	rectangle := canvas.NewRectangle(color.White)
	rectangle.CornerRadius = 90
	w.SetContent(rectangle)

	w.ShowAndRun()
}
