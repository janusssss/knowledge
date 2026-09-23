package main

import (
	"fmt"
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/canvas"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/widget"
	"image/color"
)

func main() {
	// 创建应用和窗口
	myApp := app.New()
	myWindow := myApp.NewWindow("Fyne List 示例")

	// 列表数据源
	data := []string{"苹果", "香蕉", "橙子", "葡萄", "西瓜", "芒果", "蓝莓"}

	// 创建列表组件
	list := widget.NewList(
		func() int {
			return 1
		},
		func() fyne.CanvasObject {
			return canvas.NewText("hello world", color.White)
		},
		func(i widget.ListItemID, o fyne.CanvasObject) {
			//o.(*widget.Label).SetText(data[i])
		},
	)

	// 点击事件处理
	list.OnSelected = func(id widget.ListItemID) {
		fmt.Printf("你点击了第 %d 项: %s\n", id, data[id])
	}

	// 设置窗口内容并显示
	myWindow.SetContent(container.NewScroll(list))
	myWindow.Resize(fyne.NewSize(300, 200))
	myWindow.ShowAndRun()
}
