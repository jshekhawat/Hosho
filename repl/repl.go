package repl

import (
	"bufio"
	"bytes"
	"io"

	"github.com/jshekhawat/hosho/lang/lexer"
)

/*
Repl allows for:

1) file based eval
2) expression eval

*/

// REPL is the main data structue for repl context
type REPL struct {
	out   io.Writer
	input io.Reader
}

// New returns a new REPL instance
func New(input io.Reader, out io.Writer) *REPL {
	return &REPL{
		out:   out,
		input: input,
	}

}

//Cor stands for Carry on Repl-ling, the loop part of REPL
func (r *REPL) Cor() {

	scanner := bufio.NewScanner(r.input)

	for {

		scanned := scanner.Scan()
		if !scanned {
			return
		}

		line := scanner.Text()

		l := lexer.New(bytes.NewBufferString(line))
		l.Tokenise()

		for _, t := range l.Tokens {
			r.out.Write([]byte("Lexeme: " + t.Lexeme + "\n"))
			r.out.Write([]byte("Type: " + t.Type + "\n"))
		}

	}

}
