package gorules

import "github.com/quasilyte/go-ruleguard/dsl"

func nopAppend(m dsl.Matcher) {
	m.Match(`append($_)`).
		Report(`append called with 1 argument does nothing`)
}
