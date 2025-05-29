package parser

import (
	"fmt"
	"prawn/lexer/tokenspec"
	"prawn/utils/lexer/review"
	"prawn/utils/parser/errors"
)

// contiene la declaracion de una variable tipo(Name "miVariable" Value: NumberExpr)
type VarDeclare[T any] struct {
	Name  string
	Value T
}

func (varDecl *VarDeclare[T]) Payload() map[string]interface{} {
	return map[string]interface{}{
		"Ident": varDecl.Name,
		"Value": varDecl.Value,
	}
}

func (parser *Parser) ParseVarDeclare() *VarDeclare[any] {
	//pasa al siguiente token 'IDENT' (nombre de la variable)
	parser.NextToken()
	if parser.currentToken().Type != tokenspec.IDENT {
		parser.errors = append(parser.errors, errors.CreateErrorExpected(parser.position, 2, tokenspec.IDENT, string(parser.currentToken().Type)))
		return nil
	}
	//aguarda el nombre de la variable
	varName := &VarExpr{Name: parser.currentToken().Literal}
	//pasa al siguiente token 'ASSIGN'(=)
	parser.NextToken()
	// si no encuentra el token tira error y lo almacena en un slice de errores
	if parser.currentToken().Type != tokenspec.ASSIGN {
		//hay que mejorar este mensaje de error
		parser.errors = append(parser.errors, errors.CreateErrorExpected(parser.position, 2, tokenspec.ASSIGN, string(parser.currentToken().Type)))
		fmt.Println(parser.errors)
		//no retorna nada
		return nil
	}
	/*si no se encontro ningun error sigue y lo proximo deberia ser
	el contenido
	*/
	parser.NextToken()
	if review.IsArithmeticSymbol(parser.previewNextToken().Literal) {
		varValueLeft := parser.ParseExpressionType()
		varValue := parser.CreateBinaryExpression(varValueLeft)
		if parser.currentToken().Type != tokenspec.SEMICOLON {
			parser.errors = append(parser.errors, fmt.Sprintf("Expected ';' but found '%s'", parser.currentToken().Literal))
			return nil
		}
		/*avanza al siguiente token para no dejar al currentToken con el mismo si no se pone esto
		podria causar error*/
		parser.NextToken()
		return &VarDeclare[any]{
			Name:  varName.Name,
			Value: *varValue,
		}
	}
	varValue := parser.ParseExpressionType()
	parser.NextToken()
	if parser.currentToken().Type != tokenspec.SEMICOLON {
		parser.errors = append(parser.errors, fmt.Sprintf("Expected ';' but found '%s'", parser.currentToken().Literal))
		return nil
	}
	parser.NextToken()
	return &VarDeclare[any]{
		Name:  varName.Name,
		Value: varValue,
	}
}
