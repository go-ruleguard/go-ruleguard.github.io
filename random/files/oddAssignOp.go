package gorules

import "github.com/quasilyte/go-ruleguard/dsl"

func oddAssignOp(m dsl.Matcher) {
	m.Match("$x += $x + $_",
		"$x += $x - $_").
		Report("odd += expression")

	m.Match("$x -= $x + $_",
		"$x -= $x - $_").
		Report("odd -= expression")
}
