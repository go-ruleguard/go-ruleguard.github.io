package gorules

import "github.com/quasilyte/go-ruleguard/dsl"

func nopSlicing(m dsl.Matcher) {
	m.Match(`$s[:]`).
		Where(m["s"].Type.Is(`string`) || m["s"].Type.Is(`[]$_`)).
		Report(`can simplify $$ to $s`)
}
