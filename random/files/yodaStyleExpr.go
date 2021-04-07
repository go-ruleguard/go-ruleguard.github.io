package gorules

import "github.com/quasilyte/go-ruleguard/dsl"

func yodaStyleExpr(m dsl.Matcher) {
	m.Match(`nil != $x`).Report(`rewrite as $x != nil`)
}
