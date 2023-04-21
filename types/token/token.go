package token

import (
  "fmt"
  . "minimal/text"
)


type Token struct {
  kind TokenKind
  literal string

  position int
  length int

  span TextSpan
}

func New(kind TokenKind, literal string, position int) Token {
  length := len(literal)

  if literal == "\x00" { length = 0 }

  return Token {
    kind,
    literal,

    position,
    length,

    NewTextSpan(position, length),
  }
}


func (self Token) GetKind() TokenKind { return self.kind }
func (self Token) GetLiteral() string { return self.literal }

func (self Token) GetPosition() int { return self.position }
func (self Token) GetLength() int { return self.length }

func (self Token) GetSpan() TextSpan { return self.span }

func (self Token) Print() {
  fmt.Printf(
    "Token:%s { val: \"%s\", pos: %d, len: %d }\n",
    self.kind,
    self.literal,
    self.position,
    self.length,
  );
}

func (self Token) GetUnaryOperatorPrecetence() uint8 {
  switch self.kind {
    case Plus, Minus: return 3

    default: return 0
  }
}

func (self Token) GetBinaryOperatorPrecedence() uint8 {
  switch self.kind {
    case Asterisk, Slash: return 2
    case Plus, Minus: return 1

    default: return 0
  }
}
