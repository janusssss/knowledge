package main

import (
	"fmt"
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/data/binding"
	"fyne.io/fyne/v2/widget"
	"time"
)

func main() {
	a := app.New()
	w := a.NewWindow("item")
	w.Resize(fyne.NewSize(300, 300))

	item := binding.NewItem(func(t [2]time.Time, t2 [2]time.Time) bool {
		return t == t2
	})
	err := item.Set([2]time.Time{time.Now(), time.Now()})
	if err != nil {
		panic(err)
	}

	change := widget.NewButton("change", func() {
		period, err := item.Get()
		if err != nil {
			panic(err)
		}
		fmt.Println(period)
	})

	content := container.NewVBox(change)
	w.SetContent(content)
	w.ShowAndRun()
}
