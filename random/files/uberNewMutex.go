package gorules

import "github.com/quasilyte/go-ruleguard/dsl"

// https://github.com/uber-go/guide/blob/master/style.md#zero-value-mutexes-are-valid
func uberNewMutex(m dsl.Matcher) {
	m.Match(`$mu := new(sync.Mutex); $mu.Lock()`).
		Report(`use zero mutex value instead, 'var $mu sync.Mutex'`).
		At(m["mu"])
}
