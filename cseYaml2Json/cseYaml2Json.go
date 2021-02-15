package cseYaml2Json

import (
	"bytes"
	"encoding/json"
	"fmt"
	"gopkg.in/yaml.v2"
	"log"
	"strconv"
	//"io/ioutil"
)

var info = map[string]string{}
var ikey = []string{}

//func main() {
//	buffer, err := ioutil.ReadFile("./config_sjdf.yaml")
//	if err != nil {
//		log.Fatalf(err.Error())
//	}
//	mString := YAML2JSON(buffer)
//	fmt.Println("==== out ====")
//	fmt.Println(mString)
//}

func YAML2JSON(y []byte) string {
	var yamlObj interface{}
	err := yaml.Unmarshal(y, &yamlObj)
	if err != nil {
		log.Fatalf(err.Error())
	}

	CseJson(yamlObj)


	//mjson,_ :=json.MarshalIndent(info,"\r\n","")
	mjson,_ :=json.MarshalIndent(info,"","\t")
	mString := string(mjson)

	info = map[string]string{}
	ikey = []string{}

	//return strings.Replace(mString,"\\","",-1)
	return mString
}


func CseJson(yamlObj interface{})  {
	switch typedYAMLObj := yamlObj.(type) {
	case map[interface{}]interface{}:
		for k, v := range typedYAMLObj {
			switch typedValue := v.(type) {
			case map[interface{}]interface{}:
				ikey = append(ikey, k.(string))
				//fmt.Println(typedValue)
				CseJson(typedValue)
			case string:
				if len(ikey) >= 1 {
					ikey = append(ikey, k.(string))
					sk := sli2str(ikey)
					if typedValue == "" {
						typedValue = `""`
					}
					info[sk] = typedValue
					ikey = ikey[:len(ikey)-1]
				}else {
					info[k.(string)] = typedValue
				}
			case []interface {}:
				outstr := sli2slimap(typedValue)
				if len(ikey) >= 1 {
					ikey = append(ikey, k.(string))
					sk := sli2str(ikey)
					json_str,_ := json.Marshal(outstr)
					info[sk] = string(json_str)
					ikey = ikey[:len(ikey)-1]
				}else {
					json_str,_ := json.Marshal(outstr)
					info[k.(string)] = string(json_str)
				}
			case int:
				if len(ikey) >= 1 {
					ikey = append(ikey, k.(string))
					sk := sli2str(ikey)
					info[sk] = strconv.Itoa(typedValue)
					ikey = ikey[:len(ikey)-1]
				}else {
					info[k.(string)] = strconv.Itoa(typedValue)
				}
			case int64:
				if len(ikey) >= 1 {
					ikey = append(ikey, k.(string))
					sk := sli2str(ikey)
					info[sk] = strconv.FormatInt(typedValue, 10)
					ikey = ikey[:len(ikey)-1]
				}else {
					info[k.(string)] = strconv.FormatInt(typedValue, 10)
				}
			case float64:
				if len(ikey) >= 1 {
					ikey = append(ikey, k.(string))
					sk := sli2str(ikey)
					info[sk] = strconv.FormatFloat(typedValue,'g', -1, 32)
					ikey = ikey[:len(ikey)-1]
				}else {
					info[k.(string)] = strconv.FormatFloat(typedValue,'g', -1, 32)
				}
			case uint64:
				if len(ikey) >= 1 {
					ikey = append(ikey, k.(string))
					sk := sli2str(ikey)
					info[sk] = strconv.FormatUint(typedValue, 10)
					ikey = ikey[:len(ikey)-1]
				}else {
					info[k.(string)] = strconv.FormatUint(typedValue, 10)
				}
			case bool:
				var flages string
				if typedValue {
					flages = "true"
				}else {
					flages = "false"
				}
				if len(ikey) >= 1 {
					ikey = append(ikey, k.(string))
					sk := sli2str(ikey)
					info[sk] = flages
					ikey = ikey[:len(ikey)-1]
				}else {
					info[k.(string)] = flages
				}
			case nil:
				if len(ikey) >= 1 {
					ikey = append(ikey, k.(string))
					sk := sli2str(ikey)
					info[sk] = `""`
					ikey = ikey[:len(ikey)-1]
				}else {
					info[k.(string)] = `""`
				}

			default:
				fmt.Printf("unkonw type >>. %T \n", v)
			}
		}
		if len(ikey) <= 1 {
			ikey = []string{}
		}else {
			ikey = ikey[:len(ikey)-1]
		}

	}

}

func sli2slimap(sli []interface{}) []interface{} {
	var slibuf []interface{}
	for _, v := range sli {
		switch tv := v.(type) {
		case map[interface{}]interface{}:
			var sm map[string]interface{}
			sm = map2str(tv)
			slibuf = append(slibuf, sm)
		case string:
			slibuf = append(slibuf, tv)
		case int:
			slibuf = append(slibuf, strconv.FormatInt(int64(tv), 10))
		case int64:
			slibuf = append(slibuf, strconv.FormatInt(tv, 10))
		case float64:
			slibuf = append(slibuf, strconv.FormatFloat(tv,'g', -1, 32))
		case uint64:
			slibuf = append(slibuf, strconv.FormatUint(tv, 10))
		case bool:
			if tv {
				slibuf = append(slibuf, "true")
			}else {
				slibuf = append(slibuf, "false")
			}
		}
	}
	return slibuf
}

func map2str(smap map[interface{}]interface{}) map[string]interface{} {
	var mapstr = map[string]interface{}{}
	for k,v := range smap {
		switch mapv := v.(type) {
		case map[interface{}]interface{}:
			mm := map2str(v.(map[interface{}]interface{}))
			mapstr[k.(string)] = mm
		case []interface{}:
			sl := sli2slimap(v.([]interface{}))
			mapstr[k.(string)] = sl
		case string:
			mapstr[k.(string)] = mapv
		case int:
			mapstr[k.(string)] = strconv.FormatInt(int64(mapv), 10)
		case int64:
			mapstr[k.(string)] = strconv.FormatInt(mapv, 10)
		case float64:
			mapstr[k.(string)] = strconv.FormatFloat(mapv,'g', -1, 32)
		case uint64:
			mapstr[k.(string)] = strconv.FormatUint(mapv, 10)
		case bool:
			if mapv {
				mapstr[k.(string)] = "true"
			}else {
				mapstr[k.(string)] = "false"
			}
		default:
			mapstr[k.(string)] = `""`
		}
	}
	return mapstr
}

func sli2str(sli []string) string {
	var sk string

	//fmt.Println("------------------------> ",sli)
	for _, v := range sli {
		if sk != ""{
			sk = sk + "." + v
		}else {
			sk = v
		}

	}
	return sk
}

// 去转义
func drop_tm(ss string) string {
	var buf bytes.Buffer
	for _, c := range ss {
		if c == '\\' {
			continue
		}
		buf.WriteRune(c)
	}
	//fmt.Println("           tm      :", ss)
	return ss
}