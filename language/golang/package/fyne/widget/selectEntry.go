package main

import (
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/data/binding"
	"fyne.io/fyne/v2/widget"
)

func main() {
	a := app.New()
	w := a.NewWindow("Entry")
	w.Resize(fyne.NewSize(600, 400))
	bindStr := binding.NewString()
	texts := []string{"aa", "bb", "cc", "dd", "ee"}
	entry := widget.NewSelectEntry(texts)
	entry.Bind(bindStr)

	bindStr.AddListener(binding.NewDataListener(func() {
		text := entry.Text
		println(text)
	}))

	w.SetContent(container.NewVBox(entry))
	w.ShowAndRun()
}
