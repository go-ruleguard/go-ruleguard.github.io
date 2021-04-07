package gorules

import "github.com/quasilyte/go-ruleguard/dsl"

func suggestStringsReplaceAll(m dsl.Matcher) {
	m.Match(`strings.Replace($_, $_, $_, -1)`).
		Report(`use strings.ReplaceAll`)
}
