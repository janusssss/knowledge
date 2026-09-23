package main

import (
	"fmt"
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/layout"
	"fyne.io/fyne/v2/widget"
	"log"
	"math"
	"time"
)

func main() {
	myApp := app.New()
	window := myApp.NewWindow("Popup")
	var scroll *container.Scroll
	button := widget.NewButton("显示 Popup", func() {
		weekdays := []string{"星期一", "星期二", "星期三", "星期四", "星期五", "星期六", "星期日"}
		boxLayout := layout.NewCustomPaddedVBoxLayout(5)
		box := container.New(boxLayout)
		box.Add(widget.NewButton("空白", nil))
		for _, weekday := range weekdays {
			box.Add(widget.NewButton(weekday, nil))
		}
		box.Add(widget.NewButton("空白", nil))

		scroll = container.NewVScroll(box)

		size := box.Objects[1].MinSize()
		fmt.Println("size:", size)
		c := container.NewGridWrap(fyne.NewSize(size.Width*2, size.Height*3+10), scroll)

		values := []float32{0, 41, 82, 123, 164, 205, 246, math.MaxInt64}
		move := make(chan float32, 1)
		scroll.OnScrolled = func(position fyne.Position) {
			move <- position.Y
			fmt.Println(position)

			//for i, v := range values {
			//	if position.Y > v {
			//		scroll.ScrollToOffset(fyne.NewPos(0, values[i+1]))
			//		return
			//	}
			//}
		}
		change := widget.NewButton("change", func() {
			scroll.ScrollToOffset(fyne.NewPos(0, 123))
		})

		go func() {
			for {
				select {
				case h := <-move:
					select {
					case temp := <-move:
						move <- temp
					case <-time.After(time.Millisecond * 300):
						fmt.Println("h:", h)
						for i, v := range values {
							if v <= h && h <= values[i+1] {
								if h-v < values[i+1]-h {
									h = v
								} else {
									h = values[i+1]
								}
								fyne.Do(func() {
									scroll.ScrollToOffset(fyne.NewPos(0, h))
									log.Println("move,h: ", h)
								})
								break
							}
						}
					}
				}
			}
		}()

		cc := container.NewBorder(change, nil, nil, nil, c)
		content := container.NewCenter(cc)
		popup := widget.NewPopUp(content, window.Canvas())
		popup.Resize(fyne.NewSize(400, 300))
		popup.Move(fyne.NewPos(0, 100))
		popup.Show()
		scroll.ScrollToOffset(fyne.NewPos(0, 82))
	})

	window.SetContent(container.NewVBox(button))
	window.Resize(fyne.NewSize(400, 400))
	window.ShowAndRun()
}
