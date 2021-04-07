package gorules

import "github.com/quasilyte/go-ruleguard/dsl"

func negationSimplify(m dsl.Matcher) {
	m.Match(`!!$x`).Report(`can simplify !!$x to $x`)
	m.Match(`!($x != $y)`).Suggest(`$x == $y`)
	m.Match(`!($x == $y)`).Suggest(`$x != $y`)
}
