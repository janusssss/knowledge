package main

import (
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/widget"
)

func main() {
	a := app.New()
	w := a.NewWindow("entry")

	entry := widget.NewEntry()
	entry.SetText("Hello World")
	entry.TextStyle.TabWidth = 100
	entry.TextStyle.Underline = true
	entry.TextStyle.Bold = true

	w.SetContent(container.NewVBox(entry))
	w.Resize(fyne.NewSize(300, 300))
	w.ShowAndRun()
}
