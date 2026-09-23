package main

import (
	"fmt"
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/data/binding"
	"fyne.io/fyne/v2/widget"
)

func main() {
	a := app.New()
	w := a.NewWindow("")
	data := make([]string, 0, 10)
	for _, v := range []string{"a", "b", "c"} {
		data = append(data, v)
	}
	fmt.Printf("data: %p\n", &data)
	//selectWidget := widget.NewSelect(data, nil)
	b := binding.NewString()
	selectWidget := widget.NewSelectWithData(data, b)
	selectWidget.SetSelectedIndex(0)

	button := widget.NewButton("change data", nil)

	button.OnTapped = func() {
		//for i, v := range data {
		//	data[i] = strings.ToUpper(v)
		//}
		data = append(data, "D")
		fmt.Printf("data: %p\n", &data)
		button.Refresh()

		fmt.Println(selectWidget.Options)
		fmt.Println(data)
		//selectWidget.SetOptions(data)
		//selectWidget.SetSelectedIndex(len(data) - 1)
	}

	grib := container.NewGridWithColumns(1)
	grib.Add(selectWidget)
	grib.Add(button)

	w.SetContent(grib)
	w.Resize(fyne.NewSize(300, 300))
	w.ShowAndRun()
}
