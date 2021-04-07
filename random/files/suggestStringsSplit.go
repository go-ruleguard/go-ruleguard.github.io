package gorules

import "github.com/quasilyte/go-ruleguard/dsl"

func suggestStringsSplit(m dsl.Matcher) {
	m.Match(`strings.SplitN($_, $_, -1)`).
		Report(`use strings.Split`)
}
