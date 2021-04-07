package gorules

import "github.com/quasilyte/go-ruleguard/dsl"

func nilValReturn(m dsl.Matcher) {
	m.Match(`if $*_; $v == nil { return $v }`).
		Report(`returned expr is always nil; replace $v with nil`)
}
