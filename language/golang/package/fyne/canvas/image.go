package main

import (
	"bytes"
	"fmt"
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/canvas"
	"github.com/wcharczuk/go-chart"
	"image"
)

func main() {
	a := app.New()
	w := a.NewWindow("image_pie")
	w.Resize(fyne.NewSize(400, 400))
	
	pie := chart.PieChart{
		Width:  512,
		Height: 512,
		Values: []chart.Value{
			{Value: 5, Label: "蓝色"},
			{Value: 2, Label: "Two"},
			{Value: 1, Label: "One"},
		},
	}
	buffer := bytes.NewBuffer([]byte{})
	err := pie.Render(chart.PNG, buffer)
	if err != nil {
		fmt.Printf("Error rendering pie go-chart: %v\n", err)
	}

	decode, s, err := image.Decode(buffer)
	if err != nil {
		fmt.Printf("Error decoding pie go-chart: %v\n", err)
	}
	fmt.Println(s)

	pieChart := canvas.NewImageFromImage(decode)

	w.SetContent(pieChart)

	w.ShowAndRun()
}
