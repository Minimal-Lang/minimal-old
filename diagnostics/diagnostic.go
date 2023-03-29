package diagnostics

import . "minimal/text"


type Diagnostic struct {
  span TextSpan
  message string
}

func NewDiagnostic(span TextSpan, message string) Diagnostic {
  return Diagnostic { span, message }
}


func (self Diagnostic) GetSpan() TextSpan { return self.span }
func (self Diagnostic) GetMessage() string { return self.message }
