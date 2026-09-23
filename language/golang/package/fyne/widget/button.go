package main

import (
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/widget"
)

func main() {
	a := app.New()
	w := a.NewWindow("button")
	w.Resize(fyne.NewSize(300, 300))

	label := widget.NewLabel("label")
	button := widget.NewButton("", func() {
		println("touch button")
	})

	button1 := widget.NewButton("", func() {
		println("touch button1")
	})
	stack := container.NewStack(button, button1, label)
	w.SetContent(stack)
	w.ShowAndRun()
}
