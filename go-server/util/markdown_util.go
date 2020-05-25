package util

import (
	"io/ioutil"
	"log"
	"os"
	"os/exec"
)

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
