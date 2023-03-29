package token


type TokenKind string

const (
  EOF = "EOF"
  Illegal = "Illegal"

  Number = "Number"

  Plus = "Plus"
  Minus = "Minus"
  Asterisk = "Asterisk"
  Slash = "Slash"

  OpenParentheses = "OpenParentheses"
  CloseParentheses = "CloseParentheses"
)
