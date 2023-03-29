package utils


// func IsLetter(char rune) bool {
//   return char > 'a' && char < 'z' || char > 'A' && char < 'Z'
// }

func IsNumber(char rune) bool {
  return char > '0' && char < '9'
}
