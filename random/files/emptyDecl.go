package gorules

import "github.com/quasilyte/go-ruleguard/dsl"

func emptyDecl(m dsl.Matcher) {
	m.Match(`var()`).Report(`empty var() block`)
	m.Match(`const()`).Report(`empty const() block`)
	m.Match(`type()`).Report(`empty type() block`)
}
