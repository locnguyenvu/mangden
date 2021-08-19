package template

import (
	"bytes"
	"embed"
	"html/template"
	"log"
	"strings"
)

//go:embed templates/common/*.html
var commonTemplateFiles embed.FS

//go:embed templates/views/*.html
var viewTemplateFiles embed.FS

type Engine struct {
	templates map[string]*template.Template
}

func NewEngine() *Engine {
	e := Engine{
		templates: make(map[string]*template.Template, 0),
	}
	e.ParseTemplate()
	return &e
}

func (e *Engine) ParseTemplate() error {
	var commonContent strings.Builder
	dirEntries, err := commonTemplateFiles.ReadDir("templates/common")
	if err != nil {
		return err
	}

	for _, dirEntry := range dirEntries {
		fileData, err := commonTemplateFiles.ReadFile("templates/common/" + dirEntry.Name())
		if err != nil {
			return err
		}
		commonContent.Write(fileData)
	}

	dirEntries, err = viewTemplateFiles.ReadDir("templates/views")
	if err != nil {
		return err
	}
	for _, dirEntry := range dirEntries {
		templateName := dirEntry.Name()

		fileData, err := viewTemplateFiles.ReadFile("templates/views/" + dirEntry.Name())
		if err != nil {
			return err
		}

		var viewContent strings.Builder
		viewContent.WriteString(commonContent.String())
		viewContent.Write(fileData)
		e.templates[templateName] = template.Must(template.New("main").Parse(viewContent.String()))
	}

	return nil
}

func (e Engine) Render(templateName string, data map[string]interface{}) []byte {
	tpl, ok := e.templates[templateName]
	if !ok {
		log.Fatal("View does not exist")
	}
	var b bytes.Buffer
	err := tpl.ExecuteTemplate(&b, "base", data)
	if err != nil {
		log.Fatal(err)
	}
	return b.Bytes()
}
