package parser

import (
	"log"
	"markdown_to_html/ast"
	"markdown_to_html/lexer"
	"markdown_to_html/token"
)

type Parser struct {
	l         *lexer.Lexer
	curToken  token.Token
	peekToken token.Token
}

func New(l *lexer.Lexer) *Parser {
	p := &Parser{l: l}
	p.nextToken()
	p.nextToken()
	return p
}

func (p *Parser) ParseProgram() *ast.Program {
	program := &ast.Program{}

	for p.curToken.Type != token.EOF {
		for p.curToken.Type == token.NEWLINE {
			p.nextToken()
		}
		stmt := p.parseStatement()
		if stmt != nil {
			program.Statements = append(program.Statements, stmt)
		}
	}
	return program
}

func (p *Parser) nextToken() {
	p.curToken = p.peekToken
	p.peekToken = p.l.NextToken()
}

func (p *Parser) parseStatement() interface{} {
	switch p.curToken.Type {
	case token.HEADER1:
		return p.parseHeaderStatement()
	case token.HEADER2:
		return p.parseHeaderStatement()
	case token.HEADER3:
		return p.parseHeaderStatement()
	case token.HEADER4:
		return p.parseHeaderStatement()
	case token.HEADER5:
		return p.parseHeaderStatement()
	case token.SENTENCE:
		return p.parseParagraph()
	case token.OLIST:
		return p.parseOrderedlist()
	case token.DASH:
		return p.parseUnorderedlist()
	default:
		return nil
	}
}

func (p *Parser) parseHeaderStatement() *ast.HeaderStatement {
	stmt := &ast.HeaderStatement{Token: p.curToken}
	OSentence := &ast.SentenceOneline{}
	for !p.expectPeek(token.NEWLINE) {
		p.nextToken()
		switch p.curToken.Type {
		case token.SENTENCE:
			chunk := &ast.NormalChunk{Token: p.curToken}
			OSentence.Chunks = append(OSentence.Chunks, chunk)
		case token.BOLD:
			sentence := p.parseBold()
			OSentence.Chunks = append(OSentence.Chunks, sentence)
		}
	}
	stmt.Sentence = OSentence
	p.nextToken()

	return stmt
}

func (p *Parser) parseUnorderedlist() *ast.UnOrderedList {
	stmt := &ast.UnOrderedList{}

	for p.curToken.Type == token.DASH {
		stmt.Sentences = append(stmt.Sentences, p.parseOneline())
	}
	return stmt
}

func (p *Parser) parseOrderedlist() *ast.OrderedList {
	stmt := &ast.OrderedList{}

	for p.curToken.Type == token.OLIST {
		stmt.Sentences = append(stmt.Sentences, p.parseOneline())
	}
	return stmt
}

func (p *Parser) parseParagraph() *ast.Paragraph {
	stmt := &ast.Paragraph{}

	for p.curToken.Type == token.SENTENCE {
		stmt.Sentences = append(stmt.Sentences, p.parseOneline())
	}

	return stmt
}

func (p *Parser) parseOneline() *ast.SentenceOneline {
	stmt := &ast.SentenceOneline{}
	chunk := &ast.BoldChunk{Token: p.curToken}
	stmt.Chunks = append(stmt.Chunks, chunk)

	for !p.expectPeek(token.NEWLINE) {
		p.nextToken()
		switch p.curToken.Type {
		case token.SENTENCE:
			chunk := &ast.BoldChunk{Token: p.curToken}
			stmt.Chunks = append(stmt.Chunks, chunk)
		case token.BOLD:
			chunk := p.parseBold()
			stmt.Chunks = append(stmt.Chunks, chunk)
		}
	}

	p.nextToken()

	return stmt

}

func (p *Parser) parseBold() *ast.BoldChunk {
	bold := &ast.BoldChunk{Token: p.curToken}

	if !p.expectPeek(token.SENTENCE) {
		log.Fatal("syntax error")
	}

	chunk := &ast.NormalChunk{Token: p.curToken}

	bold.Chunk = chunk

	if !p.expectPeek(token.BOLD) {
		log.Fatal("syntax error")
	}

	return bold
}

func (p *Parser) expectPeek(t token.TokenType) bool {
	if p.peekTokenIs(t) {
		p.nextToken()
		return true
	} else {
		return false
	}
}

func (p *Parser) peekTokenIs(t token.TokenType) bool {
	return p.peekToken.Type == t
}
