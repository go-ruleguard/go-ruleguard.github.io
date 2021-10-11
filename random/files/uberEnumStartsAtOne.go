package gorules

import "github.com/quasilyte/go-ruleguard/dsl"

// https://github.com/uber-go/guide/blob/master/style.md#start-enums-at-one
func uberEnumStartsAtOne(m dsl.Matcher) {
	m.Match(`const ($_ $_ = iota; $*_)`).
		Report(`enums should start from 1, not 0; use iota+1`)
}
