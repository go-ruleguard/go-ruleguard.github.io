package gorules

import "github.com/quasilyte/go-ruleguard/dsl"

func newDeref(m dsl.Matcher) {
	m.Match(`*new(bool)`).Report(`replace $$ with false`)
	m.Match(`*new(string)`).Report(`replace $$ with ""`)
	m.Match(`*new(int)`).Report(`replace $$ with 0`)
}
