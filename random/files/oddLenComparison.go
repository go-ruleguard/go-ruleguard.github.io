package gorules

import "github.com/quasilyte/go-ruleguard/dsl"

func oddLenComparison(m dsl.Matcher) {
	m.Match(`len($_) >= 0`).Report(`$$ is always true`)
	m.Match(`len($_) < 0`).Report(`$$ is always false`)

	m.Match(`len($s) <= 0`).Report(`$$ is never negative, can rewrite as len($s)==0`)
}
