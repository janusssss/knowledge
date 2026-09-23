package main

import (
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/widget"
)

func main() {
	a := app.New()
	w := a.NewWindow("Grib")
	w.Resize(fyne.NewSize(400, 400))

	sample := []string{"A", "B", "C", "D", "E", "F"}
	objects := make([]fyne.CanvasObject, len(sample))
	for i, s := range sample {
		objects[i] = widget.NewLabel(s)
	}
	content := container.NewGridWithColumns(3, objects...)

	w.SetContent(container.NewVBox(content))
	w.ShowAndRun()
}
