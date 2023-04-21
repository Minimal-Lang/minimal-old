package node

import "minimal/lexer/token"


type INode interface {
  GetKind() NodeKind

  GetPosition() int
  GetLength() int

  /* For literals */
  GetToken() token.Token

  /* For unary expressions */
  GetOperation() token.Token
  GetValue() INode

  /* For binary expressions */
  GetLeft() INode
  // GetOperation() token.Token
  GetRight() INode
}
