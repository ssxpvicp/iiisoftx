package _case

import (
	"io"
	"log"
	"os"
	"path"
)

// 边读边写
func OneSideReadWriteToDest() {
	list := getFiles(sourceDir)
	for _, l := range list {
		_, name := path.Split(l)
		destFileName := destDir + "one-side/" + name
		//文件写入
		OneSideReadWrite(l, destFileName)
	}
}

func OneSideReadWrite(srcName, destName string) {
	src, err := os.Open(srcName)
	if err != nil {
		log.Fatalln(err)
	}
	defer src.Close()
	dst, err := os.OpenFile(destName, os.O_CREATE|os.O_APPEND|os.O_WRONLY, 0644)
	if err != nil {
		log.Fatalln(err)
	}
	defer dst.Close()

	buf := make([]byte, 1024)
	for {
		n, err := src.Read(buf)
		if err != nil && err != io.EOF {
			log.Fatalln(err)
		}
		if n == 0 {
			//没有读到任何字节说明 读完了
			break
		}
		_, err = dst.Write(buf[:n]) //读到多少写多少.不写:n会写默认的1024
		if err != nil {
			log.Fatalln(err)
		}
	}
}
