package gorules

import "github.com/quasilyte/go-ruleguard/dsl"

func badLock(m dsl.Matcher) {
	m.Match(`$mu.Lock(); defer $mu.RUnlock()`).Report(`maybe $mu.RLock() was intended?`)
	m.Match(`$mu.RLock(); defer $mu.Unlock()`).Report(`maybe $mu.Lock() was intended?`)

	m.Match(`$mu1.Lock(); defer $mu2.Lock()`).
		Where(m["mu1"].Text == m["mu2"].Text).
		At(m["mu2"]).
		Report(`maybe defer $mu1.Unlock() was intended?`)
	m.Match(`$mu1.RLock(); defer $mu2.RLock()`).
		Where(m["mu1"].Text == m["mu2"].Text).
		At(m["mu2"]).
		Report(`maybe defer $mu1.RUnlock() was intended?`)
}
