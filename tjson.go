package main

import (
	"cseYaml2Json"
	"github.com/lxn/walk"
	. "github.com/lxn/walk/declarative"
	"log"
)

func main() {
	var inTE, outTE *walk.TextEdit
	var mw *walk.MainWindow

	if _, err := (MainWindow{
		AssignTo: &mw,
		Title:   "cse json transform",
		MinSize: Size{600, 400},
		Layout:  VBox{},
		Children: []Widget{
			HSplitter{
				Children: []Widget{
					TextEdit{AssignTo: &inTE},
					TextEdit{AssignTo: &outTE, ReadOnly: true},
				},
				OnBoundsChanged: func() {
					//var windowIcon *walk.Icon
					//icon, err := walk.Resources.Icon("./Batman-emoticon.ico")
					//if err != nil {
					//	log.Fatal(err)
					//}
					//mw.SetIcon(icon)
					//if windowIcon != nil {
					//	windowIcon.Dispose()
					//}
					//windowIcon = icon
				},
			},
			PushButton{
				Text: "cse json transform",
				OnClicked: func() {
					var outstr string
					if inTE.Text() == "" {
						outstr = ""
					}else {
						outstr = tjson(inTE.Text())
					}
					outTE.SetText(outstr)
				},
			},
		},
	}.Run()); err != nil {
		log.Fatal(err)
	}
}

// 转 cse json
func tjson(s string) string {
	tm :=  []byte(s)
	outstr := cseYaml2Json.YAML2JSON(tm)
	return outstr
}

