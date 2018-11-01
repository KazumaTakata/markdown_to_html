package generator

import (
	"fmt"
	"markdown_to_html/lexer"
	"markdown_to_html/parser"
	"testing"
)

func TestEvalIntegerExpression(t *testing.T) {
	tests := []struct {
		input string
	}{
		{`# sentence
	# sentence2
	`},
	}

	for _, tt := range tests {
		evaluated := testEval(tt.input)
		fmt.Sprint(evaluated[0])
		// testIntegerObject(t, evaluated, tt.expected)
	}
}

func testEval(input string) []string {
	l := lexer.New(input)
	p := parser.New(l)
	program := p.ParseProgram()

	return Generator(program)
}

func TestEvalExpression(t *testing.T) {
	tests := []struct {
		input string
	}{
		{`sentence
	sentence2
	`},
	}

	for _, tt := range tests {
		evaluated := testEval(tt.input)
		fmt.Sprint(evaluated[0])
		// testIntegerObject(t, evaluated, tt.expected)
	}
}

func TestUListExpression(t *testing.T) {
	tests := []struct {
		input string
	}{
		{`- sentence
	- sentence2
	`},
	}

	for _, tt := range tests {
		evaluated := testEval(tt.input)
		fmt.Sprint(evaluated[0])
		// testIntegerObject(t, evaluated, tt.expected)
	}
}

func TestBoldExpression(t *testing.T) {
	tests := []struct {
		input string
	}{
		{`** boldtext **
	`},
	}

	for _, tt := range tests {
		evaluated := testEval(tt.input)
		fmt.Sprint(evaluated[0])
		// testIntegerObject(t, evaluated, tt.expected)
	}
}

func TestCodeExpression(t *testing.T) {
	tests := []struct {
		input string
	}{
		{"`code`"},
	}

	for _, tt := range tests {
		evaluated := testEval(tt.input)
		fmt.Sprint(evaluated[0])
		// testIntegerObject(t, evaluated, tt.expected)
	}
}

func TestLinksExpression(t *testing.T) {
	tests := []struct {
		input string
	}{
		{"[linkname](linkurl)"},
	}

	for _, tt := range tests {
		evaluated := testEval(tt.input)
		fmt.Sprint(evaluated[0])
		// testIntegerObject(t, evaluated, tt.expected)
	}
}

func TestImagesExpression(t *testing.T) {
	tests := []struct {
		input string
	}{
		{"![linkname](linkurl)"},
	}

	for _, tt := range tests {
		evaluated := testEval(tt.input)
		fmt.Sprint(evaluated[0])
		// testIntegerObject(t, evaluated, tt.expected)
	}
}
