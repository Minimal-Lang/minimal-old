package text


type TextLine struct {
  text SourceText

  start,
  length,
  lengthIncludingLineBreak,
  end int

  span,
  spanIncludingLineBreak TextSpan
}

func NewTextLine(text SourceText, start, length, lengthIncludingLineBreak int) TextLine {
  return TextLine {
    text,

    start,
    length,
    lengthIncludingLineBreak,
    start + length,

    NewTextSpan(start, length),
    NewTextSpan(start, lengthIncludingLineBreak),
  }
}


func (self TextLine) GetText() SourceText { return self.text }

func (self TextLine) GetStart() int { return self.start }
func (self TextLine) GetLength() int { return self.length }
func (self TextLine) GetLengthIncludingLineBreak() int { return self.lengthIncludingLineBreak }
func (self TextLine) GetEnd() int { return self.end }

func (self TextLine) GetSpan() TextSpan { return self.span }
func (self TextLine) GetSpanIncludingLineBreak() TextSpan { return self.spanIncludingLineBreak }


func (self TextLine) String() string {
  return self.text.StringFromSpan(self.span)
}
