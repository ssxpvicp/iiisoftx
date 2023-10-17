package _case

import (
	"fmt"
	"log"
	"os"
	"path"
)

func ReadWriteFiles() {
	list := getFiles(sourceDir)
	for _, s := range list {
		fmt.Println(s)
		bytes, err := os.ReadFile(s)
		if err != nil {
			log.Fatalln(err)
		}
		_, name := path.Split(s)
		dstName := destDir + "normal/" + name
		err = os.WriteFile(dstName, bytes, 0644)
		if err != nil {
			log.Fatalln(err)
		}

	}
}
