package main

import (
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/widget"
)

func main() {
	a := app.New()
	w := a.NewWindow("")
	w.Resize(fyne.NewSize(800, 600))

	text := "hello world"
	label := widget.NewLabel(text)
	label.Alignment = fyne.TextAlignCenter

	change := widget.NewButton("change", func() {
		text = "changed"
		label.SetText(text)
		label.Importance = widget.HighImportance
		label.Refresh()
	})
	box := container.NewVBox(label, change)
	w.SetContent(box)
	w.ShowAndRun()
}
