package lexer

type TokenType int

const (
	TokenTypeEOF TokenType = iota
	TokenTypeIllegal

	TokenTypePrefix
	TokenTypeLabel
	TokenTypeText
	TokenTypeFlagName
)
