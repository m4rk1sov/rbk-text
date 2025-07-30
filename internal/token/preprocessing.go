package token

import (
	"github.com/m4rk1sov/rbk-text/internal/parser"
	"regexp"
	"strings"
)

var contractions = map[string]bool{
	"I'm": true, "I've": true, "I'd": true, "I'll": true,
	"You're": true, "You've": true, "You'd": true, "You'll": true,
	"He's": true, "He'd": true, "He'll": true,
	"She's": true, "She'd": true, "She'll": true,
	"It's": true, "It'd": true, "It'll": true,
	"We're": true, "We've": true, "We'd": true, "We'll": true,
	"They're": true, "They've": true, "They'd": true, "They'll": true,
	"Somebody's": true, "Somebody'd": true, "Somebody'll": true,
	"Someone's": true, "Someone'd": true, "Someone'll": true,
	"Something's": true, "Something'd": true, "Something'll": true,

	"Who's": true, "Who're": true, "Who'd": true, "Who've": true, "Who'll": true,
	"What's": true, "What're": true, "What'd": true, "What've": true, "What'll": true,
	"When's": true, "When're": true, "When'd": true, "When've": true, "When'll": true,
	"Where's": true, "Where're": true, "Where'd": true, "Where've": true, "Where'll": true,
	"Why's": true, "Why're": true, "Why'd": true, "Why've": true, "Why'll": true,
	"How's": true, "How're": true, "How'd": true, "How've": true, "How'll": true,
	"Which's": true, "Which're": true, "Which'd": true, "Which've": true, "Which'll": true,

	"This's": true, "This'd": true, "This'll": true,
	"These're": true, "These'd": true, "These'll": true,
	"That's": true, "That'd": true, "That'll": true,
	"Those're": true, "Those'd": true, "Those'll": true,
	"Here's": true, "Here're": true, "Here'd": true, "Here'll": true,
	"There's": true, "There're": true, "There'd": true, "There'll": true,

	"Ain't": true, "Isn't": true, "Aren't": true, "Wasn't": true, "Weren't": true,
	"Won't": true, "Will've": true, "Can't": true, "Couldn't": true, "Could've": true,
	"Shouldn't": true, "Should've": true, "Wouldn't": true, "Would've": true,
	"Mightn't": true, "Might've": true, "Mustn't": true, "Must've": true,
	"Don't": true, "Doesn't": true, "Didn't": true, "Haven't": true, "Hasn't": true, "Hadn't": true,

	"Gimme": true, "Cuz": true, "Cause": true, "Finna": true, "Imma": true, "Gonna": true,
	"Wanna": true, "Gotta": true, "Hafta": true, "Woulda": true, "Coulda": true, "Shoulda": true,
	"Ma'am": true, "Howdy": true, "Let's": true, "Y'all": true, "A'ight": true, "Amn't": true,
	"Ol'": true, "Cept": true,
}

var (
	// fixing cases like I   ' m you' re and etc.
	contractionRegex = regexp.MustCompile(`\b(\w+)\s*'\s*(\w+)\b`)
	// no space before punctuation
	spaceBeforePunctuation = regexp.MustCompile(`\s+([.,!?;:])`)
	// space after a punctuation and special characters
	spaceAfterPunctuation       = regexp.MustCompile(`([.,!?;:])(\w)`)
	spaceAfterSpecialCharacters = regexp.MustCompile(`([.,!?;:])([('"])`)
	// spacing inside single quotes
	singleQuoteSpace = regexp.MustCompile(`(?:^|\s)'\s*([^']*?)\s*'([.,!?;:]*)`)
	// spacing inside double quotes
	// \s*"\s*([^"]*?)\s*"\s* was old one
	doubleQuoteSpace = regexp.MustCompile(`(?:^|\s)"\s*([^"]*?)\s*"([.,!?;:]*)`)
	// multiple spaces to single
	multipleSpaces = regexp.MustCompile(`\s+`)
)

func isContraction(word string) bool {
	return contractions[word] || contractions[parser.ToTitle(strings.ToLower(word))] ||
		contractions[strings.ToLower(word)] || contractions[strings.ToUpper(word)]
}

// Normalize is a process to clean the raw string
func Normalize(raw string) string {
	result := multipleSpaces.ReplaceAllString(raw, " ")

	// fix spaced-out contractions (e.g., "I ' m" to "I'm")
	result = contractionRegex.ReplaceAllStringFunc(result, func(match string) string {
		parts := contractionRegex.FindStringSubmatch(match)
		if len(parts) == 3 {
			candidate := parts[1] + "'" + parts[2]
			if isContraction(candidate) {
				return candidate
			}
		}
		return match
	})

	// spacing around quotes
	result = singleQuoteSpace.ReplaceAllString(result, " '$1'$2")
	result = doubleQuoteSpace.ReplaceAllString(result, ` "$1"$2 `)

	// fix spacing around punctuation
	result = spaceBeforePunctuation.ReplaceAllString(result, "$1")
	result = spaceAfterPunctuation.ReplaceAllString(result, "$1 $2")
	result = spaceAfterSpecialCharacters.ReplaceAllString(result, "$1 $2")

	// in case when iside quotes
	//result = strings.ReplaceAll(result, `. "`, `."`)
	//result = strings.ReplaceAll(result, `, "`, `,"`)
	//result = strings.ReplaceAll(result, `! "`, `!"`)
	//result = strings.ReplaceAll(result, `? "`, `?"`)
	//result = strings.ReplaceAll(result, `; "`, `;"`)
	//
	//result = strings.ReplaceAll(result, `. '`, `.'`)
	//result = strings.ReplaceAll(result, `, '`, `,'`)
	//result = strings.ReplaceAll(result, `! '`, `!'`)
	//result = strings.ReplaceAll(result, `? '`, `?'`)
	//result = strings.ReplaceAll(result, `; '`, `;'`)

	// cleanup
	result = multipleSpaces.ReplaceAllString(result, " ")

	return strings.TrimSpace(result)
}
