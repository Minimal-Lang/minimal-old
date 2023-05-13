package diagnostics

import (
	"fmt"

	. "minimal/text"
	"minimal/types/token"

	. "github.com/pandasoli/colorstrings"
)


func ColorCode(source string, tokens []token.Token) (result ColorString) {
  result, _ = NewColorString(source)

  for _, current := range tokens {
    var color string

    switch current.GetKind() {
      case token.Plus, token.Minus,
           token.Asterisk, token.Slash,
           token.OpenParentheses, token.CloseParentheses:
        color = "2"

      case token.Illegal:
        color = "1;4:3"
    }

    result.Colorize(color, uint(current.GetPosition()))
    result.Colorize("0", uint(current.GetPosition() + current.GetLength()))
  }

  return result
}

func (self Diagnostic) PrintDiagnostic(source_text SourceText, tokens []token.Token) {
  fmt.Printf("\033[47;30m ERROR \033[0m: %s\n", self.message)
  fmt.Println()

  error_line_index := source_text.GetLineIndex(self.GetSpan().GetStart())
  error_line := source_text.Lines[error_line_index]

  colored_code := ColorCode(source_text.String(), tokens)
  colored_lines, _ := colored_code.SplitString("\n")

  print_line := func(index uint, line string) {
    colored_line := colored_lines[index]
    line_number_color := "\033[90m"

    if index == uint(error_line_index) {
      line_number_color = "\033[37m"
    }

    fmt.Printf("  %s%d\033[90m▕\033[0m %s\n", line_number_color, index + 1, colored_line.Join())
  }

  if error_line_index > 0 {
    line_index := error_line_index - 1
    line := source_text.Lines[line_index]

    print_line(uint(line_index), line.String())
  }

  print_line(uint(error_line_index), error_line.String())

  if error_line_index < len(source_text.Lines) - 1 {
    line_index := error_line_index + 1
    line := source_text.Lines[line_index]

    print_line(uint(line_index), line.String())
  }

  fmt.Println()
}
