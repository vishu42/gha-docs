package impl

import (
	"html/template"
	"os"

	_ "embed"

	f "github.com/vishu42/gha-docs/pkg/file"

	"github.com/spf13/cobra"
	"gopkg.in/yaml.v2"
)

//go:embed README.md.tpl
var tmpl []byte

type RootOpts struct {
	Debug      bool
	ActionFile string
}

func RunRoot(cmd *cobra.Command, args []string, o *RootOpts) {
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

		// custom field
		Documentation string
	}

	fileName := o.ActionFile

	// open yaml file
	file, err := os.ReadFile(fileName)
	if err != nil {
		panic(err)
	}

	osfile, err := os.OpenFile(fileName, os.O_RDONLY, 0o644)
	if err != nil {
		panic(err)
	}
	content, err := f.ReadYamlComments(osfile)
	if err != nil {
		panic(err)
	}

	// print content
	println(content)

	action.Documentation = content

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
