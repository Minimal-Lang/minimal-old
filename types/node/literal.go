package node

import "minimal/types/token"


type LiteralExpression struct {
  kind NodeKind

  position,
  length int

  expression token.Token
}

func NewLiteral(expression token.Token) INode {
  return LiteralExpression {
    kind: Literal,

    position: expression.GetPosition(),
    length: expression.GetLength(),

    expression: expression,
  }
}


func (self LiteralExpression) GetKind() NodeKind { return self.kind }

func (self LiteralExpression) GetPosition() int { return self.position }
func (self LiteralExpression) GetLength() int { return self.length }

func (self LiteralExpression) GetExpression() token.Token { return self.expression }
