package main

import (
	"fmt"
	"os"
	"text/template"

	"gopkg.in/yaml.v3"
)

func main() {
	args := os.Args
	if len(args) < 4 {
		fmt.Println("Usage : ste file.tmpl data.yaml outputFile")
		return
	}

	tmplPath := args[1]
	dataPath := args[2]
	outPath := args[3]

	tmplBody, err := os.ReadFile(tmplPath)
	if err != nil {
		fmt.Println(err)
		return
	}

	tmpl, err := template.New("tmpl").Parse(string(tmplBody))
	if err != nil {
		fmt.Println(err)
		return
	}

	dataBody, err := os.ReadFile(dataPath)
	if err != nil {
		fmt.Println(err)
		return
	}

	tmplArgs := map[string]any{}
	err = yaml.Unmarshal(dataBody, tmplArgs)
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

	err = tmpl.Execute(file, tmplArgs)
	if err != nil {
		fmt.Println(err)
		return
	}

	fmt.Println(outPath, "generated")
}
