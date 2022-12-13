package main

import (
	"fmt"
	"os"
	"strconv"
	"text/template"
)

func main() {
	args := os.Args

	tmplPath := args[1]
	outPath := args[2]

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

	file, err := os.Create(outPath)
	if err != nil {
		fmt.Println(err)
		return
	}
	defer file.Close()

	tmplArgs := map[string]any{}
	values := args[3:]
	tmplArgs["values"] = values
	for i, value := range values {
		tmplArgs["value"+strconv.Itoa(i+1)] = value
	}

	err = tmpl.Execute(file, tmplArgs)
	if err != nil {
		fmt.Println(err)
		return
	}

	fmt.Println(outPath, "generated")
}
