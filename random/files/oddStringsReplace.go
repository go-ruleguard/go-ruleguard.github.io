package gorules

import "github.com/quasilyte/go-ruleguard/dsl"

func oddStringsReplace(m dsl.Matcher) {
	m.Match(`strings.Replace($_, $_, $_, 0)`).
		Report(`n=0 argument does nothing, maybe n=-1 is indended?`)
}
