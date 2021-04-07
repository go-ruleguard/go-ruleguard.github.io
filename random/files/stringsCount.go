package gorules

import "github.com/quasilyte/go-ruleguard/dsl"

func stringsCount(m dsl.Matcher) {
	m.Match(`strings.Count($_, $_) >= 0`).Report(`statement always true`)
	m.Match(`bytes.Count($_, $_) >= 0`).Report(`statement always true`)
}
