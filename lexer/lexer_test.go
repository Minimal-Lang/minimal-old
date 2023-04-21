package lexer

import (
	"testing"

	. "minimal/diagnostics"
	"minimal/types/token"
	"github.com/stretchr/testify/assert"
)


type TokenTest struct {
  Content string
  Result []token.Token
  DiagnosticsCount int
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

func TestTokens(t *testing.T) {
  assert := assert.New(t)

  tests := []TokenTest {
    {
      "123 * -(12 / -2)+6",
      []token.Token {
        token.New(token.Number, "123", 0),
        token.New(token.Asterisk, "*", 4),
        token.New(token.Minus, "-", 6),
        token.New(token.OpenParentheses, "(", 7),
        token.New(token.Number, "12", 8),
        token.New(token.Slash, "/", 11),
        token.New(token.Minus, "-", 13),
        token.New(token.Number, "2", 14),
        token.New(token.CloseParentheses, ")", 15),
        token.New(token.Plus, "+", 16),
        token.New(token.Number, "6", 17),
        token.New(token.EOF, "\x00", 18),
      },
      0,
    },
  }

  for _, test := range tests {
    tokens, diagnostics := lexall(test.Content)

    assert.Equal(len(diagnostics), test.DiagnosticsCount)
    assert.Equal(len(tokens), len(test.Result))

    for i := range make([]int, len(tokens)) {
      res := tokens[i]
      expected := test.Result[i]

      assert.Equal(expected.GetKind(), res.GetKind())
      assert.Equal(expected.GetLiteral(), res.GetLiteral())
      assert.Equal(expected.GetPosition(), res.GetPosition())
      assert.Equal(expected.GetLength(), res.GetLength())
    }
  }
}
