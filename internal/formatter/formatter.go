package formatter

import (
	"strings"
	"unicode/utf8"
)

func isMultiCharPunctuation(token string) bool {
	return token == "..." || token == "?!" || token == "!?"
}

func isPunctuation(token string) bool {
	return strings.ContainsAny(token, ".,!?:;")
}

func isQuote(token string) bool {
	return token == "'"
}

func isVowel(token string) bool {
	if token == "" {
		return false
	}
	r, _ := utf8.DecodeRuneInString(strings.ToLower(token))
	switch r {
	case 'a', 'e', 'i', 'o', 'u', 'h':
		return true
	default:
		return false
	}
}

func JoinTokens(tokens []string) string {
	var b strings.Builder
	i := 0
	
	for i < len(tokens) {
		token := tokens[i]
		
		// checking for an article
		if strings.ToLower(token) == "a" && i+1 < len(tokens) {
			next := tokens[i+1]
			if isVowel(next) {
				b.WriteString("an")
				i++
				continue
			}
		}
		
		if i == 0 {
			b.WriteString(token)
			i++
			continue
		}
		
		prev := tokens[i-1]
		
		// quotes
		if isQuote(token) {
			if i > 0 && !isPunctuation(prev) {
				b.WriteString(token)
			} else {
				b.WriteString("'")
				b.WriteString(token)
			}
			i++
			continue
		}
		
		if isMultiCharPunctuation(token) {
			b.WriteString(token)
			i++
			continue
		}
		
		if isPunctuation(token) {
			b.WriteString(token)
			i++
			continue
		}
		if isPunctuation(prev) || isQuote(prev) {
			b.WriteString(" ")
			b.WriteString(token)
		} else {
			b.WriteString(" ")
			b.WriteString(token)
		}
		i++
	}
	
	return b.String()
}
