package gorules

import "github.com/quasilyte/go-ruleguard/dsl"

func appendCombine(m dsl.Matcher) {
	m.Match(`$dst = append($x, $a); $dst = append($x, $b)`).
		Report(`$dst=append($x,$a,$b) is faster`)
}
