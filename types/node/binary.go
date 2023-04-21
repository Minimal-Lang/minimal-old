package node

import "minimal/types/token"


type BinaryExpression struct {
  kind NodeKind

  position,
  length int

  left INode
  operation token.Token
  right INode
}

func NewBinary(left INode, operation token.Token, right INode) INode {
  return BinaryExpression {
    kind: Binary,

    position: left.GetPosition(),
    length: (right.GetPosition() + right.GetLength()) - left.GetPosition(),

    left: left,
    operation: operation,
    right: right,
  }
}


func (self BinaryExpression) GetKind() NodeKind { return self.kind }

func (self BinaryExpression) GetPosition() int { return self.position }
func (self BinaryExpression) GetLength() int { return self.length }

func (self BinaryExpression) GetLeft() INode { return self.left }
func (self BinaryExpression) GetOperation() token.Token { return self.operation }
func (self BinaryExpression) GetRight() INode { return self.right }
