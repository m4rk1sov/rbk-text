package parser

import (
	"strconv"
	"strings"
	"unicode"
)

func TransformTokens(tokens []string) ([]string, error) {
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
				transformed, err := applyTransform(word, cmd.Type)
				if err != nil {
					return nil, err
				}
				result[idx] = transformed
			}
		} else {
			result = append(result, token)
			if isWord(token) {
				wordIndexes = append(wordIndexes, len(result)-1)
			}
		}
	}

	return result, nil
}

func applyTransform(word, typ string) (string, error) {
	switch typ {
	case "up":
		return strings.ToUpper(word), nil
	case "low":
		return strings.ToLower(word), nil
	case "cap":
		return toTitle(word), nil
	case "hex":
		return fromHex(word)
	case "bin":
		return fromBin(word)
	default:
		return word, nil
	}
}

func toTitle(word string) string {
	if len(word) == 0 {
		return word
	}
	return strings.ToUpper(string(word[0])) + strings.ToLower(word[1:])
}

func fromHex(word string) (string, error) {
	n, err := strconv.ParseInt(word, 16, 64)
	if err != nil {
		return word, err
	}
	return strconv.FormatInt(n, 10), nil
}

func fromBin(word string) (string, error) {
	n, err := strconv.ParseInt(word, 2, 64)
	if err != nil {
		return word, err
	}
	return strconv.FormatInt(n, 10), nil
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
