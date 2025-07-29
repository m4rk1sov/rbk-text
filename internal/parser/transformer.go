package parser

import (
	"strconv"
	"strings"
	"unicode"
)

func TransformTokens(tokens []string) ([]string, error) {
	var result []string
	var wordIndexes []int
	
	for i := 0; i < len(tokens)-1; i++ {
		currentLower := strings.ToLower(tokens[i])
		
		if currentLower != "a" && currentLower != "an" {
			continue
		}
		
		// lookahead for word
		j := i + 1
		for j < len(tokens) && !isWord(tokens[j]) {
			j++
		}
		
		if j >= len(tokens) {
			continue
		}
		
		nextWord := tokens[j]
		needsAn := isVowel(nextWord)
		if needsAn && currentLower == "a" {
			tokens[i] = "an"
		} else if !needsAn && currentLower == "an" {
			tokens[i] = "a"
		}
		
	}
	
	for i := 0; i < len(tokens); i++ {
		token := tokens[i]
		
		if cmd, ok := ParseCommand(token); ok {
			count := cmd.Count
			start := len(wordIndexes) - count
			if start < 0 {
				start = 0
			}
			
			// applying the transformation to the last counted words
			for _, wordIndex := range wordIndexes[start:] {
				word := result[wordIndex]
				transformed, err := applyTransform(word, cmd.Type)
				if err != nil {
					// change of design, we simply leave the word if error occurs
					continue
				}
				result[wordIndex] = transformed
			}
			//clearing the applied words by the command
			wordIndexes = wordIndexes[:start]
		} else {
			result = append(result, token)
			// index of a word that can be transformed
			if isWord(token) {
				wordIndexes = append(wordIndexes, len(result)-1)
			}
		}
	}
	
	return result, nil
}

func applyTransform(word, typ string) (string, error) {
	quote := ""
	if (strings.HasPrefix(word, "'") && strings.HasSuffix(word, "'")) ||
		(strings.HasPrefix(word, `"`) && strings.HasSuffix(word, `"`)) {
		quote = string(word[0])
		word = word[1 : len(word)-1]
	}
	
	var transformed string
	var err error
	
	switch typ {
	case "up":
		transformed = strings.ToUpper(word)
	case "low":
		transformed = strings.ToLower(word)
	case "cap":
		transformed = toTitle(word)
	case "hex":
		transformed, err = fromHex(word)
	case "bin":
		transformed, err = fromBin(word)
	default:
		transformed = word
	}
	
	if quote != "" {
		transformed = quote + transformed + quote
	}
	
	return transformed, err
}

func toTitle(word string) string {
	if len(word) == 0 {
		return ""
	}
	// simpler transformation by rune
	runes := []rune(word)
	runes[0] = unicode.ToUpper(runes[0])
	return string(runes)
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

// handling only words, not the punctuation and quotes
func isWord(token string) bool {
	if token == "" {
		return false
	}
	
	// command checking
	_, isCommand := ParseCommand(token)
	if isCommand {
		return false
	}
	
	// unquote if quoted
	if (strings.HasPrefix(token, "'") && strings.HasSuffix(token, "'")) ||
		(strings.HasPrefix(token, `"`) && strings.HasSuffix(token, `"`)) {
		token = token[1 : len(token)-1]
	}
	
	// only punctuations
	chars := `.,!?;:"'()\`
	if strings.Trim(token, chars) == "" {
		return false
	}
	
	return true
}

func isVowel(token string) bool {
	if token == "" {
		return false
	}
	
	if strings.HasPrefix(token, "'") && strings.HasSuffix(token, "'") && len(token) > 2 {
		inner := token[1 : len(token)-1]
		return startsWithVowel(inner)
	}
	
	if strings.HasPrefix(token, `"`) && strings.HasSuffix(token, `"`) && len(token) > 2 {
		inner := token[1 : len(token)-1]
		return startsWithVowel(inner)
	}
	
	if isWord(token) {
		return startsWithVowel(token)
	}
	
	return false
}

func startsWithVowel(word string) bool {
	if word == "" {
		return false
	}
	
	lowerWord := strings.ToLower(word)
	firstChar := rune(lowerWord[0])
	
	vowels := "aeiouh"
	if strings.ContainsRune(vowels, firstChar) {
		return true
	}
	
	return false
}
