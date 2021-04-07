package gorules

import "github.com/quasilyte/go-ruleguard/dsl"

func flagDeref(m dsl.Matcher) {
	m.Match(`*flag.Bool($*_)`,
		`*flag.Float64($*_)`,
		`*flag.Duration($*_)`,
		`*flag.Int($*_)`,
		`*flag.Int64($*_)`,
		`*flag.String($*_)`,
		`*flag.Uint($*_)`,
		`*flag.Uint64($*_)`).
		Report(`immediate deref in $$ is most likely an error`)
}
