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
		stmt := p.parseStatement()
		if stmt != nil {
			program.Statements = append(program.Statements, stmt)
		}

		p.nextToken()
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
	default:
		return nil
	}
}

func (p *Parser) parseHeaderStatement() *ast.HeaderStatement {
	stmt := &ast.HeaderStatement{Token: p.curToken}

	for !p.expectPeek(token.NEWLINE) {
		p.nextToken()
		switch p.curToken.Type {
		case token.SENTENCE:
			sentence := ast.Sentence{Token: p.curToken}
			stmt.Body = append(stmt.Body, sentence)
		case token.BOLD:
			sentence := p.parseBold()
			stmt.Body = append(stmt.Body, sentence)
		}
	}

	return stmt
}

func (p *Parser) parseParagraph() *ast.Paragraph {
	stmt := &ast.Paragraph{Token: token.Token{Type: token.PARA}}
	sentence := ast.Sentence{Token: p.curToken}
	stmt.Body = append(stmt.Body, sentence)

	for !p.expectPeek(token.NEWLINE) {
		p.nextToken()
		switch p.curToken.Type {
		case token.SENTENCE:
			sentence := ast.Sentence{Token: p.curToken}
			stmt.Body = append(stmt.Body, sentence)
		case token.BOLD:
			sentence := p.parseBold()
			stmt.Body = append(stmt.Body, sentence)
		}
	}

	return stmt
}

func (p *Parser) parseBold() *ast.BoldSentence {
	bold := &ast.BoldSentence{Token: p.curToken}

	if !p.expectPeek(token.SENTENCE) {
		log.Fatal("syntax error")
	}

	sentence := ast.Sentence{Token: p.curToken}

	bold.Sentence = append(bold.Sentence, sentence)

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
