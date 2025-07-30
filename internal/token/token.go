package token

import (
	"errors"
	"regexp"
	"strings"
)

var tokenRegex = regexp.MustCompile(`\((?:up|low|cap|hex|bin)(?:,\s*\d+)?\)|'[^']*'|"[^"]*"|[\\&]|'|"|\w+(?:'\w+)*|[.,!?;:]+|[()]|\\n+`)

var quotedContentRegex = regexp.MustCompile(`\w+(?:'\w+)*|[.,!?;:]+|\\n+|\S`)

//var tokenRegex = regexp.MustCompile(`\((?:up|low|cap|hex|bin)(?:,\s*\d+)?\)|[\\&]|'|"|\w+(?:'\w+)*|[.,!?;:]+|[()]|\\n+`)

//func Tokenize(clean string) ([]string, error) {
//	matches := tokenRegex.FindAllString(clean, -1)
//	if matches == nil {
//		return nil, errors.New("no valid tokens found")
//	}
//
//	var filtered []string
//	for _, token := range matches {
//		if strings.TrimSpace(token) != "" {
//			filtered = append(filtered, token)
//		}
//	}
//	return filtered, nil
//}

func Tokenize(clean string) ([]string, error) {
	matches := tokenRegex.FindAllString(clean, -1)
	if matches == nil {
		return nil, errors.New("no valid tokens found")
	}
	
	var filtered []string
	
	for _, token := range matches {
		if strings.TrimSpace(token) == "" {
			continue
		}
		if isQuotedContent(token) {
			splitTokens := splitQuotedContent(token)
			filtered = append(filtered, splitTokens...)
		} else {
			filtered = append(filtered, token)
		}
	}
	
	return filtered, nil
}

func isQuotedContent(token string) bool {
	if strings.HasPrefix(token, "'") && strings.HasSuffix(token, "'") && len(token) > 2 {
		return true
	}
	if strings.HasPrefix(token, `"`) && strings.HasSuffix(token, `"`) && len(token) > 2 {
		return true
	}
	return false
}

func splitQuotedContent(quotedToken string) []string {
	var result []string
	
	// quote type and contents
	var quoteChar string
	var content string
	
	if strings.HasPrefix(quotedToken, "'") && strings.HasSuffix(quotedToken, "'") {
		quoteChar = "'"
		content = quotedToken[1 : len(quotedToken)-1]
	} else if strings.HasPrefix(quotedToken, `"`) && strings.HasSuffix(quotedToken, `"`) {
		quoteChar = `"`
		content = quotedToken[1 : len(quotedToken)-1]
	} else {
		return []string{quotedToken}
	}
	
	// opening quote
	result = append(result, quoteChar)
	
	if strings.TrimSpace(content) != "" {
		contentTokens := quotedContentRegex.FindAllString(content, -1)
		for _, contentToken := range contentTokens {
			trimmed := strings.TrimSpace(contentToken)
			if trimmed != "" {
				result = append(result, trimmed)
			}
		}
	}
	
	result = append(result, quoteChar)
	
	return result
}
