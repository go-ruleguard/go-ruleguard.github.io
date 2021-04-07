package gorules

import "github.com/quasilyte/go-ruleguard/dsl"

func appendAssign(m dsl.Matcher) {
	m.Match(`$x = append($y, $_, $*_)`).
		Where(m["x"].Text != m["y"].Text &&
			m["x"].Text != "_" &&
			m["x"].Node.Is(`Ident`) &&
			m["y"].Node.Is(`Ident`)).
		Report("append result not assigned to the same slice")
}
