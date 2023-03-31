package parser

import (
	"fmt"
	"testing"

	. "minimal/diagnostics"
	. "minimal/lexer"
	"minimal/lexer/token"
	"minimal/parser/node"

	"github.com/stretchr/testify/assert"
)


type TokenTest struct {
  Content string
  Result node.INode
  DiagnosticsCount int
}


func TestTokens(t *testing.T) {
  assert := assert.New(t)

  tests := []TokenTest {
    {
      "123 * -(12 / -2)+6",
      node.NewBinary(
        node.NewBinary(
          node.NewLiteral(token.New(token.Number, "123", 0)),
          token.New(token.Asterisk, "*", 3),
          node.NewUnary(
            token.New(token.Minus, "-", 5),
            node.NewBinary(
              node.NewLiteral(token.New(token.Number, "12", 8)),
              token.New(token.Slash, "/", 10),
              node.NewUnary(
                token.New(token.Minus, "-", 12),
                node.NewLiteral(token.New(token.Number, "2", 14)),
              ),
            ),
          ),
        ),
        token.New(token.Plus, "+", 15),
        node.NewLiteral(token.New(token.Number, "6", 17)),
      ),
      0,
    },
  }

  for _, test := range tests {
    tokens, _ := lexall(test.Content)

    par := NewParser(tokens)
    ast := par.Parse()

    assert.Equal(test.DiagnosticsCount, len(par.GetDiagnostics()))
    test_node(assert, ast, test.Result)
  }
}

func test_node(assert *assert.Assertions, current, expected node.INode) {
  switch current.GetKind() {
    case node.Literal:
      val := current.GetToken()
      exp_val := expected.GetToken()

      assert.Equal(exp_val.GetKind(), val.GetKind())
      assert.Equal(exp_val.GetLiteral(), val.GetLiteral())
      assert.Equal(expected.GetPosition(), current.GetPosition())
      assert.Equal(expected.GetLength(), current.GetLength())

    case node.Unary:
      expr := current.GetValue()
      op := current.GetOperation()

      exp_expr := expected.GetLeft()
      exp_op := expected.GetOperation()

      test_node(assert, expr, exp_expr)
      assert.Equal(exp_op.GetKind(), op.GetKind())

    case node.Binary:
      left := current.GetLeft()
      op := current.GetOperation()
      right := current.GetRight()

      exp_left := expected.GetLeft()
      exp_op := expected.GetOperation()
      exp_right := expected.GetRight()

      test_node(assert, left, exp_left)
      test_node(assert, right, exp_right)
      assert.Equal(exp_op.GetKind(), op.GetKind())
  }
}

func lexall(content string) (tokens []token.Token, diagnostics []Diagnostic) {
  lex := NewLexer(content)
  current := token.New(token.Number, "12", 0)

  for current.GetKind() != token.EOF {
    current = lex.Lex()
    tokens = append(tokens, current)
  }

  return tokens, lex.GetDiagnostics()
}
