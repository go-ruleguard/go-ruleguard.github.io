package gorules

import "github.com/quasilyte/go-ruleguard/dsl"

func boolLiteralComparison(m dsl.Matcher) {
	m.Match(
		`$x == true`,
		`$x != true`,
		`$x == false`,
		`$x != false`,
	).Report(`omit bool literal in expression`)
}
