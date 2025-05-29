package tokenspec

type TokenType string

type Token struct {
	Type     TokenType
	Literal  string
	Position int
	Line     int
}

const (
	ILLEGAL          = "ILLEGAL"          // ILLEGAL VALUE
	EOF              = "EOF"              //EOF NUMBER
	IDENT            = "IDENT"            // IDENT VARIABLE NAME,FUNCTION NAME ETC
	INT              = "INT"              // INT example: (var edad = 25;)
	STRING           = "STRING"           // STRING example: (var nombre = "Pedro")
	LPAREN           = "LPAREN"           //LPAREN "("
	RPAREN           = "RPAREN"           //RPAREN ")"
	ASSIGN           = "ASSIGN"           //ASSIGN "="
	PLUS             = "PLUS"             // PLUS "+"
	MINUS            = "MINUS"            // MINUS "-"
	ASTERISK         = "ASTERISK"         //ASTERISK "*"
	SLASH            = "SLASH"            //SLASH "/"
	MOD              = "MOD"              //MOD "%"
	SEMICOLON        = "SEMICOLON"        //SEMICOLON ";"
	EQUAL            = "EQUAL"            //EQUAL "=="
	NOT_EQUAL        = "NOT_EQUAL"        //NOT_EQUAL "!="
	LESS             = "LESS"             // LESS "<"
	GREATER          = "GREATER"          //GREATER ">"
	LESS_OR_EQUAL    = "LESS_OR_EQUAL"    // LESS_EQUAL "<="
	GREATER_OR_EQUAL = "GREATER_OR_EQUAL" //GREATER_EQUAL ">="
	BANG             = "BANG"             //BANG "!"

	//keywords
	VAR   = "VAR"
	WRITE = "WRITE"
)

// Lo usa el lexer nadamas
var SymbolTokens = map[string]TokenType{
	//Simbolos unitarios
	"+": PLUS,
	"-": MINUS,
	"*": ASTERISK,
	"/": SLASH,
	"%": MOD,
	"=": ASSIGN,
	"(": LPAREN,
	")": RPAREN,
	";": SEMICOLON,
	"!": BANG,
	"<": LESS,
	">": GREATER,

	// Multi caracters
	"==": EQUAL,
	"!=": NOT_EQUAL,
	"<=": LESS_OR_EQUAL,
	">=": GREATER_OR_EQUAL,
}

// Lo usa el parser nadamas
var SymbolsArithmetic = map[string]string{
	"+": PLUS,
	"-": MINUS,
	"*": ASTERISK,
	"/": SLASH,
	"%": MOD,
}

// Lo usan el parser y lexer
var Keywords = map[string]TokenType{
	"write": WRITE,
	"var":   VAR,
}

// funcion para crear el token que usa la 'struct' TokenType y 'c' que es el currentChar
func NewToken(tokenType TokenType, literal string, position int, line int) Token {
	//retorna el token con el tipo y el literal
	return Token{Type: tokenType, Literal: literal, Position: position, Line: line}
}

func LookupIdent(ident string) TokenType {
	// verifica si el IDENT que le pasamos existe en keywords
	/*Que ase?
	1. Crea 2 variables una del token y otro de la confirmacion
	2. entra a el mapa 'keywords' a la clave ident
	3. guarda en token la key ident como TokenType
	4. Retorna el tokentype
	*/
	if token, exists := Keywords[ident]; exists {
		return token
	}
	// Si no existe retorna ident
	return IDENT
}
