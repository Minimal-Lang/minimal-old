package node

import (
  "fmt"
  "minimal/lexer/token"
)


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

func (self BaseNode) Print() {
  var loop func(node INode, indent string)
  loop = func(node INode, last_indent string) {
    indent := last_indent

    switch node.GetKind() {
      case Literal:
        node.GetToken().Print()

      case Unary:
        fmt.Printf("Unary: %d - %d {\n", node.GetPosition(), node.GetLength())
        indent += "  "

        fmt.Print(indent)
        node.GetOperation().Print()

        fmt.Print(indent)
        loop(node.GetRight(), indent)

        fmt.Printf("%s}\n", last_indent)

      case Binary:
        fmt.Printf("Binary: %d - %d {\n", node.GetPosition(), node.GetLength())
        indent += "  "

        fmt.Print(indent)
        loop(node.GetLeft(), indent)

        fmt.Print(indent)
        node.GetOperation().Print()

        fmt.Print(indent)
        loop(node.GetRight(), indent)

        fmt.Printf("%s}\n", last_indent)
    }
  }

  loop(self, "")
}
