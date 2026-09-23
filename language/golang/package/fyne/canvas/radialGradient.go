package main

import (
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/canvas"
	"image/color"
)

func main() {
	a := app.New()
	w := a.NewWindow("Radial Gradient")
	w.Resize(fyne.NewSize(400, 400))

	//gradient := canvas.NewRadialGradient(color.White, color.Transparent)
	gradient := canvas.NewLinearGradient(color.White, color.Black, 45)

	w.SetContent(gradient)

	w.ShowAndRun()
}
