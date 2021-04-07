package gorules

import "github.com/quasilyte/go-ruleguard/dsl"

func blankAssign(m dsl.Matcher) {
	m.Match(`_ = $v`).
		Where(m["v"].Pure).
		Report(`please remove the assignment to _`)
}
