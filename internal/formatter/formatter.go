package formatter

import (
	"github.com/m4rk1sov/rbk-text/internal/parser"
	"strings"
)

func isWord(token string) bool {
	if token == "" {
		return false
	}

	// command checking
	_, isCommand := parser.ParseCommand(token)
	if isCommand {
		return false
	}

	// newline sequence
	if strings.HasPrefix(token, "\\n") {
		return false
	}

	// only punctuations
	chars := `.,!?;:"'()\`
	if strings.Trim(token, chars) == "" {
		return false
	}

	return true
}

func isPunctuation(token string) bool {
	if token == "" {
		return false
	}
	// A token is punctuation if it ONLY contains punctuation characters.
	return strings.Trim(token, ".,!?;:") == ""
}

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
		if isPunctuation(token) || token == ")" {
			addSpace = false
		}

		// Rule 2: No space after an opening bracket.
		if prevToken == "(" {
			addSpace = false
		}

		// Rule 3: The quote/apostrophe logic using state.
		if token == "'" {
			// If the previous token is a word, this MUST be a contraction or a closing quote. No space needed.
			// Examples: "don'", "Casey's", "word'"
			if isWord(prevToken) {
				addSpace = false
			}
		}

		if prevToken == "'" {
			// If the previous token was a single quote, we need to decide if we just entered a quote.
			if inSingleQuote {
				// We just entered a quote (e.g., 'word). No space.
				addSpace = false
			}
		}

		// Rule 4: Same logic for double quotes.
		if token == `"` && isWord(prevToken) {
			addSpace = false
		}
		if prevToken == `"` && inDoubleQuote {
			addSpace = false
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

		// Update state *after* processing the current token.
		if token == "'" {
			inSingleQuote = !inSingleQuote
		}
		if token == `"` {
			inDoubleQuote = !inDoubleQuote
		}
	}

	return b.String()
}
