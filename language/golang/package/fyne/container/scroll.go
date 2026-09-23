package main

import (
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/widget"
)

func main() {
	a := app.New()
	w := a.NewWindow("自定义滚轮选择星期")

	weekdays := []string{"星期一", "星期二", "星期三", "星期四", "星期五", "星期六", "星期日"}

	list := widget.NewList(
		func() int { return len(weekdays) },
		func() fyne.CanvasObject { return widget.NewLabel("") },
		func(id widget.ListItemID, obj fyne.CanvasObject) {
			obj.(*widget.Label).SetText(weekdays[id])
		},
	)

	// 限制列表高度，模拟滚轮效果
	scroll := container.NewVScroll(list)
	scroll.Resize(fyne.NewSize(200, 100)) // 只显示一个选项的高度

	// 监听选择（需结合滚动位置计算，此处简化）
	list.OnSelected = func(id widget.ListItemID) {
		println("选择了:", weekdays[id])
	}

	w.Resize(fyne.NewSize(300, 300))
	w.SetContent(container.NewCenter(scroll))
	w.ShowAndRun()
}
