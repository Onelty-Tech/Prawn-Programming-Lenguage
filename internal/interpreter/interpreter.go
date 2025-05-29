package main

import (
	"fmt"
	"prawn/lexer"
	"prawn/parser"
)

var InterLexer = lexer.InitLexer(`var myVar = "Hola Mundo";
var myInt = 200;
write(myInt);
write(myVar);
write("Hola Wey");`)

// le pasamos el campo de los tokenchan del lexer
var ParserGuard = parser.NewParser(InterLexer.TokenChan)

var AST, errors = ParserGuard.Parse()
var Program = AST["Program"]

func indentTypes(node interface{}) {
	//verificamos de que tipo de dato es el interface
	switch eval := node.(type) {
	case parser.StringExpr:
		fmt.Print(eval.Value)
	case parser.NumberExpr:
		fmt.Println(eval.Value)
	case parser.VarExpr:
		for _, node := range Program {
			if varIdent, exists := node["VarDeclare"]; exists {
				if varInf, exists := varIdent.(map[string]interface{}); exists {
					if varInf["Ident"] == eval.Name {
						indentTypes(varInf["Value"])
					}
				}
			}
		}
	case parser.BinaryExpr:
		//em breve pronto voce podra criar esse :n
		fmt.Println(eval)
	default:
		//
	}
}

func main() {
	for _, node := range Program {
		if varDecl, exists := node["VarDeclare"]; exists {
			payload := varDecl.(map[string]interface{})["Payload"]
			fmt.Println("Declaracion de variable: ", varDecl)
			indentTypes(payload)
		} else if writeDecl, exists := node["Write"]; exists {
			payload := writeDecl.(map[string]interface{})
			indentTypes(payload["Value"])
		}
	}
}
