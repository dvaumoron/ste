package main

import (
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
		panic(err)
	}

	tmpl, err := template.New("tmpl").Parse(string(tmplBody))
	if err != nil {
		panic(err)
	}

	file, err := os.Create(outPath)
	if err != nil {
		panic(err)
	}
	defer file.Close()

	err = tmpl.Execute(file, args[3:])
	if err != nil {
		panic(err)
	}
}
