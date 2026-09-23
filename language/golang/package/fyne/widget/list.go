package main

import (
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/widget"
)

func main() {
	myApp := app.New()
	myWindow := myApp.NewWindow("List Widget")
	myWindow.Resize(fyne.NewSize(400, 400))
	data := []string{"1", "2", "3", "4", "5"}
	var list *widget.List
	list = widget.NewList(
		func() int {
			return len(data)
		},
		func() fyne.CanvasObject {
			return widget.NewLabel("1")
		},
		func(id widget.ListItemID, o fyne.CanvasObject) {
			o.(*widget.Label).SetText(data[id])
		})

	myWindow.SetContent(container.NewBorder(nil, widget.NewButton("change", func() {
		data = append(data, "6")
		list.Refresh()
	}), nil, nil, list))




	myWindow.ShowAndRun()
}
