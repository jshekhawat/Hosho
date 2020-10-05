package token

//Comments
const (
	ILLEGAL  = "ILLEGAL"
	EQUAL    = "="
	GT       = ">"
	GTE      = ">="
	LT       = "<"
	LTE      = "<="
	BANG     = "!"
	PLUS     = "+"
	MINUS    = "-"
	SLASH    = "/"
	STAR     = "*"
	LPAREN   = "("
	RPAREN   = ")"
	LBRACKET = "["
	RBRACKET = "]"
	LBRACE   = "{"
	RBRACE   = "}"
	COLON    = ":"
	CARET    = "^"
	ASSIGN   = ":="
	DOT      = "."
	STRING   = "'"
	NUMBER   = "NUMBER"
	COMMENT  = "#"
	NOT      = "NOT"
	EOF      = "EOF"
	AND      = "AND"
	OR       = "OR"
	IN       = "IN"
	IDENT    = "IDENT"
	IF       = "IF"
	ELSE     = "ELSE"
	FUNCTION = "FUNCTION"
	LET      = "LET"
	TRUE     = "TRUE"
	FALSE    = "FALSE"
	LIKE     = "LIKE"
)

//Type is
type Type string

//Token is
type Token struct {
	Lexeme string
	Type   Type
}

var keywords = map[string]Type{
	"func":  FUNCTION,
	"let":   LET,
	"true":  TRUE,
	"false": FALSE,
	"if":    IF,
	"else":  ELSE,
	"in":    IN,
	"and":   AND,
	"or":    OR,
	"like":  LIKE,
}

//LookupKeyWord ...
func LookupKeyWord(lex string) Type {
	if tok, ok := keywords[lex]; ok {
		return tok
	}
	return IDENT
}
