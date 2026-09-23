package main

import (
	"fmt"
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/canvas"
	"image/color"
)

func main() {
	a := app.New()
	w := a.NewWindow("plot")
	w.Resize(fyne.NewSize(400, 400))

	circle := canvas.NewCircle(color.White)

	position := circle.Position()
	fmt.Println(position)
	circle.Position1 = fyne.Position{X: 50, Y: 50}
	circle.Position2 = fyne.Position{X: 100, Y: 100}

	//circle.Resize(fyne.NewSize(400, 400))

	w.SetContent(circle)
	w.ShowAndRun()
}
