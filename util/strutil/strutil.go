package strutil

import (
	"fmt"
	"regexp"
)

var (
	alphaNumReg *regexp.Regexp
)

func init() {
	r, err := regexp.Compile("[^a-zA-Z0-9]+")
	if err != nil {
		panic(fmt.Sprintf("regex compile: %v", err))
	}
	alphaNumReg = r
}

func AlphaNum(s string) string {
	return alphaNumReg.ReplaceAllString(s, "-")
}
