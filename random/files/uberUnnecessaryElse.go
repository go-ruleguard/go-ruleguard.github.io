package gorules

import "github.com/quasilyte/go-ruleguard/dsl"

// https://github.com/uber-go/guide/blob/master/style.md#unnecessary-else
func uberUnnecessaryElse(m dsl.Matcher) {
	m.Match(`var $v $_; if $cond { $v = $x } else { $v = $y }`).
		Where(m["y"].Pure).
		Report(`rewrite as '$v := $y; if $cond { $v = $x }'`)
}
