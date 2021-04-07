package gorules

import "github.com/quasilyte/go-ruleguard/dsl"

func emptyError(m dsl.Matcher) {
	m.Match(`fmt.Errorf("")`, `errors.New("")`).
		Report(`empty errors are hard to debug`)
}
