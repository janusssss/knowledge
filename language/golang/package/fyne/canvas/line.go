package main

import (
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/canvas"
	"fyne.io/fyne/v2/container"
	"image/color"
)

func main() {
	a := app.New()
	w := a.NewWindow("Draw Line Example")

	line := canvas.NewLine(color.White)
	container.NewWithoutLayout()

	line.Position1 = fyne.NewPos(10, 50)
	line.Position2 = fyne.NewPos(390, 50)

	//w.SetContent(line)
	w.SetContent(container.NewGridWrap(fyne.NewSize(400, 100), line))

	w.Resize(fyne.NewSize(400, 100))
	w.ShowAndRun()
}
