package gorules

import "github.com/quasilyte/go-ruleguard/dsl"

func regexpMust(m dsl.Matcher) {
	m.Match(`regexp.Compile($pat)`, `regexp.CompilePOSIX($pat)`).
		Where(m["pat"].Const).
		Report(`can use MustCompile for const patterns`)
}
