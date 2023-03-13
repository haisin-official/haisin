package utils

import (
	"regexp"
)

func CheckRegExp(reg string, str string) bool {
	return regexp.MustCompile(reg).Match([]byte(str))
}
