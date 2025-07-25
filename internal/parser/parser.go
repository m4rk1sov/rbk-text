package parser

import (
	"strconv"
	"strings"
)

type Command struct {
	Type  string // "cap", "up", "bin" and others
	Count int    // 1, 2 or none (will default to 1) (the number of previous tokens to apply)
}

func ParseCommand(token string) (*Command, bool) {
	token = strings.TrimSpace(token)

	runes := []rune(token)

	length1 := len(runes)
	if length1 < 3 || token[0] != '(' || token[length1-1] != ')' {
		return nil, false
	}

	// contents of the command
	body := token[1 : length1-1]
	parts := strings.Split(body, ",")

	// checking if there is more than 2 parts (e.g. (cap, 2, fasdf)
	length2 := len(parts)
	if length2 > 2 {
		return nil, false
	}

	// default count is 1
	cmd := Command{Count: 1}

	// removing the whitespaces in type
	cmd.Type = strings.TrimSpace(parts[0])
	cmd.Type = strings.ToLower(cmd.Type)
	if cmd.Type == "" {
		return nil, false
	}

	if len(parts) == 2 {
		rawCount := strings.TrimSpace(parts[1])
		if rawCount == "" {
			return nil, false
		}

		n, err := strconv.Atoi(rawCount)
		if err != nil || n <= 0 {
			return nil, false
		}
		cmd.Count = n
	}

	switch cmd.Type {
	case "cap", "low", "up", "bin", "hex":
		return &cmd, true
	default:
		return nil, false
	}
}
