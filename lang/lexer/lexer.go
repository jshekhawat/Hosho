package lexer

import (
	"io"
	"io/ioutil"

	"github.com/jshekhawat/hosho/lang/token"
)

//Lexer only supports ascii
type Lexer struct {
	current  rune
	source   []byte
	Tokens   []token.Token
	position int //the position of the current byte
	line     int
	errors   []Error
}

//Msg is the message string for the thrown error
type Msg string

//Error is the error thrown while lexing
type Error struct {
	Msg  string
	line int
	col  int
}

//New returns an initialised lexer
func New(r io.Reader) *Lexer {
	var bs, _ = ioutil.ReadAll(r)
	l := &Lexer{
		current:  -1,
		source:   bs,
		position: 0,
		line:     1,
	}
	l.nextToken()
	return l
}

//Tokenise takes an input stream and tokenises it, Tokens are stored in the Tokens slice
func (l *Lexer) Tokenise() {

	for !l.isAtEnd() {
		l.nextToken()
	}
}

func (l *Lexer) advance() {
	//this is so dumb
	if l.isAtEnd() {
		l.current = -1
		l.position++
		return
	}
	l.current = rune(l.source[l.position])
	l.position++
}

func (l *Lexer) match(ch byte) bool {
	if l.isAtEnd() {
		return false
	}
	return ch == l.source[l.position]
}

func (l *Lexer) addToken(ch rune, tt token.Type) {
	l.Tokens = append(l.Tokens, token.Token{
		Lexeme: string(ch),
		Type:   tt,
	})
}

func (l *Lexer) addStringToken(str string, tt token.Type) {
	l.Tokens = append(l.Tokens, token.Token{
		Lexeme: str,
		Type:   tt,
	})
}

func (l *Lexer) peek() rune {
	if l.isAtEnd() {
		return -1
	}
	return rune(l.source[l.position])
}

func (l *Lexer) nextToken() {

	l.advance()

	switch l.current {
	case -1:
		l.addToken(l.current, token.EOF)
	case '=':
		l.addToken(l.current, token.EQUAL)
	case '>':
		if l.match('=') {
			l.addStringToken(">=", token.GTE)
		} else {
			l.addToken(l.current, token.GT)
		}
	case '<':
		if l.match('=') {
			l.addStringToken("<=", token.LTE)
		} else {
			l.addToken(l.current, token.LT)
		}
	case '+':
		l.addToken(l.current, token.PLUS)
	case '-':
		l.addToken(l.current, token.MINUS)
	case '*':
		l.addToken(l.current, token.STAR)
	case '/':
		l.addToken(l.current, token.SLASH)
	case '(':
		l.addToken(l.current, token.LPAREN)
	case ')':
		l.addToken(l.current, token.RPAREN)
	case '[':
		l.addToken(l.current, token.LBRACKET)
	case ']':
		l.addToken(l.current, token.RBRACKET)
	case '{':
		l.addToken(l.current, token.LBRACE)
	case '}':
		l.addToken(l.current, token.RBRACE)
	case '^':
		l.addToken(l.current, token.CARET)
	case ':':
		if l.match('=') {
			l.addStringToken(":=", token.ASSIGN)
		}
	case '.':
		if isDigit(l.peek()) {
			l.number()
		} else {
			l.addToken(l.current, token.DOT)
		}
	case '\'':
		l.string()
	case '\n':
		l.line++
		break
	case '\t':
	case '\r':
	case ' ':
		break
	default:
		if isDigit(l.current) {
			l.number()
		} else if isAlpha(l.current) {
			l.identifier()
		} else {
			l.addToken(l.current, token.ILLEGAL)
		}

	}

}

/*
A number can be:
20
20.00
.05
*/

func (l *Lexer) number() {
	start := l.position - 1

	if l.current != '.' {
		for isDigit(l.current) {
			l.advance()
		}
	}

	if l.current == '.' {
		var decimal = false
		l.advance()
		for isDigit(l.current) {
			l.advance()
			decimal = true
		}
		if !decimal {

		}
	}

	l.addStringToken(string(l.source[start:l.position-1]), token.NUMBER)

}

func (l *Lexer) identifier() {

	start := l.position - 1

	for isAlpha(l.current) || isDigit(l.current) {
		l.advance()
	}
	lit := string(l.source[start : l.position-1])
	tt := token.LookupKeyWord(lit)
	l.addStringToken(lit, tt)
}

func isDigit(ch rune) bool {
	return ch >= '0' && ch <= '9'
}

func isAlpha(ch rune) bool {
	return (ch >= 'a' && ch <= 'z') || (ch >= 'A' && ch <= 'Z') || ch == '_'
}

func (l *Lexer) string() {
	start := l.position
	l.advance()
	for l.current != '\'' {
		l.advance()
	}
	l.addStringToken(string(l.source[start:l.position-1]), token.STRING)
}

func (l *Lexer) isAtEnd() bool {
	if l.position >= len(l.source) {
		return true
	}
	return false
}
