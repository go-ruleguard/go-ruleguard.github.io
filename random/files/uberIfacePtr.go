package gorules

import "github.com/quasilyte/go-ruleguard/dsl"

// https://github.com/uber-go/guide/blob/master/style.md#pointers-to-interfaces
func uberIfacePtr(m dsl.Matcher) {
	m.Match(`*$x`).
		Where(m["x"].Type.Underlying().Is(`interface{ $*_ }`)).
		Report(`don't use pointers to an interface`)
}
