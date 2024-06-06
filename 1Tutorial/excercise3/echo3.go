package excercise3

import (
	"strings"
)

func JoinArgs(args []string) string {
	return strings.Join(args[1:], " ")
}

func ConcatArgs(args []string) string {
	s, sep := "", ""
	for _, arg := range args {
		s += sep + arg
		sep = " "

	}
	return s
}
