package gorules

import "github.com/quasilyte/go-ruleguard/dsl"

func durationExprSimplify(m dsl.Matcher) {
	m.Match(`time.Duration($x) * time.Second`).
		Where(m["x"].Const).
		Report(`rewrite as '$x * time.Second'`)

	m.Match(`int64(time.Since($t) / time.Microsecond)`).
		Suggest(`time.Since($t).Microseconds()`)
	m.Match(`int64(time.Since($t) / time.Millisecond)`).
		Suggest(`time.Since($t).Milliseconds()`)
}
