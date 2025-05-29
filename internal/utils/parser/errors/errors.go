package errors

import "fmt"

// errores para caracteres esperados pero enves de eso pusieron otros JSSJSJJSJSJSJ
func CreateErrorExpected(line int, column int, expected string, found string) string {
	return fmt.Sprintf("Expected '%s' but found '%s' in %d:%d", expected, found, line, column)
}

func CreateErrorUnrecognizableToken(line int, column int, tokenLiteral string) string {
	return fmt.Sprintf("Token Unrecognizable '%s' in %d:%d", tokenLiteral, line, column)
}
