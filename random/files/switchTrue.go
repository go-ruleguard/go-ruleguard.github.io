package gorules

import "github.com/quasilyte/go-ruleguard/dsl"

func switchTrue(m dsl.Matcher) {
	m.Match(`switch true {$*_}`).Report(`can omit true in switch`)
}
