package main

import (
	"bytes"
	"html/template"
	"io/ioutil"
	"log"
	"os"
)

func main() {
	//	const tpl = `
	//<!DOCTYPE html>
	//<html>
	//	<head>
	//		<meta charset="UTF-8">
	//		<title>{{.Title}}</title>
	//	</head>
	//	<body>
	//		{{range .Items}}<div>{{ . }}</div>{{else}}<div><strong>no rows</strong></div>{{end}}
	//	</body>
	//</html>`

	check := func(err error) {
		if err != nil {
			log.Fatal(err)
		}
	}
	// 读取
	originPath := "resource/demo.html"
	open, _ := os.Open(originPath)
	htmlData, err := ioutil.ReadAll(open)
	path := "Data/report/files/1.html"
	temp := template.New(path)
	t, err := temp.Parse(string(htmlData))
	check(err)

	data := struct {
		Title string
		Items []string
	}{
		Title: "111",
		Items: []string{
			"222",
			"333",
		},
	}
	outputBuf := new(bytes.Buffer)
	err = t.Execute(outputBuf, data)
	//err = t.Execute(os.Stdout, data)
	check(err)

	//noItems := struct {
	//	Title string
	//	Items []string
	//}{
	//	Title: "444",
	//	Items: []string{},
	//}

	//err = t.Execute(outputBuf, noItems)
	//err = t.Execute(os.Stdout, noItems)
	//check(err)
	f, err := os.OpenFile("test.html", os.O_WRONLY|os.O_CREATE|os.O_APPEND, 0600)
	err = t.Execute(f, data)
	//err = ioutil.WriteFile(path, outputBuf.Bytes(), 0666)
}
