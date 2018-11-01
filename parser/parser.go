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
	case token.OLIST:
		return p.parseOrderedlist()
	case token.DASH:
		return p.parseUnorderedlist()
	default:
		return p.parseParagraph()
	}
}

func (p *Parser) parseHeaderStatement() *ast.HeaderStatement {
	stmt := &ast.HeaderStatement{Token: p.curToken}
	OSentence := p.parseOneline()
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

	for p.iftokenissentence(p.curToken.Type) {
		stmt.Sentences = append(stmt.Sentences, p.parseOneline())
	}

	return stmt
}

func (p *Parser) parseOneline() ast.SentenceOneline {
	stmt := ast.SentenceOneline{}

	for {

		switch p.curToken.Type {
		case token.SENTENCE:
			chunk := ast.Chunk{Token: p.curToken, Kind: "normal"}
			stmt.Chunks = append(stmt.Chunks, chunk)
		case token.BOLD:
			p.nextToken()
			chunk := ast.Chunk{Token: p.curToken, Kind: "bold"}
			stmt.Chunks = append(stmt.Chunks, chunk)
			p.nextToken()
			if p.curToken.Type != token.BOLD {
				log.Fatalf("error")
			}
		case token.TICK:
			p.nextToken()
			chunk := ast.Chunk{Token: p.curToken, Kind: "code"}
			stmt.Chunks = append(stmt.Chunks, chunk)
			p.nextToken()
			if p.curToken.Type != token.TICK {
				log.Fatalf("error")
			}
		case token.EXACLAMATION:
			p.nextToken()
			if p.curToken.Type != token.LSQUARE {
				log.Fatalf("error")
			}

			links := ast.Links{}
			p.nextToken()
			if p.curToken.Type != token.SENTENCE {
				log.Fatalf("error")
			}

			links.Name = p.curToken.Literal

			p.nextToken()
			if p.curToken.Type != token.RSQUARE {
				log.Fatalf("error")
			}

			p.nextToken()
			if p.curToken.Type != token.LPAREN {
				log.Fatalf("error")
			}

			p.nextToken()
			if p.curToken.Type != token.SENTENCE {
				log.Fatalf("error")
			}

			links.Url = p.curToken.Literal

			p.nextToken()
			if p.curToken.Type != token.RPAREN {
				log.Fatalf("error")
			}

			chunk := ast.Chunk{Kind: "image", Links: links}
			stmt.Chunks = append(stmt.Chunks, chunk)
		case token.LSQUARE:
			links := ast.Links{}
			p.nextToken()
			if p.curToken.Type != token.SENTENCE {
				log.Fatalf("error")
			}

			links.Name = p.curToken.Literal

			p.nextToken()
			if p.curToken.Type != token.RSQUARE {
				log.Fatalf("error")
			}

			p.nextToken()
			if p.curToken.Type != token.LPAREN {
				log.Fatalf("error")
			}

			p.nextToken()
			if p.curToken.Type != token.SENTENCE {
				log.Fatalf("error")
			}

			links.Url = p.curToken.Literal

			p.nextToken()
			if p.curToken.Type != token.RPAREN {
				log.Fatalf("error")
			}

			chunk := ast.Chunk{Kind: "link", Links: links}
			stmt.Chunks = append(stmt.Chunks, chunk)

		}

		if p.expectPeek(token.NEWLINE) {
			break
		}
		if p.expectPeek(token.EOF) {
			break
		}
	}
	p.nextToken()
	return stmt
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

func (p *Parser) iftokenissentence(tokentype token.TokenType) bool {
	if tokentype == token.SENTENCE || tokentype == token.TICK || tokentype == token.BOLD || tokentype == token.LSQUARE || tokentype == token.EXACLAMATION {
		return true
	} else {
		return false
	}
}
