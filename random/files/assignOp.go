package gorules

import "github.com/quasilyte/go-ruleguard/dsl"

func assignOp(m dsl.Matcher) {
	m.Match(`$x = $x + 1`).Report(`can simplify to $x++`)
	m.Match(`$x = $x - 1`).Report(`can simplify to $x--`)

	m.Match(`$x = $x + $y`).Report(`can simplify to $x+=$y`)
	m.Match(`$x = $x - $y`).Report(`can simplify to $x-=$y`)
	m.Match(`$x = $x * $y`).Report(`can simplify to $x*=$y`)
}
