package gorules

import "github.com/quasilyte/go-ruleguard/dsl"

func selfAssign(m dsl.Matcher) {
	m.Match("$x = $x").Report("useless self-assignment")
}
