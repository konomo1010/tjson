package main

import (
	"cseYaml2Json"
	"fmt"
	"github.com/akamensky/argparse"
	"io/ioutil"
	"log"
	"os"
	"path"
	"path/filepath"
	"regexp"
	"strings"
)

func main() {
	//var args []string
	// Create new parser object
	parser := argparse.NewParser("print", "Prints provided string to stdout")
	var myFlag *bool = parser.Flag("r", "recursion", &argparse.Options{ Help: "recursion ... ... ..."})
	var files *[]string = parser.StringList("f", "files", &argparse.Options{ Help: "input dirname"})
	var dir *[]string = parser.StringList("d","dirs", &argparse.Options{Help: "input dirname"})
	var out *string = parser.String("o","outpath", &argparse.Options{Help: "output path"})
	err := parser.Parse(os.Args)
	if err != nil {
		fmt.Print(parser.Usage(err))
	}

	fmt.Println(*myFlag, *files, *dir, *out)
	if len(*files) != 0 {
		for _,v := range *files {
			if PathExists(v) && IsFile(v) {
				var outflag bool
				if len(*out) != 0 {
					outflag = true
				}else {
					outflag = false
				}
				fmt.Println("file : " + v)
				path_ex, files := filepath.Split(v)
				file_ext := path.Ext(files)
				file_name := strings.Split(files,file_ext)[0]
				if judge,_ := regexp.MatchString(".yaml$|.yml$",v); judge {
					fmt.Println("here ")
					buffer, err := ioutil.ReadFile(v)
					if err != nil {
						log.Fatalf(err.Error())
					}
					mString := cseYaml2Json.YAML2JSON(buffer)
					s := []byte(mString)
					if outflag {
						fmt.Println(*out+file_name+".json")
						ioutil.WriteFile(*out+file_name+".json",s,0666)
					}else {
						fmt.Println(path_ex+file_name+".json")
						ioutil.WriteFile(path_ex+file_name+".json",s,0666)
					}
				}else {
					errorInfo()
				}
			}
		}
	}

	if len(*dir) != 0 {
		var outflag bool = false
		if len(*out) != 0 {
			outflag = true
		}
		if !*myFlag {
			for _, dv := range *dir {
				dir_filelist, direrr := ioutil.ReadDir(dv)
				if direrr != nil {
					log.Fatalf(direrr.Error())
				}
				for _, dirfile := range dir_filelist {
					if !dirfile.IsDir() {
						if judge,_ := regexp.MatchString(".yaml$|.yml$",dirfile.Name()); judge {
							file_ext := path.Ext(dirfile.Name())
							file_name := strings.Split(dirfile.Name(),file_ext)[0]
							var pathfile string
							if outflag {
								pathfile = *out + "/" + file_name + ".json"
							}else {
								pathfile = dv + "/" + file_name+".json"
							}
							buffer, err := ioutil.ReadFile(dv+"/"+dirfile.Name())
							if err != nil {
								log.Fatalf(err.Error())
							}
							mString := cseYaml2Json.YAML2JSON(buffer)
							s := []byte(mString)
							fmt.Println("-->"+pathfile)
							ioutil.WriteFile(pathfile,s,0666)
						}
					}
				}
			}
		}else {
			for _,dirstr := range *dir {
				filepath.Walk(dirstr, func(p string, info os.FileInfo, err error) error {
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
								if outflag {
									fmt.Println(*out+"/"+file_name+".json")
									ioutil.WriteFile(*out+"/"+file_name+".json",s,0666)
								}else {
									fmt.Println(path_ex+"/"+file_name+".json")
									ioutil.WriteFile(path_ex+"/"+file_name+".json",s,0666)
								}

							}
						}
					}
					return nil
				})
			}
		}

	}

}

	//if *myFlag {
	//	args = delslielm(os.Args, "-r")
	//	for _,v := range args[1:] {
	//		filepath.Walk(v, func(p string, info os.FileInfo, err error) error {
	//			if info == nil {
	//				return err
	//			}
	//			if info.IsDir() {
	//				return nil
	//			}else {
	//				path_ex, files := filepath.Split(p)
	//				file_ext := path.Ext(files)
	//				file_name := strings.Split(files,file_ext)[0]
	//
	//
	//				if judge,_ := regexp.MatchString(".yaml$|.yml$",files); judge {
	//					//fmt.Printf("=======> : %s\n",files,p)
	//					if checkFileIsExist(p) {
	//						buffer, err := ioutil.ReadFile(p)
	//						if err != nil {
	//							log.Fatalf(err.Error())
	//						}
	//						mString := cseYaml2Json.YAML2JSON(buffer)
	//						s := []byte(mString)
	//						fmt.Println(path_ex+file_name+".json")
	//						ioutil.WriteFile(path_ex+file_name+".json",s,0666)
	//					}else {
	//						errorInfo()
	//					}
	//				}
	//			}
	//			return nil
	//		})
	//	}
	//} else {
	//	fmt.Println("no -r ")
	//}



func PathExists(path string) (bool) {
	_, err := os.Stat(path)
	if err == nil {
		return true
	}
	if os.IsNotExist(err) {
		return false
	}
	fmt.Println(path + "is not exist")
	return false
}

func IsDir(path string) bool {
	s, err := os.Stat(path)
	if err != nil {
		return false
	}
	return s.IsDir()
}


func IsFile(path string) bool {
	return !IsDir(path)
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

// 删除 切片 指定元素
func delslielm(sli []string, str string) []string {
	for i, v := range sli {
		if v == str {
			sli = append(sli[:i], sli[i+1:]...)
			break
		}
	}
	return sli
}