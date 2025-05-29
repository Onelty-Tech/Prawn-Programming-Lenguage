package parser

import (
	"prawn/lexer/tokenspec"
)

type WriteDecl struct {
	Value Node
}

func (wrt *WriteDecl) Payload() map[string]interface{} {
	return map[string]interface{}{
		"Value": wrt.Value,
	}
}

// Parsea write("Hola Mundo")
func (parser *Parser) ParseWriteDecl() *WriteDecl {
	/*current token 'write' asi que asemos un parser.NextToken
	y pasamos al siguiente token
	*/
	parser.NextToken()
	//el current token '(' y verificamos si existe si no agregamos el error y retornamos nil
	if parser.currentToken().Type != tokenspec.LPAREN {
		parser.errors = append(parser.errors, "Expected '(' but found '%s'.", parser.currentToken().Literal)
		return nil
	}
	//pasa al contenido
	parser.NextToken()
	//aqui guarda el contenido en un tipo de dato 'Node'
	writeContentValue := parser.ParseExpressionType()
	//pasa al siguiente token que deberia de ser ')'
	parser.NextToken()
	if parser.currentToken().Type != tokenspec.RPAREN {
		parser.errors = append(parser.errors, "Expected ')' but found '%s'.", parser.currentToken().Literal)
		return nil
	}
	parser.NextToken()
	if parser.currentToken().Type != tokenspec.SEMICOLON {
		parser.errors = append(parser.errors, "Expected ';' but found '%s'.", parser.currentToken().Literal)
		return nil
	}
	return &WriteDecl{
		Value: writeContentValue,
	}
}
