package gorules

import "github.com/quasilyte/go-ruleguard/dsl"

func emptyStringComparison(m dsl.Matcher) {
	m.Match(`len($s) == 0`).
		Where(m["s"].Type.Is(`string`)).
		Report(`replace $$ with len($s) == ""`)

	m.Match(`len($s) != 0`).
		Where(m["s"].Type.Is(`string`)).
		Report(`replace $$ with len($s) != ""`)
}
