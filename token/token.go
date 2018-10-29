package token

type TokenType string

type Token struct {
	Type    TokenType
	Literal string
}

const (
	HEADER1  = "#"
	HEADER2  = "##"
	HEADER3  = "###"
	HEADER4  = "####"
	HEADER5  = "#####"
	SENTENCE = "SENTENCE"
	NEWLINE  = "NEWLINE"
	BOLD     = "BOLD"
	ITALIC   = "ITALIC"
	OLIST    = "OLIST"
	DASH     = "-"
	EOF      = "EOF"
	PARA     = "PARA"
)
