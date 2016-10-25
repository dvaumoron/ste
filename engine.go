package main

import (
    "fmt"
	"io/ioutil"
	"os"
	"text/template"
)

func main() {
	args := os.Args

	tmplPath := args[1]
	outPath := args[2]

	tmplBody, err := ioutil.ReadFile(tmplPath)
	if err != nil {
        fmt.Println(err)
		return
	}

	tmpl, err := template.New("tmpl").Parse(string(tmplBody))
	if err != nil {
        fmt.Println(err)
		return
	}

	file, err := os.Create(outPath)
	if err != nil {
        fmt.Println(err)
		return
	}
	defer file.Close()

	err = tmpl.Execute(file, args[3:])
	if err != nil {
        fmt.Println(err)
		return
	}

    fmt.Println(outPath, "generated")
}
