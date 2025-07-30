package parser

import (
	"strings"
)

func isNewlineSequence(token string) bool {
	return strings.HasPrefix(token, "\\n")
}

func JoinTokens(tokens []string) string {
	if len(tokens) == 0 {
		return ""
	}

	var b strings.Builder
	inSingleQuote := false
	inDoubleQuote := false

	// Write the first token directly.
	b.WriteString(tokens[0])
	if tokens[0] == "'" {
		inSingleQuote = true
	}
	if tokens[0] == `"` {
		inDoubleQuote = true
	}

	// Loop starting from the second token.
	for i := 1; i < len(tokens); i++ {
		token := tokens[i]
		prevToken := tokens[i-1]

		// --- Logic to decide if a space is needed BEFORE the current token ---
		addSpace := true

		// Rule 1: No space before punctuation or a closing bracket.
		if IsPunctuation(token) || token == ")" {
			addSpace = false
		}

		// Rule 2: No space after an opening bracket.
		if prevToken == "(" {
			addSpace = false
		}

		// Rule 3: The quote/apostrophe logic using state.
		if token == "'" {
			if inSingleQuote {
				addSpace = false
			} else {
				if prevToken == "(" || IsPunctuation(prevToken) {
					addSpace = false
				} else {
					addSpace = true
				}
			}
			inSingleQuote = !inSingleQuote
		}

		if prevToken == "'" {
			if inSingleQuote {
				addSpace = false
			} else {
				if IsPunctuation(token) || token == ")" {
					addSpace = false
				} else {
					addSpace = true
				}
			}
		}

		// Rule 4: Same logic for double quotes.
		if token == `"` {
			if inDoubleQuote {
				addSpace = false
			} else {
				if prevToken == "(" || IsPunctuation(prevToken) {
					addSpace = false
				} else {
					addSpace = true
				}
			}
			inDoubleQuote = !inDoubleQuote
		}

		if prevToken == `"` {
			if inDoubleQuote {
				addSpace = false
			} else {
				if IsPunctuation(token) || token == ")" {
					addSpace = false
				} else {
					addSpace = true
				}
			}
		}

		// Rule 5: Checking for \ (in case of \n)
		if prevToken == `\` {
			addSpace = false
		}

		// Rule 6: No space in newline sequences
		if isNewlineSequence(token) || isNewlineSequence(prevToken) {
			addSpace = false
		}

		// Decision and write
		if addSpace {
			b.WriteString(" ")
		}
		b.WriteString(token)

	}

	return b.String()
}
