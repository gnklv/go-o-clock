package types

import "unicode/utf8"

type Symbol []string

func (s Symbol) Width() int {
	return utf8.RuneCountInString(s[0])
}

func (s Symbol) Height() int {
	return len(s)
}
