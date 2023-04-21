package node

import "minimal/types/token"


type UnaryExpression struct {
  kind NodeKind

  position,
  length int

  operation token.Token
  expression INode
}

func NewUnary(operation token.Token, value INode) INode {
  // If the operation of postfix so the value comes before the operator.
  var position int
  var length int

  if operation.GetPosition() < value.GetPosition() {
    position = operation.GetPosition()
    length = (value.GetPosition() + value.GetLength()) - position
  } else /* if value.GetPosition() < operation.GetPosition() */ {
    position = value.GetPosition()
    length = (operation.GetPosition() + operation.GetLength()) - position
  }

  return UnaryExpression {
    kind: Unary,

    position: position,
    length: length,

    operation: operation,
    expression: value,
  }
}


func (self UnaryExpression) GetKind() NodeKind { return self.kind }

func (self UnaryExpression) GetPosition() int { return self.position }
func (self UnaryExpression) GetLength() int { return self.length }

func (self UnaryExpression) GetOperation() token.Token { return self.operation }
func (self UnaryExpression) GetExpression() INode { return self.expression }
