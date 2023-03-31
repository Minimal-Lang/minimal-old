package node

import "minimal/lexer/token"


type LiteralExpression struct {
  BaseNode
}

func NewLiteral(token token.Token) INode {
  return LiteralExpression {
    BaseNode: BaseNode {
      kind: Literal,
      token: token,
    },
  }
}
