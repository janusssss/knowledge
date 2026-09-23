package main

import (
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/theme"
	"fyne.io/fyne/v2/widget"
)

func main() {
	a := app.New()
	w := a.NewWindow("RichTextWithText")
	w.Resize(fyne.NewSize(400, 400))

	richText := widget.NewRichText(
		&widget.TextSegment{
			Text: "红色粗体",
			Style: widget.RichTextStyle{
				ColorName: theme.ColorNameError,
			},
		},
		&widget.TextSegment{
			Text: " 斜体",
			Style: widget.RichTextStyle{
				ColorName: theme.ColorNameSuccess,
				TextStyle: fyne.TextStyle{
					Italic:    true,
					Underline: true,
					Bold:      true,
				},
			},
		},
		&widget.TextSegment{
			Text: " 斜体",
			Style: widget.RichTextStyle{
				ColorName: theme.ColorNameSuccess,
			},
		},
	)

	w.SetContent(richText)
	w.ShowAndRun()
}
