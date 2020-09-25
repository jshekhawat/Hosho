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
}

//LookupKeyWord ...
func LookupKeyWord(lex string) Type {
	if tok, ok := keywords[lex]; ok {
		return tok
	}
	return IDENT
}
