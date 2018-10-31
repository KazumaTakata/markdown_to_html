package generator

import (
	"markdown_to_html/lexer"
	"markdown_to_html/parser"
	"testing"
	"writinginterpreter/object"
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
		// testIntegerObject(t, evaluated, tt.expected)
	}
}

func testEval(input string) object.Object {
	l := lexer.New(input)
	p := parser.New(l)
	program := p.ParseProgram()

	return Eval(program, env)
}
