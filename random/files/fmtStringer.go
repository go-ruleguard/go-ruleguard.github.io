package gorules

import "github.com/quasilyte/go-ruleguard/dsl"

func fmtStringer(m dsl.Matcher) {
	m.Match(`fmt.Sprint($x)`).
		Where(m["x"].Type.Implements(`fmt.Stringer`)).
		Suggest(`$x.String()`)
}
