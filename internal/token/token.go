package token

import "strings"

func Tokenize(input string) ([]string, error) {
	var tokens []string
	var current strings.Builder
	
	runes := []rune(input)
	for i := 0; i < len(runes); {
		r := runes[i]
		
		switch {
		case r == '(':
			end := strings.IndexRune(input[i:], ')')
			if end != -1 {
				tokens = append(tokens, input[i:i+end+1])
				i += end + 1
			} else {
				// not a command, consume the rune
				current.WriteRune(r)
				i++
			}
		case isLetterOrDigit(r):
			current.Reset()
			for i < len(runes) && isLetterOrDigit(runes[i]) {
				current.WriteRune(runes[i])
				i++
			}
			tokens = append(tokens, current.String())
		case isPunctuation(r):
			tokens = append(tokens, string(r))
			i++
		case isSpace(r):
			i++
		default:
			i++
		}
	}
	
	return tokens, nil
}

func isLetterOrDigit(r rune) bool {
	return (r >= 'a' && r <= 'z') ||
		(r >= 'A' && r <= 'Z') ||
		(r >= '0' && r <= '9')
}

func isPunctuation(r rune) bool {
	return strings.ContainsRune(".,!?;:\"", r)
}

func isSpace(r rune) bool {
	return r == ' ' || r == '\t' || r == '\n'
}
