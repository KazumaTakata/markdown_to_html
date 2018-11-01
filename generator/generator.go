package generator

import (
	"markdown_to_html/ast"
	"markdown_to_html/taggenerator"
)

func Generator(node interface{}) []string {
	switch node := node.(type) {
	case *ast.Program:
		return evalStatements(node.Statements)
	case *ast.HeaderStatement:
		return evalHeader(node)
	case *ast.Paragraph:
		return evalParagraph(node.Sentences)
	case *ast.UnOrderedList:
		return evalUnOrderedList(node.Sentences)

	}

	return nil
}

func evalStatements(stmt []interface{}) []string {
	htmllist := []string{}
	for _, statement := range stmt {
		result := Generator(statement)
		htmllist = append(htmllist, result...)
	}
	return htmllist
}

func evalHeader(node *ast.HeaderStatement) []string {
	htmllist := []string{}
	content := evalChunks(node.Sentence.Chunks)

	tagstring := taggenerator.GenerateTagString("h1", content)
	htmllist = append(htmllist, tagstring)

	return htmllist
}

func evalParagraph(node []ast.SentenceOneline) []string {
	htmllist := []string{}
	content := ""
	for _, sentence := range node {
		content += evalChunks(sentence.Chunks)
		content += "\n"
	}

	tagstring := taggenerator.GenerateTagString("p", content)
	htmllist = append(htmllist, tagstring)
	return htmllist
}

func evalUnOrderedList(node []ast.SentenceOneline) []string {
	listitems := []string{}
	for _, sentence := range node {
		content := evalChunks(sentence.Chunks)
		listitems = append(listitems, content)
	}

	listitems = append(listitems, taggenerator.GenerageUnOrderedString(listitems))

	return listitems
}

func evalChunks(chunks []ast.Chunk) string {
	content := ""
	for _, chunk := range chunks {
		switch chunk.Kind {
		case "bold":
			tmpstring := "<strong>"
			tmpstring += chunk.Token.Literal
			tmpstring += "</strong>"
			content += tmpstring
		case "code":
			tmpstring := "<code>"
			tmpstring += chunk.Token.Literal
			tmpstring += "</code>"
			content += tmpstring
		case "normal":
			content += chunk.Token.Literal
		case "link":
			content = taggenerator.GenerateTagStringwithAttri("a", chunk.Links.Name, "href", chunk.Links.Url)
		case "image":
			content = taggenerator.GenerateImageTag(chunk.Links.Url, chunk.Links.Name)
		}

	}

	return content
}
