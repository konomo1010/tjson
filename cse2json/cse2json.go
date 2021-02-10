package main

import (
	"cseYaml2Json"
	"fmt"
	"io/ioutil"
	"log"
	"path"
	"path/filepath"
	"strings"

	//"io/ioutil"
	//"log"
	"os"
	"regexp"
)

func main() {

	if len(os.Args) == 1 {
		errorInfo()
		os.Exit(100)
	}else {
		//fmt.Println(os.Args[0])// args 第一个片 是文件路径
		for _,v := range os.Args[1:] {
			paths, files := filepath.Split(v)
			file_ext := path.Ext(files)
			file_name := strings.Split(files,file_ext)[0]
			fmt.Println(paths, files, file_name, file_ext)
			if judge,_ := regexp.MatchString(".[yaml|yml]$",files); judge {
				//fmt.Println(v)
				//fmt.Println(strings.Split(filepath.Base(v),".")[0])




				// 判断文件是否存在
				if checkFileIsExist(v) {
					buffer, err := ioutil.ReadFile(v)
					if err != nil {
						log.Fatalf(err.Error())
					}
					mString := cseYaml2Json.YAML2JSON(buffer)
					s := []byte(mString)
					ioutil.WriteFile(paths+file_name+".json",s,0666)
				}else {
					errorInfo()
				}


			} else {
				errorInfo()
			}
		}
	}

}

func errorInfo() {
	fmt.Println("Error : sorry, You need to enter [yaml|yml] file!~, please check!~")
}

func checkFileIsExist(filename string) bool {
	var exist = true
	if _, err := os.Stat(filename); os.IsNotExist(err) {
		exist = false
	}
	return exist
}