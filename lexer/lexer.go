package lexer

import (
	"unicode"
  "strings"

  "minimal/lexer/token"
  . "minimal/utils"
	. "minimal/diagnostics"
)


type Lexer struct {
  source []rune
  current rune

  position int
  next_position int

  diagnostics DiagnosticsBag
}

func NewLexer(content string) *Lexer {
  return &Lexer {
    source: []rune(content),
    diagnostics: NewDiagnosticBag(),
  }
}


func (self Lexer) GetDiagnostics() []Diagnostic { return self.diagnostics.GetDiagnostics() }


func (self Lexer) peek() rune {
  if self.next_position >= len(self.source) {
    return '\x00'
  }

  return self.source[self.next_position]
}

func (self *Lexer) next() rune {
  if self.next_position >= len(self.source) {
    self.current = '\x00'
  } else {
    self.current = self.source[self.next_position]
  }

  self.position = self.next_position
  self.next_position++

  return self.current
}

func (self *Lexer) read(buffer *strings.Builder, test func(rune) bool) {
  for test(self.peek()) {
    buffer.WriteRune(self.next())
  }
}

func (self *Lexer) skip_whitespace() {
  for unicode.IsSpace(self.peek()) {
    self.next()
  }
}

func (self *Lexer) Lex() token.Token {
  self.skip_whitespace()
  self.next()

  var buffer strings.Builder
  var kind token.TokenKind = token.Illegal
  position := self.position

  buffer.WriteRune(self.current)

  switch self.current {
    case '\x00': kind = token.EOF

    case '+': kind = token.Plus
    case '-': kind = token.Minus
    case '*': kind = token.Asterisk
    case '/': kind = token.Slash

    case '(': kind = token.OpenParentheses
    case ')': kind = token.CloseParentheses

    default:
      if IsNumber(self.current) {
        self.read(&buffer, IsNumber)
        kind = token.Number
      } else {
        self.diagnostics.ReportIllegalChar(position, self.current)
      }
  }

  return token.New(kind, buffer.String(), position)
}
