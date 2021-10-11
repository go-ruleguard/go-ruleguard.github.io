package gorules

import "github.com/quasilyte/go-ruleguard/dsl"

// https://github.com/uber-go/guide/blob/master/style.md#handle-type-assertion-failures
func uberUncheckedTypeAssert(m dsl.Matcher) {
	m.Match(
		// simple assignments
		`$_ := $_.($_)`,
		`$_ = $_.($_)`,
		// as a function call argument
		`$_($*_, $_.($_), $*_)`,
		// as a composite literal member
		`$_{$*_, $_.($_), $*_}`,
		`$_{$*_, $_: $_.($_), $*_}`,
	).Report(`avoid unchecked type assertions as they can panic`)
}
