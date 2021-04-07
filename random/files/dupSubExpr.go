package gorules

import "github.com/quasilyte/go-ruleguard/dsl"

func dupSubExpr(m dsl.Matcher) {
	m.Match(`$x || $x`,
		`$x && $x`,
		`$x | $x`,
		`$x & $x`).
		Where(m["x"].Pure).
		Report(`suspicious identical LHS and RHS`)
}
