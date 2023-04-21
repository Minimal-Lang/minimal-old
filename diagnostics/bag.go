package diagnostics

import (
  "fmt"
  "minimal/types/token"
  . "minimal/text"
)


type DiagnosticsBag struct {
  diagnostics []Diagnostic
}

func NewDiagnosticBag() DiagnosticsBag {
  return DiagnosticsBag {}
}


func (self DiagnosticsBag) GetDiagnostics() []Diagnostic { return self.diagnostics }


func (self *DiagnosticsBag) report(span TextSpan, message string) {
  diagnostic := NewDiagnostic(span, message)
  self.diagnostics = append(self.diagnostics, diagnostic)
}

func (self *DiagnosticsBag) ReportIllegalChar(position int, char rune) {
  span := NewTextSpan(position, 1)
  message := fmt.Sprintf("Illegal character input: '%c'", char)

  self.report(span, message)
}

func (self *DiagnosticsBag) ReportUnexpectedToken(span TextSpan, kind, expected token.TokenKind) {
  message := fmt.Sprintf("Unexpected token %s, expected %s", kind, expected)
  self.report(span, message)
}
