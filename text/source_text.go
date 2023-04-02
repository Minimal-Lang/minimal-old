package text


type SourceText struct {
  Lines []TextLine
  text []rune
}

func NewSourceText(text string) SourceText {
  self := SourceText {

  }

  self.text = []rune(text)
  self.Lines = self.parseLines()

  return self
}


func (self *SourceText) GetLineIndex(position int) int {
  lower := 0
  upper := len(self.Lines) - 1

  for lower <= upper {
    index := lower + (upper - lower) / 2
    start := self.Lines[index].GetStart()

    if position == start {
      return index
    }

    if start > position {
      upper = index - 1
    } else {
      lower = index + 1
    }
  }

  return lower - 1
}

func (self *SourceText) parseLines() []TextLine {
  var res []TextLine

  position := 0
  lineStart := 0

  for position < len(self.text) {
    lineBreakLength := self.getLineBreakLength(position)

    if lineBreakLength == 0 {
      position++
    } else {
      self.addLine(&res, position, lineStart, lineBreakLength)

      position += lineBreakLength
      lineStart = position
    }
  }

  if position >= lineStart {
    self.addLine(&res, position, lineStart, 0)
  }

  return res
}

func (self SourceText) addLine(res *[]TextLine, position, lineStart, lineBreakLength int) {
  length := position - lineStart
  lineLengthIncludingLineBreak := length + lineBreakLength
  line := NewTextLine(self, lineStart, length, lineLengthIncludingLineBreak)
  *res = append(*res, line)
}

func (self *SourceText) getLineBreakLength(position int) int {
  char := self.text[position]
  var lookahead rune

  if position + 1 >= len(self.text) {
    lookahead = '\x00'
  } else {
    lookahead = self.text[position + 1]
  }

  if char == '\r' && lookahead == '\n' {
    return 2
  }

  if char == '\r' || char == '\n' {
    return 1
  }

  return 0
}

func (self SourceText) String() string {
  return string(self.text)
}

func (self SourceText) StringFromSpan(span TextSpan) string {
  return self.StringFromBounds(span.GetStart(), span.GetLength())
}

func (self SourceText) StringFromBounds(start, length int) string {
  return string(self.text[start:start + length])
}
