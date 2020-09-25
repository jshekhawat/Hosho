package lexer

import (
	"io"
	"io/ioutil"

	"github.com/jshekhawat/pogo/lang/token"
)

//Lexer is the struct for scanning and generating tokens
type Lexer struct {
	source       []byte
	position     int
	readPosition int
	current      byte
	line         int
}

type error struct {
}

// New initialises the lexer for a given source
func New(source io.Reader) *Lexer {
	bs, _ := ioutil.ReadAll(source)
	l := &Lexer{source: bs}
	l.readChar()
	return l
}

func (l *Lexer) match(ch byte) bool {
	return ch == l.source[l.readPosition]
}

func (l *Lexer) readChar() {
	if l.readPosition >= len(l.source) {
		l.current = 0
	} else {
		l.current = l.source[l.readPosition]
	}
	l.position = l.readPosition
	l.readPosition++
}

func addToken(lexeme byte, tt token.Type) token.Token {
	return token.Token{
		Lexeme: string(lexeme),
		Type:   tt,
	}
}

func getString(l *Lexer) token.Token {
	start := l.position
	for l.current != '\'' {
		l.readChar()
	}
	return token.Token{
		Lexeme: string(l.source[start:l.position]),
		Type:   token.STRING,
	}
}

//NextToken allows you to iterate through all the tokens in the input stream
func (l *Lexer) NextToken() token.Token {

	l.readChar()
	tok := token.Token{
		Lexeme: "",
		Type:   token.EOF,
	}
	switch l.current {
	case '=':
		return addToken(l.current, token.EQUAL)
	case '>':
		if l.match('=') {
			l.readChar()
			return token.Token{
				Lexeme: ">=",
				Type:   token.GTE,
			}
		}
		return addToken(l.current, token.GT)
	case '<':
		if l.match('=') {
			l.readChar()
			return token.Token{
				Lexeme: "<=",
				Type:   token.LTE,
			}
		}
		return addToken(l.current, token.LTE)
	case '!':
		return addToken(l.current, token.BANG)
	case '+':
		return addToken(l.current, token.PLUS)
	case '-':
		return addToken(l.current, token.MINUS)
	case '/':
		return addToken(l.current, token.SLASH)
	case '*':
		return addToken(l.current, token.STAR)
	case '(':
		return addToken(l.current, token.LPAREN)
	case ')':
		return addToken(l.current, token.RPAREN)
	case '[':
		return addToken(l.current, token.LBRACKET)
	case ']':
		return addToken(l.current, token.RBRACKET)
	case '{':
		return addToken(l.current, token.LBRACE)
	case '}':
		return addToken(l.current, token.RBRACE)
	case ':':
		if l.match('=') {
			return token.Token{
				Lexeme: ":=",
				Type:   token.ASSIGN,
			}
		}
		return addToken(l.current, token.COLON)
	case '^':
		return addToken(l.current, token.CARET)
	case '.':
		return addToken(l.current, token.DOT)
	case '\'':
		return getString(l)
	case '#':
		return addToken(l.current, token.DOT)
	}
	return tok
}
