package main

import (
	"fmt"
	"prawn/lexer"
	"prawn/parser"
)

func main() {
	lexer := lexer.InitLexer(`var test = 200 + 200;
	var myxd = 400;
	write("Hola Mundo XDDDDSODSODSD-A2-2");`)

	// Crear el parser y pasarle el canal de tokens
	parser := parser.NewParser(lexer.TokenChan)

	// Parsear los tokens que vayan llegando

	/*no sabe el tipo de dato
	 */
	AST, errors := parser.Parse()
	//aqui imprime el AST que retorna el Parser
	// esque como es un channel, para imprimir un chan es asis
	for k, v := range AST {
		fmt.Println("main: ", k, v)
	}
	fmt.Println("Errors: ", errors)

	program := AST["Program"]

	//ejemplo de como acceder a los campos (By ChatGPT)
	for i, node := range program {
		if varDecl, ok := node["VarDeclare"]; ok {
			fmt.Printf("VarDeclare #%d:\n", i+1)

			// Aquí accedes al payload
			payload := varDecl.(map[string]map[string]interface{})["Payload"]

			name := payload["Ident"].(string)
			value := payload["Value"]

			fmt.Println("  Name:", name)
			fmt.Println("  Value:", value)
		} else if writeDecl, ok := node["Write"]; ok {
			fmt.Printf("WriteDeclare#%d:\n", i+1)
			payload := writeDecl.(map[string]map[string]interface{})["Payload"]
			value := payload["Value"]
			fmt.Printf("Value: %s\n", value)
		}
	}
}
