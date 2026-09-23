package main

import (
	"fmt"
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/canvas"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/data/binding"
	"fyne.io/fyne/v2/widget"
	"github.com/labstack/gommon/log"
	"image/color"
)

func main() {
	a := app.New()
	w := a.NewWindow("text")
	w.Resize(fyne.NewSize(400, 400))

	text := canvas.NewText("text", color.White)
	text.Alignment = fyne.TextAlignCenter

	data := binding.NewFloat()
	data.Set(100.0)
	data.AddListener(binding.NewDataListener(func() {
		d, _ := data.Get()
		t := fmt.Sprintf("%v", d)
		text.Text = t
	}))

	change := widget.NewButton("changed", func() {
		err := data.Set(2.0)
		if err != nil {
			log.Error(err.Error())
		}
	})

	c := container.NewVBox(text, change)
	w.SetContent(c)
	w.ShowAndRun()
}
