package taggenerator

import (
	"bytes"
	"strings"
	"text/template"
)

type TerminalTag struct {
	Tag     string
	Content string
}

func GenerateTagString(tagtype string, content string) string {

	td := TerminalTag{tagtype, content}

	t, err := template.New("todos").Parse("<{{.Tag}}> {{ .Content}} </{{.Tag}}>")
	if err != nil {
		panic(err)
	}
	var tpl bytes.Buffer
	err = t.Execute(&tpl, td)
	if err != nil {
		panic(err)
	}
	result := tpl.String()

	return result
}

func GenerageUnOrderedString(contents []string) string {

	stringlist := []string{}
	stringlist = append(stringlist, "<ul>")
	stringlist = append(stringlist, "\n")

	for _, content := range contents {
		stringlist = append(stringlist, GenerateTagString("li", content))
		stringlist = append(stringlist, "\n")
	}

	stringlist = append(stringlist, "</ul>")
	stringlist = append(stringlist, "\n")

	return strings.Join(stringlist[:], " ")
}

type TerminalTagWithAttri struct {
	Tag        string
	Content    string
	AttriName  string
	AttriValue string
}

func GenerateTagStringwithAttri(tagtype string, content string, attrname string, attrvalue string) string {

	td := TerminalTagWithAttri{tagtype, content, attrname, attrvalue}

	t, err := template.New("todos").Parse("<{{.Tag}} {{.AttriName}}=\"{{.AttriValue}}\">{{.Content}}</{{.Tag}}>")
	if err != nil {
		panic(err)
	}
	var tpl bytes.Buffer
	err = t.Execute(&tpl, td)
	if err != nil {
		panic(err)
	}
	result := tpl.String()

	return result
}

type ImageTag struct {
	Url   string
	Attri string
}

func GenerateImageTag(src string, attri string) string {

	td := ImageTag{src, attri}

	t, err := template.New("todos").Parse("<img src=\"{{.Url}}\" alt=\"{{.Attri}}\"/>")
	if err != nil {
		panic(err)
	}
	var tpl bytes.Buffer
	err = t.Execute(&tpl, td)
	if err != nil {
		panic(err)
	}
	result := tpl.String()

	return result

}
