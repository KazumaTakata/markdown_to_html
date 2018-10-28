package lexer

import "testing"
import "markdown_to_html/token"

func TestNextToken(t *testing.T) {
	input := `# sentence 1 eee
	## sentent 3 
	### sentent 4 
	`

	tests := []struct {
		expectedType    token.TokenType
		expectedLiteral string
	}{
		{token.HEADER1, "#"},
		{token.SENTENCE, "sentence 1 eee"},
		{token.NEWLINE, "\n"},
		{token.HEADER2, "##"},
		{token.SENTENCE, "sentent 3 "},
		{token.NEWLINE, "\n"},
		{token.HEADER3, "###"},
		{token.SENTENCE, "sentent 4 "},
		{token.NEWLINE, "\n"},
	}

	l := New(input)

	for i, tt := range tests {
		tok := l.NextToken()

		if tok.Type != tt.expectedType {
			t.Fatalf("tests[%d] - tokenType wrong. expected=%q, got=%q", i, tt.expectedType, tok.Type)
		}

		if tok.Literal != tt.expectedLiteral {
			t.Fatalf("tests[%d] - literal wrong. expected=%q, got=%q", i, tt.expectedLiteral, tok.Literal)
		}
	}

}
