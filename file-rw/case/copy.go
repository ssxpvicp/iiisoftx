package _case

import (
	"fmt"
	"io"
	"log"
	"os"
	"path"
)

func CopyDirToDir() {
	list := getFiles(sourceDir)

	for _, f := range list {
		if f == "" {
			continue
		}
		_, name := path.Split(f)
		destFileName := destDir + "copy/" + name
		//复制文件
		//fmt.Println(destFileName)
		fmt.Println(f)
		CopyFile(f, destFileName)
	}
}

// 复制文件方法
func CopyFile(srcName, destName string) (int64, error) {
	src, err := os.Open(srcName)
	if err != nil {
		log.Fatalln(srcName, err)
	}
	defer src.Close()

	file, err := os.OpenFile(destName, os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		log.Fatalln(err)
	}
	defer file.Close()
	return io.Copy(file, src)
}
