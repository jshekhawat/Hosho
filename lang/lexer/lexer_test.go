package lexer

import (
	"bytes"
	"fmt"
	"testing"

	"github.com/jshekhawat/pogo/lang/token"
)

func TestNonLiterals(t *testing.T) {

	// tests := []struct {
	// 	expectedType   token.Type
	// 	expectedLexeme string
	// }{
	// 	{token.PLUS, "+"},
	// 	{token.ASSIGN, ":="},
	// 	{token.GTE, ">="},
	// 	{token.EQUAL, "="},
	// 	{token.LTE, "<="},
	// 	{token.LT, "<"},
	// }

	// lex := New(bytes.NewBufferString(`+:=><=<`))

	// for i, tt := range tests {
	// 	tok := lex.NextToken()

	// 	t.Logf("Expected Type %x - Got %x \n", tt.expectedType, tok.Type)
	// 	t.Logf("Expected Lexeme %q - Got %q \n", tt.expectedLexeme, tok.Lexeme)

	// 	if tok.Type != tt.expectedType {

	// 		t.Errorf("tests[%d] - expected=%q, got=%q", i, tt.expectedType, tok.Type)
	// 	}

	// 	if tok.Lexeme != tt.expectedLexeme {
	// 		t.Errorf("tests[%d] - expected=%q, got=%q", i, tt.expectedLexeme, tok.Lexeme)
	// 	}

	// }

}

func TestIdentifiers(t *testing.T) {
	tests := []struct {
		expectedType   token.Type
		expectedLexeme string
	}{

		{token.LET, "let"},
	}

	lex := New(bytes.NewBufferString(`let a = 5`))
	lex.Tokenise()

	for i, tt := range tests {
		tok := lex.tokens[0]

		fmt.Printf("Expected Type %q - Got %q \n", tt.expectedType, tok.Type)

		if tok.Type == tt.expectedType {
			for _, tok1 := range lex.tokens {
				t.Errorf("Token Type - %q | Token Value - %q \n", tok1.Type, tok1.Lexeme)

			}
		}

		if tok.Lexeme != tt.expectedLexeme {
			t.Errorf("tests[%d] - expected=%q, got=%q", i, tt.expectedLexeme, tok.Lexeme)
		}

	}
}
