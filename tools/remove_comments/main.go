package main

import (
	"log"
	"os"
	"path/filepath"
	"regexp"
	"strings"
)

func removeComments(input string) string {
	// Remove block comments
	re1 := regexp.MustCompile(`(\n|^)/\*(\n|.)*?\*/`)
	// Remove line comments
	re2 := regexp.MustCompile(`(\n|^)[[:blank:]]*//.*`)

	result := re1.ReplaceAllString(input, "\n")
	result = re2.ReplaceAllString(result, "\n")
	return result
}

func main() {
	contents, err := os.ReadDir(".")
	if err != nil {
		log.Fatal(err)
	}
	for _, file := range contents {
		if file.IsDir() {
			continue
		} else if strings.TrimRight(file.Name(), ".go") == file.Name() {
			continue
		}
		fpath := filepath.Join(".", file.Name())
		content, err := os.ReadFile(fpath)
		if err != nil {
			log.Fatal(err)
		}
		cleaned := removeComments(string(content))
		f, err := os.CreateTemp(".", "file-XXXXXXXX.go")
		if err != nil {
			log.Fatal(err)
		}
		os.WriteFile(f.Name(), []byte(cleaned), 0644)
		os.Rename(f.Name(), fpath)
	}
}
