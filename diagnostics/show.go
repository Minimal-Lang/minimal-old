package diagnostics

import (
	"fmt"

	. "minimal/text"

  . "github.com/pandasoli/colorstrings"
)


func (self Diagnostic) PrintDiagnostic(source_text SourceText) {
  fmt.Printf("\033[47;30m ERROR \033[0m: %s\n", self.message)
  fmt.Println()

  error_line_index := source_text.GetLineIndex(self.GetSpan().GetStart())
  error_line := source_text.Lines[error_line_index]
  error_x := self.GetSpan().GetStart() - error_line.GetStart()

  print_line := func(index uint, line string) {
    colored_line, _ := NewColorString(line)
    line_number_color := "\033[90m"

    if index == uint(error_line_index) {
      colored_line.Colorize("4:3;58;2;191;97;106", uint(error_x))
      colored_line.Colorize("0", uint(error_x + 1))
      line_number_color = "\033[37m"
    }

    fmt.Printf("  %s%d\033[90m▕\033[0m %s\n", line_number_color, index + 1, line)
  }

  if error_line_index > 0 {
    line_index := error_line_index - 1
    line := source_text.Lines[line_index]

    print_line(uint(line_index), line.String())
  }

  print_line(uint(error_line_index), error_line.String())

  if error_line_index < len(source_text.Lines) {
    line_index := error_line_index + 1
    line := source_text.Lines[line_index]

    print_line(uint(line_index), line.String())
  }

  fmt.Println()
}
