package main

import (
	"fmt"
	"os"
	"text/template"

	"gopkg.in/yaml.v2"

	_ "embed"
)

//go:embed README.md.tpl
var tmpl []byte

func main() {
	// parse yaml file into a struct
	var action struct {
		Name        string `yaml:"name"`
		Description string `yaml:"description"`
		Inputs      map[string]struct {
			Description string `yaml:"description"`
			Required    bool   `yaml:"required"`
			Default     string `yaml:"default"`
		} `yaml:"inputs"`

		Outputs map[string]struct {
			Description string `yaml:"description"`
		} `yaml:"outputs"`

		Author string `yaml:"author"`
	}

	var file []byte
	fileName := ""

	if len(os.Args) >= 2 {
		fileName = os.Args[1]
	}

	// if no file name is provided, use the default
	if len(os.Args) < 2 {
		fileName = "action.yml"
		fmt.Println("No file name provided, using default action.yml")
	}

	// open yaml file
	file, err := os.ReadFile(fileName)
	if err != nil {
		panic(err)
	}

	// parse yaml file into action struct
	err = yaml.Unmarshal(file, &action)
	if err != nil {
		panic(err)
	}

	// execute http template
	t, err := template.New("readme").Parse(string(tmpl))
	if err != nil {
		panic(err)
	}

	// create a readme.md file
	f, err := os.Create("README.md")
	if err != nil {
		panic(err)
	}

	err = t.Execute(f, action)
	if err != nil {
		panic(err)
	}
}
