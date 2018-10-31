package ast

import "markdown_to_html/token"

type OnelineSentence interface {
	onelinesentence()
}

type Chunk interface {
	chunk()
}

type Program struct {
	Statements []interface{}
}

type HeaderStatement struct {
	Token    token.Token
	Sentence OnelineSentence
}

type Paragraph struct {
	Token     token.Token
	Sentences []OnelineSentence
}

type SentenceOneline struct {
	Chunks []Chunk
}

func (o *SentenceOneline) onelinesentence() {}

type BoldChunk struct {
	Token token.Token
	Chunk Chunk
}

func (b *BoldChunk) chunk() {
}

type NormalChunk struct {
	Token token.Token
}

func (n *NormalChunk) chunk() {
}

type OrderedList struct {
	Sentences []OnelineSentence
}

type UnOrderedList struct {
	Sentences []OnelineSentence
}
