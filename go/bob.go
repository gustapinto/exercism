package bob

import (
	"strings"
	"unicode"
)

func Hey(remark string) string {
	remark = strings.TrimSpace(remark)
	isNothing := func() bool {
		return remark == ""
	}
	isAQuestion := func() bool {
		return strings.HasSuffix(remark, "?")
	}
	isAYell := func() bool {
		return strings.IndexFunc(remark, unicode.IsLetter) >= 0 &&
			remark == strings.ToUpper(remark)
	}

	switch {
	case isNothing():
		return "Fine. Be that way!"
	case isAQuestion() && isAYell():
		return "Calm down, I know what I'm doing!"
	case isAQuestion():
		return "Sure."
	case isAYell():
		return "Whoa, chill out!"
	default:
		return "Whatever."
	}
}
