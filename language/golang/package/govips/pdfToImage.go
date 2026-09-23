package main

import (
	"fmt"
	"github.com/davidbyttow/govips/v2/vips"
	"os"
	"strconv"
)

func checkError(err error) {
	if err != nil {
		fmt.Println("error:", err)
		os.Exit(1)
	}
}

func main() {
	vips.Startup(nil)
	defer vips.Shutdown()

	filePath := "govips/example.pdf"
	file, err := vips.NewImageFromFile(filePath)
	checkError(err)
	pages := file.Pages()
	file.Close()

	for page := range pages {
		params := vips.NewImportParams()
		params.Page.Set(page)
		image, err := vips.LoadImageFromFile(filePath, params)
		checkError(err)
		native, _, err := image.ExportNative()
		checkError(err)
		s := "govips/" + strconv.Itoa(page) + ".jpeg"
		err = os.WriteFile(s, native, 0644)
		checkError(err)
		image.Close()
	}
}
