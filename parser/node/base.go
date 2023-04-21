package node

import "minimal/lexer/token"


type BaseNode struct {
  kind NodeKind

  position int
  length int

  /* For literals */
  token token.Token

  /* For unary expressions */
  operation token.Token
  value INode

  /* For binary expressions */
  left INode
  right INode
}


func (self BaseNode) GetKind() NodeKind { return self.kind }

func (self BaseNode) GetPosition() int { return self.position }
func (self BaseNode) GetLength() int { return self.length }

/* For literals */
func (self BaseNode) GetToken() token.Token { return self.token }

/* For unary expressions */
func (self BaseNode) GetOperation() token.Token { return self.operation }
func (self BaseNode) GetValue() INode { return self.value }

/* For binary expressions */
func (self BaseNode) GetLeft() INode { return self.left }
func (self BaseNode) GetRight() INode { return self.right }
