package gorules

import "github.com/quasilyte/go-ruleguard/dsl"

// https://github.com/uber-go/guide/blob/master/style.md#defer-to-clean-up
func uberDeferCleanup(m dsl.Matcher) {
	m.Match(`$mu.Lock(); $next`).
		Where(!m["next"].Text.Matches(`defer .*\.Unlock\(\)`)).
		Report(`$mu.Lock() should be followed by a deferred Unlock`)

	m.Match(`$res, $err := $open($*_); if $*_ { $*_ }; $next`).
		Where(m["res"].Type.Implements(`io.Closer`) &&
			m["err"].Type.Implements(`error`) &&
			!m["next"].Text.Matches(`defer .*[cC]lose`)).
		Report(`$res.Close() should be deferred right after the $open error check`)
}
