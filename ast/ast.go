package ast

import "markdown_to_html/token"

type OnelineSentence interface {
	onelinesentence()
}

type Program struct {
	Statements []interface{}
}

type HeaderStatement struct {
	Token    token.Token
	Sentence SentenceOneline
}

type Paragraph struct {
	Token     token.Token
	Sentences []SentenceOneline
}

type SentenceOneline struct {
	Chunks []Chunk
}

func (o *SentenceOneline) onelinesentence() {}

type Chunk struct {
	Token token.Token
	Kind  string
	Links Links
}

type OrderedList struct {
	Sentences []SentenceOneline
}

type UnOrderedList struct {
	Sentences []SentenceOneline
}

type Links struct {
	Name string
	Url  string
}
