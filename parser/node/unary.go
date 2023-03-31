package node

import "minimal/lexer/token"


type UnaryExpression struct {
  BaseNode
}

func NewUnary(operation token.Token, value INode) INode {
  return UnaryExpression {
    BaseNode: BaseNode {
      kind: Unary,
      operation: operation,
      value: value,
    },
  }
}
