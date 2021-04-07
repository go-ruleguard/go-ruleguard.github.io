package gorules

import "github.com/quasilyte/go-ruleguard/dsl"

func oddAssign(m dsl.Matcher) {
	m.Match("$x = $y; $y = $x").Report("odd sequence of assignments")
}
