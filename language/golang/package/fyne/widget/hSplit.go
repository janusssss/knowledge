package main

import (
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/widget"
)

func main() {
	a := app.New()
	w := a.NewWindow("hSplit")
	w.Resize(fyne.NewSize(400, 400))

	leading := widget.NewLabel("Leading")
	trailing := widget.NewLabel("Trailing")
	split := container.NewHSplit(leading, trailing)

	w.SetContent(split)
	w.ShowAndRun()
}
