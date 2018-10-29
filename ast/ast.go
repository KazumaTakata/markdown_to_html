package ast

import "markdown_to_html/token"

type Program struct {
	Statements []interface{}
}

type Sentence struct {
	Token token.Token
}

type HeaderStatement struct {
	Token token.Token
	Body  []interface{}
}

type Paragraph struct {
	Token token.Token
	Body  []interface{}
}

type BoldSentence struct {
	Token    token.Token
	Sentence []Sentence
}
