package token

import (
	"errors"
	"regexp"
)

var wordRegex = regexp.MustCompile(`\([^)]+\)|\w+(?:'\w+)*|[.,!?;:'"]`)

func Tokenize(input string) ([]string, error) {
	matches := wordRegex.FindAllString(input, -1)
	if matches == nil {
		return nil, errors.New("no valid tokens found")
	}
	return matches, nil
}

//
//func Tokenize(input string) ([]string, error) {
//	var tokens []string
//	var current strings.Builder
//
//	runes := []rune(input)
//	for i := 0; i < len(runes); {
//		r := runes[i]
//
//		switch {
//		case r == '(':
//			j := i
//			for j < len(runes) && runes[j] != ')' {
//				j++
//			}
//			if j < len(runes) && runes[j] == ')' {
//				tokens = append(tokens, string(runes[i:j+1]))
//				i = j + 1
//			}
//			if j >= len(runes) {
//				return nil, errors.New("unclosed command: missing ')'")
//			} else {
//				current.WriteRune(r)
//				i++
//			}
//		case isLetterOrDigit(r):
//			current.Reset()
//			for i < len(runes) && isLetterOrDigit(runes[i]) {
//				current.WriteRune(runes[i])
//				i++
//			}
//			tokens = append(tokens, current.String())
//		case isPunctuation(r):
//			tokens = append(tokens, string(r))
//			i++
//		case isSpace(r):
//			i++
//		default:
//			i++
//		}
//	}
//
//	return tokens, nil
//}
//
//func isLetterOrDigit(r rune) bool {
//	return (r >= 'a' && r <= 'z') ||
//		(r >= 'A' && r <= 'Z') ||
//		(r >= '0' && r <= '9')
//}
//
//func isPunctuation(r rune) bool {
//	return strings.ContainsRune(".,!?;:'\"", r)
//}
//
//func isSpace(r rune) bool {
//	return r == ' ' || r == '\t' || r == '\n'
//}
