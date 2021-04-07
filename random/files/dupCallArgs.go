package gorules

import "github.com/quasilyte/go-ruleguard/dsl"

func dupCallArgs(m dsl.Matcher) {
	m.Match(`strings.Replace($_, $x, $x, $_)`).
		Report(`replace 'old' and 'new' parameters are identical`)
}
