package main

import (
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/data/binding"
	"fyne.io/fyne/v2/widget"
	"math"
)

func main() {
	a := app.New()
	w := a.NewWindow("sliderWithData")
	w.Resize(fyne.NewSize(400, 400))

	data := binding.NewFloat()
	slider := widget.NewSliderWithData(0, math.MaxFloat64, data)
	change := widget.NewButton("changed", func() {
		err := data.Set(2.0)
		if err != nil {
			return
		}
	})

	box := container.NewVBox(slider, change)
	w.SetContent(box)
	w.ShowAndRun()

}
