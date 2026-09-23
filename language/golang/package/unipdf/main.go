package main

import (
	"fmt"
	"github.com/unidoc/unipdf/v3/common/license"
	"os"

	"github.com/unidoc/unipdf/v3/extractor"
	"github.com/unidoc/unipdf/v3/model"
)

func init() {
	// Make sure to load your metered License API key prior to using the library.
	// If you need a key, you can sign up and create a free one at https://cloud.unidoc.io
	err := license.SetMeteredKey(`2b12654a18ae0241742eda3addced0a23f2c1eca0fa19ca565c58a4488ee6e40`)
	if err != nil {
		panic(err)
	}
}

func main() {
	// 打开PDF文件
	inputPath := "compose/unipdf/通用规范汉字表.pdf"
	f, err := os.Open(inputPath)
	if err != nil {
		fmt.Printf("Error: %v\n", err)
		return
	}
	defer f.Close()

	// 创建一个新的PDF阅读器
	pdfReader, err := model.NewPdfReader(f)
	if err != nil {
		fmt.Printf("Error: %v\n", err)
		return
	}

	// 获取PDF文件的页数
	numPages, err := pdfReader.GetNumPages()
	if err != nil {
		fmt.Printf("Error: %v\n", err)
		return
	}

	for i := 0; i < numPages; i++ {
		// 提取每一页的文本
		page, err := pdfReader.GetPage(i + 1)
		if err != nil {
			fmt.Printf("Error: %v\n", err)
			return
		}

		ext, err := extractor.New(page)
		if err != nil {
			fmt.Printf("Error: %v\n", err)
			panic(err)
		}
		extractText, err := ext.ExtractText()
		if err != nil {
			fmt.Printf("Error: %v\n", err)
			panic(err)
		}
		fmt.Printf("Page %d:\n%s\n", i+1, extractText)
	}
}
