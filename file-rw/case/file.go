package _case

import (
	"fmt"
	"log"
	"os"
	"strings"
)

// 源文件目录
const sourceDir = "file-rw/source-file/"

// 目录文件目录
const destDir = "file-rw/dest-file/"

func getFiles(dir string) []string {
	fs, err := os.ReadDir(dir)
	if err != nil {
		log.Fatalln(err)

	}
	list := make([]string, 0)

	for _, f := range fs {
		if f.IsDir() {
			continue
		}
		fmt.Println(f.Name())
		fullName := strings.Trim(dir, "/") + "/" + f.Name()
		fmt.Println(fullName)
		fmt.Println("")
		list = append(list, fullName)
	}
	return list
}
