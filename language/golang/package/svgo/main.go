package main

import (
	svg "github.com/ajstarks/svgo"
	"os"
)

func main() {
	width, height := 500, 400
	canvas := svg.New(os.Stdout)
	canvas.Start(width, height)
	canvas.Rect(100, 100, 300, 200, "fill:one;stroke:black;stroke-width:3")
	canvas.End()
}
