package main

import (
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/widget"
)

func main() {
	a := app.New()
	w := a.NewWindow("Table")
	w.Resize(fyne.NewSize(400, 400))

	sample := [][]string{
		{"1", "2", "3"},
		{"A", "B", "C"},
	}
	table := widget.NewTable(
		func() (rows int, cols int) {
			return len(sample), len(sample[0])
		},
		func() fyne.CanvasObject {
			return widget.NewLabel("")
		},
		func(id widget.TableCellID, object fyne.CanvasObject) {
			object.(*widget.Label).SetText(sample[id.Row][id.Col])
		},
	)

	w.SetContent(table)
	w.ShowAndRun()
}
