package main

import (
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/widget"
)

func main() {
	// 创建应用和窗口
	myApp := app.New()
	myWindow := myApp.NewWindow("Card 示例")

	// 创建一个简单的 Card
	card := widget.NewCard(
		"这是标题",    // 标题（title）
		"这是一个副标题", // 副标题（subtitle）
		widget.NewLabel("这里是卡片的内容区域。你可以放任何 fyne.CanvasObject，比如图片、按钮、列表等。"), // 内容（content）
	)

	// 设置窗口内容
	myWindow.Resize(fyne.Size{Width: 300, Height: 300})
	myWindow.SetContent(container.NewPadded(card))
	myWindow.ShowAndRun()
}
