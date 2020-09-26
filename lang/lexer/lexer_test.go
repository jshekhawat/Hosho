package lexer

import (
	"bytes"
	"testing"

	"github.com/jshekhawat/pogo/lang/token"
)

func TestNonLiterals(t *testing.T) {

	tests := []struct {
		expectedType   token.Type
		expectedLexeme string
	}{
		{token.ASSIGN, "="},
	}

	lex := New(bytes.NewBufferString(`=`))

	for i, tt := range tests {
		tok := lex.NextToken()

		if tok.Type != tt.expectedType {
			t.Fatalf("tests[%d] - expected=%q, got=%q", i, tt.expectedType, tok.Type)
		}

	}

}
