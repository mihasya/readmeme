package readmegen

import (
	"text/template"
	"io"
)

type Readme struct {
	Project string
	Who string
	What string
	When string
	Where string
	Why string
}

var readmeTemplate string =
`# {{.Project}}

## Who?
{{.Who}}

## What?
{{.What}}

## When?
{{.When}}

## Where?
{{.Where}}

## Why?
{{.Why}}
`

func Render(w io.Writer, r Readme) error {
	tpl, err := template.New("markdown").Parse(readmeTemplate)
	if err != nil {
		return err
	}
	err = tpl.Execute(w, r)
	if err != nil {
		return err
	}
	return nil
}
