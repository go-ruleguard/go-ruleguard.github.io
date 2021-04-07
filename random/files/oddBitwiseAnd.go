package gorules

import "github.com/quasilyte/go-ruleguard/dsl"

func oddBitwiseAnd(m dsl.Matcher) {
	m.Match("$x & $x",
		"$x & ^$x",
		"^$x & $x").
		Report("odd bitwise AND")
}
