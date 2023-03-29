package text


type TextSpan struct {
  start int
  length int
  end int
}

func NewTextSpan(start, length int) TextSpan {
  return TextSpan {
    start,
    length,
    start + length,
  }
}


func (self TextSpan) GetStart() int { return self.start }
func (self TextSpan) GetLength() int { return self.length }
func (self TextSpan) GetEnd() int { return self.end }
