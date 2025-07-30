package tests

import (
	"github.com/m4rk1sov/rbk-text/internal/parser"
	"github.com/m4rk1sov/rbk-text/internal/token"
	"testing"
)

func TestWorkflow(t *testing.T) {
	var tests = []struct {
		in  string
		out string
	}{
		{`She '  d said ' hello ', then " goodbye ".\n`, `She'd said 'hello', then "goodbye". \n`},
		{`   "  This is what I'  ll call an ' tricky' situation  . " (up) `, `"This is what I'll call a 'tricky' SITUATION."`},
		{`((word1    2word word3  ,   word4)  ), (cap, 3)`, `((word1 2word Word3, Word4)),`},
		{`"    \n\n\n\nnnn word \nnn n \\\n \n" (cap, 5)    `, `"\n\n\n\nnnn N \nnn N \\\n \N"`},
		{`'word (cap)'`, `'word (cap)'`},
		{`'word2' (cap)`, `'Word2'`},
		{`word"test test2 \n test3"(up, 2)"str1 str2 str3 you'll  . "    (cap, 2)  `, `word "test test2 \N TEST3" "str1 str2 Str3 You'll."`},
		{`'   I '  m here and    will "always   be for you in "  11110 "" (bin)   years '`, `'I'm here and will "always be for you in" 30 "" years '`},
		//{``, ``},
		//{``, ``},
		//{``, ``},
		//{``, ``},
		//{``, ``},
		//{``, ``},
		//{``, ``},
		//{``, ``},
		//{``, ``},
	}

	for _, tt := range tests {
		t.Run(tt.in, func(t *testing.T) {
			clean := token.Normalize(tt.in)
			tokens, err := token.Tokenize(clean)
			if err != nil {
				t.Errorf("failed to tokenize the text: %v", err)
			}
			transformed, err := parser.TransformTokens(tokens)
			if err != nil {
				t.Errorf("failed to tokenize the text: %v", err)
			}
			formatted := parser.JoinTokens(transformed)
			formatted = token.Normalize(formatted)
			if formatted != tt.out {
				t.Errorf("expected: %q, but got: %q", tt.out, formatted)
			}
		})
	}
}
