package gorules

import "github.com/quasilyte/go-ruleguard/dsl"

func valSwap(m dsl.Matcher) {
	m.Match(`$tmp := $x; $x = $y; $y = $tmp`).
		Report(`can use parallel assignment like $x,$y=$y,$x`)
}
