package gorules

import "github.com/quasilyte/go-ruleguard/dsl"

func dupCond(m dsl.Matcher) {
	m.Match(
		`$x && $x`,
		`$x && $_ && $x`,
		`$x && $_ && $_ && $x`).
		Report(`suspicious duplicated $x in && condition`)

	m.Match(
		`$x || $x`,
		`$x || $_ || $x`,
		`$x || $_ || $_ || $x`).
		Report(`suspicious duplicated $x in || condition`)
}
