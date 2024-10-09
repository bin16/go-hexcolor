package hexcolor

import "regexp"

var rx = regexp.MustCompile("^#?([0-9a-fA-F]{3}|[0-9a-fA-F]{4}|[0-9a-fA-F]{6}|[0-9a-fA-F]{8})$")

func IsValid(s string) bool {
	return rx.MatchString(s)
}
