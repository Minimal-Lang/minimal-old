package parser

import (
	. "minimal/diagnostics"
	"minimal/types/token"
	"minimal/types/node"
)


type Parser struct {
  tokens []token.Token
  position int
  diagnostics DiagnosticsBag
}

func NewParser(tokens []token.Token) *Parser {
  return &Parser {
    tokens: tokens,
    diagnostics: NewDiagnosticBag(),
  }
}


func (self Parser) GetDiagnostics() []Diagnostic { return self.diagnostics.GetDiagnostics() }


func (self Parser) peek(offset int) token.Token {
  i := self.position + offset

  if i >= len(self.tokens) {
    return self.tokens[len(self.tokens) - 1]
  }

  return self.tokens[i]
}

func (self Parser) get_current() token.Token { return self.peek(0) }

func (self *Parser) next() token.Token {
  current := self.get_current()
  self.position++
  return current
}

func (self *Parser) match(kind token.TokenKind) token.Token {
  if self.get_current().GetKind() == kind {
    return self.next()
  }

  current := self.get_current()
  self.diagnostics.ReportUnexpectedToken(current.GetSpan(), current.GetKind(), kind)
  return token.New(kind, "", current.GetPosition())
}

func (self *Parser) Parse() node.INode { return self.binary(0) }

func (self *Parser) binary(parent_precedence uint8) node.INode {
  unary_op_prece := self.get_current().GetUnaryOperatorPrecetence()

  var left node.INode

  if unary_op_prece != 0 && unary_op_prece >= parent_precedence {
    op := self.next()
    operand := self.binary(unary_op_prece)

    left = node.NewUnary(op, operand)
  } else {
    left = self.primary()
  }

  for {
    prece := self.get_current().GetBinaryOperatorPrecedence()
    if prece == 0 || prece <= parent_precedence {
      break
    }

    op := self.next()
    right := self.binary(prece)

    left = node.NewBinary(left, op, right)
  }

  return left
}

func (self *Parser) primary() node.INode {
  if self.get_current().GetKind() == token.OpenParentheses {
    self.next()
    expr := self.binary(0)
    self.match(token.CloseParentheses)

    return expr
  }

  return node.NewLiteral(self.match(token.Number))
}
