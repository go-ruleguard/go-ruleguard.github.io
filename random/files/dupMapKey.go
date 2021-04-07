package gorules

import "github.com/quasilyte/go-ruleguard/dsl"

func dupMapKey(m dsl.Matcher) {
	m.Match(`map[$_]$_{$*_, $k: $_, $*_, $k: $_, $*_}`).
		Where(m["k"].Pure).
		Report(`suspicious duplicate key $k`).
		At(m["k"])
}
