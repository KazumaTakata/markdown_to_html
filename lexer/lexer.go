package lexer

import (
	"fmt"
	"markdown_to_html/token"
	"strings"
)

type Lexer struct {
	input        string
	position     int
	readPosition int
	ch           byte
	ch_str       string
	header_num   int
}

func (l *Lexer) readChar() {
	if l.readPosition >= len(l.input) {
		l.ch = 0
	} else {
		l.ch = l.input[l.readPosition]
		l.ch_str = string(l.input[l.readPosition])
	}
	l.position = l.readPosition

	l.readPosition += 1
}

func (l *Lexer) readUntil(ch byte) {

}

func (l *Lexer) PrintCurrChar() {
	fmt.Println(string(l.ch))
}

func (l *Lexer) NextToken() token.Token {
	l.skipWhitespace()
	var tok token.Token

	switch l.ch {
	case '#':
		l.header_num += 1
		l.CoundHeader()
		if l.peekChar() == ' ' {
			l.readChar()
			switch l.header_num {
			case 1:
				tok = newToken(token.HEADER1, '#')
			case 2:
				tok = token.Token{Type: token.HEADER2, Literal: "##"}
			case 3:
				tok = token.Token{Type: token.HEADER3, Literal: "###"}
			case 4:
				tok = token.Token{Type: token.HEADER4, Literal: "####"}
			case 5:
				tok = token.Token{Type: token.HEADER5, Literal: "#####"}
			}
			l.readChar()
		} else {
			tok.Type = token.SENTENCE
			tok.Literal = strings.Repeat("#", l.header_num) + l.readSentence()
		}
		l.header_num = 0
	case '\n':
		tok = newToken(token.NEWLINE, l.ch)
		l.readChar()
	case '*':
		if l.peekChar() == '*' {
			l.readChar()
			tok = token.Token{Type: token.BOLD, Literal: "**"}
			l.readChar()
		} else {
			l.readChar()
			tok = token.Token{Type: token.ITALIC, Literal: "*"}
		}
	case '-':
		tok = newToken(token.DASH, l.ch)
		l.readChar()
	case 0:
		tok.Literal = ""
		tok.Type = token.EOF
	default:
		if isDigit(l.ch) {
			if l.peekChar() == '.' {
				literal := string(l.ch) + string('.')
				tok = token.Token{Type: token.OLIST, Literal: literal}
				l.readChar()
				l.readChar()
			} else {
				tok.Type = token.SENTENCE
				tok.Literal = l.readSentence()
			}
		} else {
			tok.Type = token.SENTENCE
			tok.Literal = l.readSentence()
		}

	}
	return tok
}

func (l *Lexer) peekChar() byte {
	if l.readPosition >= len(l.input) {
		return 0
	} else {
		return l.input[l.readPosition]
	}
}

func (l *Lexer) CoundHeader() {
	for l.peekChar() == '#' {
		l.readChar()
		l.header_num += 1
	}
}

func New(input string) *Lexer {
	l := &Lexer{input: input, header_num: 0}
	l.readChar()
	return l
}

func (l *Lexer) readSentence() string {
	position := l.position
	for isLetter(l.ch) || isDigit(l.ch) {
		l.readChar()
	}

	return l.input[position:l.position]
}

func newToken(tokenType token.TokenType, ch byte) token.Token {
	return token.Token{Type: tokenType, Literal: string(ch)}
}

func isLetter(ch byte) bool {
	return 'a' <= ch && ch <= 'z' || 'A' <= ch && ch <= 'Z' || ch == ' ' || ch == '#'
}

func isDigit(ch byte) bool {
	return '0' <= ch && ch <= '9'
}

func (l *Lexer) skipWhitespace() {
	for l.ch == ' ' || l.ch == '\t' || l.ch == '\r' {
		l.readChar()
	}
}
