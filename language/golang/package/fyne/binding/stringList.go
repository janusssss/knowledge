package main

import (
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/data/binding"
	"fyne.io/fyne/v2/widget"
	"log"
)

func main() {
	a := app.New()
	w := a.NewWindow("")
	w.Resize(fyne.NewSize(300, 300))

	list := binding.NewStringList()
	srt := []string{"hello", "world"}
	err := list.Set(srt)
	if err != nil {
		log.Fatal(err)
	}
	change := widget.NewButton("change", func() {
		err = list.SetValue(1, "change world")
		if err != nil {
			panic(err)
		}
		list.Remove("hello")
		list.Prepend("hello")
		log.Println("change success")
	})

	set := widget.NewButton("set", func() {
		vlaues, err := list.Get()
		if err != nil {
			panic(err)
		}
		vlaues[0] = "set"
		list.Remove(vlaues[1])
		list.Set(vlaues)
		log.Println("change success")
	})
	add := widget.NewButton("add", func() {
		err = list.Append("change world")
		if err != nil {
			panic(err)
		}
		log.Println("change success")
	})

	values, err := list.Get()
	if err != nil {
		panic(err)
	}

	subContent := objects(values)
	content := container.NewVBox(subContent, change, add, set)
	w.SetContent(content)

	list.AddListener(binding.NewDataListener(func() {
		values, err = list.Get()
		if err != nil {
			panic(err)
		}
		content.Remove(subContent)
		subContent = objects(values)
		content.Add(subContent)
		log.Println("listener success")
	}))

	w.ShowAndRun()

}

func objects(items []string) fyne.CanvasObject {
	box := container.NewVBox()
	for _, item := range items {
		box.Add(widget.NewLabel(item))
	}
	return box
}
