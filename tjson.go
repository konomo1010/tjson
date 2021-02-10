package main

import (
	"cseYaml2Json"
	"fmt"
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
					TextEdit{
						AssignTo: &inTE,
						VScroll: true,
						TextColor:walk.RGB(73,156,84),
					},
					TextEdit{
						AssignTo: &outTE,
						ReadOnly: true,
						VScroll: true,
						CompactHeight: false,
						ColumnSpan: 1,
					},
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
				//MinSize: Size{10, 20},
				OnClicked: func() {
					//var outstr string
					//if inTE.Text() == "" {
					//	outstr = ""
					//}else {
					//	outstr = tjson(inTE.Text())
					//}

					outTE.SetText(tjson(inTE.Text()))

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
	fmt.Println(outstr)
	return outstr
}

