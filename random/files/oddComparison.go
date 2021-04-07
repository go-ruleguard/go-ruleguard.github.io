package gorules

import "github.com/quasilyte/go-ruleguard/dsl"

func oddComparison(m dsl.Matcher) {
	m.Match(
		"$x - $y == 0",
		"$x - $y != 0",
		"$x - $y < 0",
		"$x - $y <= 0",
		"$x - $y > 0",
		"$x - $y >= 0",
		"$x ^ $y == 0",
		"$x ^ $y != 0",
	).Report("odd comparison")
}
