package lexer

import (
	"prawn/lexer/tokenspec"
	"prawn/utils/lexer/review"
)

/*
	structura del Lexer 1. Input completo, 2.Posicion actual,

3. Proxima Posicion a Leer,
4. El caracter actual,
*/
type Lexer struct {
	input       string
	position    int
	nextRead    int
	currentChar byte
	currentLine int
	TokenChan   chan tokenspec.Token
}

// constructor
func InitLexer(input string) *Lexer {
	Lexer := &Lexer{input: input, TokenChan: make(chan tokenspec.Token)}
	//empieza a leer el primer caracter
	Lexer.readChar()
	go Lexer.tokenize()
	//retorna la 'struct'
	return Lexer
}

/*
readChar funcion fundamental para avanzar con la lectura del codigo,
lee byte por byte(Letra por letra)
*/
func (lexer *Lexer) readChar() {
	// si llego al final de input resetea el currentChar a '0'
	if lexer.nextRead >= len(lexer.input) {
		lexer.currentChar = 0
	} else {
		// de lo contrario sigue
		lexer.currentChar = byte(lexer.input[lexer.nextRead])
	}
	/*actualizamos los datos
	lexer.position(la posicion actual) lo actualizamos al lexer.nextRead
	que es el siguiente para Leer
	*/
	lexer.position = lexer.nextRead
	lexer.nextRead++

	if lexer.currentChar == '\n' {
		lexer.currentLine++
	}
}

// da un vistazo del proximo caracter sin mover el puntero del Lexer
func (lexer *Lexer) previewNextChar() rune {
	// retorna tipo de dato 'rune' solo si aun no termina el input
	if lexer.nextRead >= len(lexer.input) {
		return 0
	}
	return rune(lexer.input[lexer.nextRead])
}

func (lexer *Lexer) Tokenizer() tokenspec.Token {
	//crea una variable token para cada token
	var token tokenspec.Token

	lexer.jumpWhitespaces()

	if lexer.currentChar == 0 {
		return tokenspec.Token{
			Type:     tokenspec.EOF,
			Literal:  "",
			Position: lexer.position,
			Line:     lexer.currentLine,
		}
	}

	//salta los espacios en blanco

	switch {
	case lexer.currentChar == '"':
		token = lexer.readStringToken()
	case review.IsLetter(lexer.currentChar):
		generatedToken := lexer.readIdentifierToken()
		token = generatedToken
	case review.IsDigit(lexer.currentChar):
		generatedToken := lexer.readNumberToken()
		token = generatedToken
	case review.IsSymbol(lexer.currentChar):
		generatedToken := lexer.readSymbolToken()
		token = generatedToken
	default:
		token = tokenspec.NewToken(tokenspec.ILLEGAL, string(lexer.currentChar), lexer.position, lexer.currentLine)
		lexer.readChar()
	}
	return token
}

// salta todos los espacios
func (lexer *Lexer) jumpWhitespaces() {
	for lexer.currentChar == ' ' || lexer.currentChar == '\n' || lexer.currentChar == '\t' {
		// incrementa solo si encuentra salto de linea
		lexer.readChar()
	}
}

// parte una porcion del codigo
func (lexer *Lexer) readIdentifierToken() tokenspec.Token {
	start := lexer.position
	for review.IsLetter(lexer.currentChar) || review.IsDigit(lexer.currentChar) {
		lexer.readChar()
	}
	literal := lexer.input[start:lexer.position]
	tokenType := tokenspec.LookupIdent(literal)
	return tokenspec.NewToken(tokenType, literal, start, lexer.currentLine)
}

func (lexer *Lexer) readNumberToken() tokenspec.Token {
	start := lexer.position
	for review.IsDigit(lexer.currentChar) {
		lexer.readChar()
	}
	literal := lexer.input[start:lexer.position]
	return tokenspec.NewToken(tokenspec.INT, literal, start, lexer.currentLine)
}

func (lexer *Lexer) readSymbolToken() tokenspec.Token {
	//tomamos la posicion inicial
	start := lexer.position
	//tomamos el currentChar inicial
	ch := string(lexer.currentChar)
	//creamos una variable para el proximo Char
	next := string(lexer.previewNextChar())

	// Intenta encontrar un token de dos caracteres
	if tokType, ok := tokenspec.SymbolTokens[ch+next]; ok {
		lexer.readChar() // avanzar uno
		lexer.readChar() // avanzar otro
		return tokenspec.NewToken(tokType, ch+next, start, lexer.currentLine)
	}

	// Intenta encontrar token de un solo carácter
	if tokType, ok := tokenspec.SymbolTokens[ch]; ok {
		lexer.readChar()
		return tokenspec.NewToken(tokType, ch, start, lexer.currentLine)
	}

	// Si no es un símbolo válido, retorna ILLEGAL
	lexer.readChar()
	return tokenspec.NewToken(tokenspec.ILLEGAL, ch, start, lexer.currentLine)
}

// esta funcion lee lo que hay dentro de las comillas y lo retorna como string ya que no es un tokentype
func (lexer *Lexer) readStringToken() tokenspec.Token {
	start := lexer.position // guarda la posición de la comilla inicial
	lexer.readChar()        // salta la comilla de apertura

	strStart := lexer.position
	for lexer.currentChar != '"' && lexer.currentChar != 0 {
		lexer.readChar()
	}
	literal := lexer.input[strStart:lexer.position]

	token := tokenspec.Token{
		Type:     tokenspec.STRING,
		Literal:  literal,
		Position: start,
		Line:     lexer.currentLine,
	}

	lexer.readChar() // salta la comilla final

	return token
}

func (lexer *Lexer) tokenize() {
	for tok := lexer.Tokenizer(); tok.Type != tokenspec.EOF; tok = lexer.Tokenizer() {
		lexer.TokenChan <- tok
	}
	close(lexer.TokenChan)
}

/*
func main() {
	fmt.Println(`el codigo enviado fue este:
	var nombre = 235;
	write(nombre)`)
	fmt.Println("----TOKENS CREADOS----")
	lexer := InitLexer(`
	var myBool = 250 != 900
	if(myBool){
		write("Es true")
	};
	write("Hola Mundo XDDDDSODSODSD SD-SOS-OS");`)
	for tok := lexer.Tokenizer(); tok.Type != tokenspec.EOF; tok = lexer.Tokenizer() {
		fmt.Printf("Type: '%-7s' Literal: '%s' Position: '%d' Line: '%d'\n", tok.Type, tok.Literal, tok.Position, tok.Line)
	}
	fmt.Scanln()
}
*/
