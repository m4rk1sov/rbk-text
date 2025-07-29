package token

import (
	"errors"
	"regexp"
	"strings"
)

var tokenRegex = regexp.MustCompile(`\((?:up|low|cap|hex|bin)(?:,\s*\d+)?\)|'[^']*'|"[^"]*"|[\\&]|'|"|\w+(?:'\w+)*|[.,!?;:]+|[()]|\\n+`)

func Tokenize(clean string) ([]string, error) {
	matches := tokenRegex.FindAllString(clean, -1)
	if matches == nil {
		return nil, errors.New("no valid tokens found")
	}

	var filtered []string
	for _, token := range matches {
		if strings.TrimSpace(token) != "" {
			filtered = append(filtered, token)
		}
	}

	return filtered, nil
}
