package main

import (
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/widget"
)

func main() {
	a := app.New()
	w := a.NewWindow("Slider")
	w.Resize(fyne.NewSize(300, 300))

	slider := widget.NewSlider(0, 300)

	w.SetContent(slider)
	w.ShowAndRun()
}
