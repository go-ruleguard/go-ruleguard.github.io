package gorules

import "github.com/quasilyte/go-ruleguard/dsl"

func arrayDeref(m dsl.Matcher) {
	m.Match(`(*$arr)[$_]`).
		Where(m["arr"].Type.Is(`*[$_]$_`)).
		Report(`explicit array deref is redundant`)
}
