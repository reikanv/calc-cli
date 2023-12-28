package strpad

import (
	"fmt"
	"strings"
)

func absPad(pad int) int {
	if pad < 0 {
		return pad * -1
	}

	return pad
}

func Right(str string, pad int) string {
	whitespace := strings.Repeat(" ", absPad(pad-len(str)))
	return fmt.Sprintf("%s%s", str, whitespace)
}

func Left(str string, pad int) string {
	whitespace := strings.Repeat(" ", absPad(pad-len(str)))
	return fmt.Sprintf("%s%s", whitespace, str)
}

func TrimFloat(str string) string {
	return strings.TrimRight(strings.TrimRight(str, "0"), ".")
}
