package parser

import (
	"markdown_to_html/lexer"
	"testing"
)

func TestHeaderStatement(t *testing.T) {

	input := `# sentence
	# sentence2
	`

	l := lexer.New(input)
	p := New(l)

	program := p.ParseProgram()

	if program == nil {

	}
	// if program == nil {
	// 	t.Fatalf("ParseProgram() returned nil")
	// }
	// if len(program.Statements) != 3 {
	// 	t.Fatalf("Program.Statements does not contain 3 statements. got=%d", len(program.Statements))
	// }

	// tests := []struct {
	// 	expectedIdentifier string
	// }{
	// 	{"x"},
	// 	{"y"},
	// 	{"foobar"},
	// }

	// for i, tt := range tests {
	// 	stmt := program.Statements[i]
	// 	if !testLetStatement(t, stmt, tt.expectedIdentifier) {
	// 		return
	// 	}
	// }

}

func TestParagraphStatement(t *testing.T) {

	input := `sentence
	sentece2

	`

	l := lexer.New(input)
	p := New(l)

	program := p.ParseProgram()

	if program == nil {

	}
}

func TestOrderedList(t *testing.T) {
	input := `1. sentence
	2. sentece2

	`

	l := lexer.New(input)
	p := New(l)

	program := p.ParseProgram()

	if program == nil {

	}
}
