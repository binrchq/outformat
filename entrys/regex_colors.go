package entrys

import "regexp"

type RegexColors struct {
	Pattern *regexp.Regexp
	Style   Style
}
