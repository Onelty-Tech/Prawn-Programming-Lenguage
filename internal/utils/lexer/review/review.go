package review

import "prawn/lexer/tokenspec"

func IsLetter(currentChar byte) bool {
	return (currentChar >= 'a' && currentChar <= 'z') ||
		(currentChar >= 'A' && currentChar <= 'Z') ||
		currentChar == '_'
}

func IsDigit(currentChar byte) bool {
	return currentChar >= '0' && currentChar <= '9'
}

func IsSymbol(currentChar byte) bool {
	if _, confirm := tokenspec.SymbolTokens[string(currentChar)]; confirm {
		return true
	}
	return false
}

func IsArithmeticSymbol(currentChar string) bool {
	if _, confirm := tokenspec.SymbolsArithmetic[currentChar]; confirm {
		return true
	}
	return false
}
