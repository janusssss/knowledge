package main

import (
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/layout"
	"fyne.io/fyne/v2/widget"
	"strconv"
)

func main() {
	a := app.New()
	w := a.NewWindow("vbox")
	w.Resize(fyne.NewSize(400, 400))

	objects := make([]fyne.CanvasObject, 6)
	for i := range len(objects) {
		objects[i] = widget.NewLabel(strconv.Itoa(i))
	}

	content := container.NewVBox(objects...)

	layout.NewVBoxLayout()
	w.SetContent(content)
	w.ShowAndRun()
}
