package node

import "minimal/lexer/token"


type BinaryExpression struct {
  BaseNode
}

func NewBinary(left INode, operation token.Token, right INode) INode {
  return BinaryExpression {
    BaseNode: BaseNode {
      kind: Binary,
      left: left,
      operation: operation,
      right: right,
    },
  }
}
