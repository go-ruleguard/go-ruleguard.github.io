package gorules

import "github.com/quasilyte/go-ruleguard/dsl"

func badRangeCheck(m dsl.Matcher) {
	m.Match(`$x < $a && $x > $b`).
		Where(m["a"].Value.Int() <= m["b"].Value.Int()).
		Report("the condition is always false because $a <= $b")

	m.Match(`$x > $a && $x < $b`).
		Where(m["a"].Value.Int() >= m["b"].Value.Int()).
		Report("the condition is always false because $a >= $b")
}
