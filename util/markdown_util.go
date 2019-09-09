package util

import (
	"github.com/b3log/lute"
	"io/ioutil"
	"log"
	"os"
	"os/exec"
)

func ConvertLuteMarkdown(markdown string) string {
	luteEngine := lute.New() // 默认已经启用 GFM 支持以及中文语境优化
	luteEngine.SetCodeSyntaxHighlightInlineStyle(true)
	html, err := luteEngine.MarkdownStr("", markdown)
	if nil != err {
		panic(err)
	}
	return html
}

// ConvertMarkdown 通过pandoc 转换markdown转换成html
func ConvertMarkdown(c string) (markdownHTML string) {
	inputFile := createTempFile("source_")
	defer os.Remove(inputFile.Name())
	inputFile.Write([]byte(c))

	outputFile := createTempFile("converted_")
	defer os.Remove(outputFile.Name())

	args := []string{
		"--standalone",
		"--from", "markdown_github",
		"--to", "html",
		"--highlight-style=kate",
		"--output", outputFile.Name(),
		inputFile.Name(),
	}

	err := exec.Command("pandoc", args...).Run()
	if err != nil {
		log.Panic(err)
	}

	data, err := ioutil.ReadAll(outputFile)
	if err != nil {
		log.Panic(err)
	}

	return string(data)
}

func createTempFile(prefix string) *os.File {
	tempFile, err := ioutil.TempFile(os.TempDir(), prefix)
	if err != nil {
		log.Fatal(err)
	}
	return tempFile
}
