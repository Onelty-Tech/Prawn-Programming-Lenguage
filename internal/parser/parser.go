package parser

import (
	"fmt"
	"prawn/lexer/tokenspec"
	"strconv"
)

type Parser struct {
	tokens   []tokenspec.Token
	position int
	errors   []string
}

// Node contiene todo tipo de datos
type Node interface{}

// contiene una expresion tipo(x + 12)
type BinaryExpr struct {
	Left  Node
	Op    string
	Right Node
}

// contiene numeros
type NumberExpr struct {
	Value int
}

type StringExpr struct {
	Value string
}

// contiene el nombre de variables
type VarExpr struct {
	Name string
}

// el constructor de la 'struct'
/*
El lexer le pasa al parser la informacion por un channel ?Como Funciona un Channel{
	al leer un indice se borra
	no puedes leer el proximo indice sin borrarlo
}
	asi que por eso decidi tener una variable para almacenarlo,primero lee
	todo el contenido del channel para poder usarlo como pegue la gana
*/
func NewParser(tokenchan chan tokenspec.Token) *Parser {
	var buffer []tokenspec.Token
	//va agregando token por token al buffer
	for tok := range tokenchan {
		//Si es EOF termina(osea el fin del codigo)
		if tok.Type == tokenspec.EOF {
			break
		}
		buffer = append(buffer, tok)
	}
	//retorna la 'struct'
	return &Parser{
		tokens:   buffer,
		position: 0,
		errors:   []string{},
	}
}

// verifica el proximo token sin saltarlo
func (parser *Parser) previewNextToken() tokenspec.Token {
	return parser.tokens[parser.position+1]
}

// regresa el token actual sin modificar el parser.position🗣️
func (parser *Parser) currentToken() tokenspec.Token {
	return parser.tokens[parser.position]
}

// version simplificada sin manejo de errores para donde no se necesita del manejo de errores
func Atoi(input string) int {
	value, _ := strconv.Atoi(input)
	return value
}

/*
Esta funcion Lee el tipo de Dato INT,STRING,IDENT y lo retorna en Un 'struct' Node
*/
//error esto regresa un Nil en un binaryexpr, no detecta que sea un int
func (parser *Parser) ParseExpressionType() Node {
	//crea un switch para validar el currentToken.Type
	switch parser.currentToken().Type {
	//si es INT crea el Nodo y lo pasa a Int por que esta en String(Necesita manejo de errores(ahorita nadamas por test))
	case tokenspec.INT:
		valueInt := Atoi(parser.currentToken().Literal)
		return NumberExpr{Value: valueInt}
	case tokenspec.STRING:
		valueStr := parser.currentToken().Literal
		return StringExpr{Value: valueStr}
	case tokenspec.IDENT:
		valueIDENT := parser.currentToken().Literal
		return VarExpr{Name: valueIDENT}
	// otros casos como booleanos, operaciones, etc.
	default:
		return nil
	}

}

// tipo constructor que crea 'BinaryExpr',left node,Operator string,right node
func (parser *Parser) CreateBinaryExpression(leftValue Node) *BinaryExpr {
	//Guarda el primer valor(lo guarda como literal)
	//avanza al siguiente token que seria el operador tipo (*,-,+)
	parser.NextToken()
	OpValue := parser.currentToken()
	parser.NextToken()
	rightValue := parser.ParseExpressionType()
	parser.NextToken()
	return &BinaryExpr{
		Left:  leftValue,
		Op:    OpValue.Literal,
		Right: rightValue,
	}
}

func (parser *Parser) NextToken() tokenspec.Token {
	if parser.position < len(parser.tokens) {
		token := parser.tokens[parser.position]
		parser.position++
		return token
	}
	return tokenspec.Token{
		Type: tokenspec.EOF,
	}
}

// parsea Nodo por nodo
func (parser *Parser) parseNode() map[string]interface{} {
	//leer que tipo de token es
	//aqui vamos a leer los tipos de tokens
	currentToken := parser.currentToken()
	switch currentToken.Type {
	case tokenspec.VAR:
		node := parser.ParseVarDeclare()
		return map[string]interface{}{
			"VarDeclare": node.Payload(),
		}
	case tokenspec.WRITE:
		node := parser.ParseWriteDecl()
		return map[string]interface{}{
			"Write": node.Payload(),
		}
	default:
		parser.errors = append(parser.errors, fmt.Sprintf("Token '%s' no reconocido position '%d'", parser.currentToken().Literal, parser.position))
		return nil
	}
}

// Crea el AST(Abstract Sintaxys Tree)
func (parser *Parser) Parse() (map[string][]map[string]interface{}, []string) {
	var program []map[string]interface{}

	for parser.position < len(parser.tokens) {
		node := parser.parseNode()
		if node != nil {
			program = append(program, node)
		} else {
			// avanzar para no quedar en bucle infinito si hay error
			parser.NextToken()
		}
	}
	return map[string][]map[string]interface{}{
		"Program": program,
	}, parser.errors
}
