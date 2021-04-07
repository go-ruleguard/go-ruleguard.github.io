package gorules

import "github.com/quasilyte/go-ruleguard/dsl"

func oddBitwiseOr(m dsl.Matcher) {
	m.Match("$x | $x",
		"$x | ^$x",
		"^$x | $x").
		Report("odd bitwise OR")
}
