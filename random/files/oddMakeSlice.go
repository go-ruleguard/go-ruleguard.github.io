package gorules

import "github.com/quasilyte/go-ruleguard/dsl"

func oddMakeSlice(m dsl.Matcher) {
	m.Match(`new([$cap]$typ)[:$len]`).
		Report(`rewrite as 'make([]$typ, $len, $cap)'`)
}
