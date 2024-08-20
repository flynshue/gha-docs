package cmd

import (
	"bytes"
	"embed"
	"fmt"
	"os"
	"path/filepath"
	"text/template"

	"gopkg.in/yaml.v3"
)

//go:embed templates/*
var templates embed.FS
var templateDir = "templates/"

var funcMap = template.FuncMap{
	"getInputs":  getInputTable,
	"getOutputs": getOutputTable,
	"genUsage":   genUsage,
}

type Action struct {
	Name        string            `yaml:"name"`
	Use         string            `yaml:"use"`
	Description string            `yaml:"description,omitempty"`
	Inputs      map[string]Input  `yaml:"inputs,omitempty"`
	Outputs     map[string]Output `yaml:"outputs"`
}

type Output struct {
	Description string `yaml:"description,omitempty"`
}

type Input struct {
	Description string `yaml:"description,omitempty"`
	Required    bool   `yaml:"required,omitempty"`
	Default     string `yaml:"default,omitempty"`
	InputType   string `yaml:"type,omitempty"`
}

func readActionFile(file string) (*Action, error) {
	b, err := os.ReadFile(file)
	if err != nil {
		return nil, err
	}
	action := &Action{}
	err = yaml.Unmarshal(b, &action)
	action.setInputType()
	action.Use = getBasePath(file)
	return action, err
}

func getBasePath(file string) string {
	absPath, err := filepath.Abs(file)
	if err != nil {
		return ""
	}
	dir, _ := filepath.Split(absPath)
	return filepath.Base(dir)
}

func writeDocs(file string, data []byte) error {
	return os.WriteFile(file, data, 0644)
}

func (a *Action) setInputType() {
	for name, input := range a.Inputs {
		if input.InputType == "" {
			input.InputType = "string"
			a.Inputs[name] = input
		}
	}
}

func (a *Action) generateDocs() ([]byte, error) {
	tplFile := "doc.md"
	t, err := template.New("").Funcs(funcMap).ParseFS(templates, templateDir+tplFile)
	if err != nil {
		return nil, err
	}
	buf := &bytes.Buffer{}
	if err := t.ExecuteTemplate(buf, tplFile, a); err != nil {
		return nil, err
	}
	return buf.Bytes(), nil
}

func getInputTable(inputs map[string]Input) string {
	buf := &bytes.Buffer{}
	buf.WriteString("| Name | Type | Default | Description | Required |\n")
	buf.WriteString("| --- | --- | --- | --- | --- |\n")
	for name, input := range inputs {
		buf.WriteString(fmt.Sprintf("| %s | %s | %s | %s | %v |\n", name, input.InputType, input.Default, input.Description, input.Required))
	}
	return buf.String()
}

func getOutputTable(outputs map[string]Output) string {
	buf := &bytes.Buffer{}
	buf.WriteString("| Name | Description |\n")
	buf.WriteString("| --- | --- |\n")
	for name, output := range outputs {
		buf.WriteString(fmt.Sprintf("| %s | %s |\n", name, output.Description))
	}
	return buf.String()
}

func genUsage(a Action) string {
	usage := fmt.Sprintf(`
- name: %s
  uses: %s
  with:`, a.Name, a.Use)
	buf := &bytes.Buffer{}
	buf.WriteString(usage)
	for input, value := range a.Inputs {
		buf.WriteString(fmt.Sprintf("\n    %s: %s", input, value.Default))
	}
	return buf.String()
}
