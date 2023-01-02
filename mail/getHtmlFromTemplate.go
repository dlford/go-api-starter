package mail

import (
	"bytes"
	"context"
	"errors"
	"fmt"
	"text/template"

	"github.com/Boostport/mjml-go"
)

func getHtmlFromTemplate(filename string, data interface{}) (string, error) {
	tmpl, err := template.ParseFiles(filename)
	if err != nil {
		panic(err)
	}
	var tmplBuff bytes.Buffer
	tmpl.Execute(&tmplBuff, data)
	html, err := mjml.ToHTML(context.Background(), tmplBuff.String(), mjml.WithMinify(true))
	var mjmlError mjml.Error
	if errors.As(err, &mjmlError) {
		fmt.Println(mjmlError.Message)
		fmt.Println(mjmlError.Details)
		return "", err
	}

	return html, nil
}
