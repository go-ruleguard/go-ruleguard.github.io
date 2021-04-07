package gorules

import "github.com/quasilyte/go-ruleguard/dsl"

func directAtomicAssign(m dsl.Matcher) {
	m.Match(`$x = atomic.AddInt32(&$x, $_)`,
		`$x = atomic.AddInt64(&$x, $_)`,
		`$x = atomic.AddUint32(&$x, $_)`,
		`$x = atomic.AddUint64(&$x, $_)`,
		`*$x = atomic.AddInt32($x, $_)`,
		`*$x = atomic.AddInt64($x, $_)`,
		`*$x = atomic.AddUint32($x, $_)`,
		`*$x = atomic.AddUint64($x, $_)`,
	).Report(`direct assignment to atomic value`)
}
