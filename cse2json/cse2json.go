package main

import (
	"cseYaml2Json"
	"fmt"
	"io/ioutil"
	"log"
	"path"
	"path/filepath"
	"regexp"
	"strings"

	"os"
)

func main() {

	for _,v := range os.Args[1:] {
		filepath.Walk(v, func(p string, info os.FileInfo, err error) error {
			if info == nil {
				return err
			}
			if info.IsDir() {
				return nil
			}else {
				path_ex, files := filepath.Split(p)
				file_ext := path.Ext(files)
				file_name := strings.Split(files,file_ext)[0]


				if judge,_ := regexp.MatchString(".yaml$|.yml$",files); judge {
					//fmt.Printf("=======> : %s\n",files,p)
					if checkFileIsExist(p) {
						buffer, err := ioutil.ReadFile(p)
						if err != nil {
							log.Fatalf(err.Error())
						}
						mString := cseYaml2Json.YAML2JSON(buffer)
						s := []byte(mString)
						fmt.Println(path_ex+file_name+".json")
						ioutil.WriteFile(path_ex+file_name+".json",s,0666)
					}else {
						errorInfo()
					}
				}
			}
			return nil
		})
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