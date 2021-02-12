package main

import (
	"fmt"
	"os"
	"regexp"
)

func main() {
	//if judge,_ := regexp.MatchString(".[yaml|yml]$","go.mod"); judge {
	//	fmt.Println("jjj")
	//}

	//if ioutil.ReadDir()
	//fileInfoList,err := ioutil.ReadDir(pwd)

	for _, v := range os.Args[1:] {
		//fmt.Println(v)

		if judge,_ := regexp.MatchString("^-",v); judge {
			switch v {
			case "-a":
				fmt.Println(v)
			default:
				fmt.Println("default")
			}
		}

		//list, err := ioutil.ReadDir(v)
		//if err != nil { log.Fatal(err.Error())}
		//for i := range list {
		//	fmt.Println(list[i].Name())  //打印当前文件或目录下的文件或目录名
		//}
	}

}
