package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"

	. "minimal/lexer"
	. "minimal/parser"
	. "minimal/text"
	"minimal/types/node"
	"minimal/types/token"
)


func main() {
  fmt.Printf("\033[1;47;30m  𝙼𝚒𝚗𝚒𝚖𝚊𝚕 - 𝙰𝚗 𝚘𝚕𝚍 𝚗𝚎𝚠 𝚙𝚛𝚘𝚐𝚛𝚊𝚖𝚖𝚒𝚗𝚐 𝚕𝚊𝚗𝚐𝚞𝚊𝚐𝚎 :𝙳  \033[0m\n")
  fmt.Println()

  render := bufio.NewReader(os.Stdin)

  var text_builder strings.Builder
  show_tokens := false
  show_ast := false

  for {
    // Here I can check if the `text_builder` is empty or not,
    // If it is, it means this is the first time the user is inputing something.
    // If it is not empty it means the user inputed some thing incomplete.
    if text_builder.Len() == 0 {
      fmt.Print("\033[1m›\033[0m ")
    } else {
      fmt.Print(" \033[1m›\033[0m ")
    }

    input, _ := render.ReadString('\n')

    input = strings.TrimSpace(input)
    is_blank := len(input) == 0

    if text_builder.Len() == 0 {
      if is_blank {
        break
      } else if input == "#show tokens" {
        show_tokens = !show_tokens

        var color string
        if show_tokens { color = "\033[32m" } else { color = "\033[33m" }

        var state string
        if show_tokens { state = "Showing" } else { state = "Not showing" }

        fmt.Printf("  🦴 %s%s lex tokens.\033[0m\n", color, state)

        continue
      } else if input == "#show tree" {
        show_ast = !show_ast

        var color string
        if show_ast { color = "\x1b[32m" } else { color = "\x1b[33m" }

        var state string
        if show_ast { state = "Showing" } else { state = "Not showing" }

        fmt.Printf("  ☠️ %s%s parse trees.\033[0m\n", color, state)

        continue
      }
    }

    text_builder.WriteString(input)
    text := text_builder.String()
    source_text := NewSourceText(text)

    lex := NewLexer(text)
    current := token.New(token.Number, "12", 0)
    var tokens []token.Token

    for current.GetKind() != token.EOF {
      current = lex.Lex()
      tokens = append(tokens, current)
    }

    par := NewParser(tokens)
    ast := par.Parse()

    if !is_blank && len(par.GetDiagnostics()) > 0 {
      continue
    }

    diagnostics := append(lex.GetDiagnostics(), par.GetDiagnostics()...)

    if show_tokens {
      for _, token := range tokens {
        token.Print()
      }

      fmt.Println()
    }

    if show_ast {
      printAST(ast, "", true)
      fmt.Println()
    }

    if len(diagnostics) > 0 {
      for _, diag := range diagnostics {
        lineIndex := source_text.GetLineIndex(diag.GetSpan().GetStart())
        lineNumber := lineIndex + 1
        line := source_text.Lines[lineIndex]
        char_position := diag.GetSpan().GetStart() - line.GetStart() + 1

        fmt.Println()
        fmt.Printf(" \033[31m( line: %d, position: %d ): %s\033[0m\n", lineNumber, char_position, diag.GetMessage())

        prefixSpan := NewTextSpan_FromBounds(line.GetStart(), diag.GetSpan().GetStart())
        suffixSpan := NewTextSpan_FromBounds(diag.GetSpan().GetEnd(), line.GetEnd())

        prefix := source_text.StringFromSpan(prefixSpan)
        suffix := source_text.StringFromSpan(suffixSpan)
        err := source_text.StringFromSpan(diag.GetSpan())

        fmt.Printf("  ╰─ %s\033[31m%s\033[0m%s\n", prefix, err, suffix)
        fmt.Println()
      }
    }

    text_builder.Reset()
  }
}

func printAST(item interface {}, indent string, isLast bool) {
  var marker string
  if isLast { marker = "╰─ " } else { marker = "├─ " }

  fmt.Print("\033[90m" + indent + marker)

  var kind string
  if tok, isTk := item.(token.Token); isTk {
    kind = string(tok.GetKind())
  } else if node, isNode := item.(node.INode); isNode {
    kind = string(node.GetKind())
  }

  fmt.Print("\033[34m" + kind)
  fmt.Print("\033[0m")

  if tok, isTk := item.(token.Token); isTk {
    fmt.Print(": ")
    fmt.Print("\033[33m" + tok.GetLiteral())
    fmt.Print("\033[0m")
  }

  fmt.Println()

  if item, isNode := item.(node.INode); isNode {
    if isLast {
      indent += "   "
    } else {
      indent += "│  "
    }

    var children []interface {}

    if item, isLiteral := item.(node.LiteralExpression); isLiteral {
      children = append(children, item.GetExpression())
    }
    if item, isUnary := item.(node.UnaryExpression); isUnary {
      children = append(children, item.GetOperation())
      children = append(children, item.GetExpression())
    }
    if item, isBinary := item.(node.BinaryExpression); isBinary {
      children = append(children, item.GetLeft())
      children = append(children, item.GetOperation())
      children = append(children, item.GetRight())
    }

    for i, child := range children {
      printAST(child, indent, i == len(children) - 1)
    }
  }
}
