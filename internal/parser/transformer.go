package parser

import (
	"strconv"
	"strings"
	"unicode"
)

func TransformTokens(tokens []string) []string {
	var result []string
	var wordIndexes []int

	for i := 0; i < len(tokens); i++ {
		token := tokens[i]

		if cmd, ok := ParseCommand(token); ok {
			count := cmd.Count
			start := len(wordIndexes) - count
			if start < 0 {
				start = 0
			}

			//
			for _, idx := range wordIndexes[start:] {
				word := result[idx]
				result[idx] = applyTransform(word, cmd.Type)
			}
		} else {
			result = append(result, token)
			if isWord(token) {
				wordIndexes = append(wordIndexes, len(result)-1)
			}
		}
	}

	return result
}

func applyTransform(word, typ string) string {
	switch typ {
	case "up":
		return strings.ToUpper(word)
	case "down":
		return strings.ToLower(word)
	case "cap":
		return toTitle(word)
	case "hex":
		return fromHex(word)
	case "bin":
		return fromBin(word)
	default:
		return word
	}
}

func toTitle(word string) string {
	if len(word) == 0 {
		return word
	}
	return strings.ToUpper(string(word[0])) + strings.ToLower(word[1:])
}

func fromHex(word string) string {
	n, err := strconv.ParseInt(word, 16, 64)
	if err != nil {
		return word
	}
	return strconv.FormatInt(n, 10)
}

func fromBin(word string) string {
	n, err := strconv.ParseInt(word, 2, 64)
	if err != nil {
		return word
	}
	return strconv.FormatInt(n, 10)
}

func isWord(s string) bool {
	if len(s) == 0 {
		return false
	}
	for _, r := range s {
		if !unicode.IsLetter(r) && !unicode.IsDigit(r) {
			return false
		}
	}
	return true
}
