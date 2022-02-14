package lexer

import "Monkey/token"

// isLetter
// 引数で渡された文字が、英数字であるかチェック
func isLetter(ch byte) bool {
	return 'a' <= ch && ch <= 'z' || 'A' <= ch && 'Z' <= ch || ch == '_'
}

// isDigit
// 引数で渡された文字が、数字であるかチェック
func isDigit(ch byte) bool {
	return '0' <= ch && ch <= '9'
}

// newToken
// 引数で渡された、tokenTypeとchを構造体に入れて返す
func newToken(tokenType token.TokenType, ch byte) token.Token {
	return token.Token{Type: tokenType, Literal: string(ch)}
}
