package gorules

import "github.com/quasilyte/go-ruleguard/dsl"

func timeEq(m dsl.Matcher) {
	m.Match("$t0 == $t1").Where(m["t0"].Type.Is("time.Time")).
		Report("using == with time.Time")
	m.Match("$t0 != $t1").Where(m["t0"].Type.Is("time.Time")).
		Report("using != with time.Time")
	m.Match(`map[$k]$v`).Where(m["k"].Type.Is("time.Time")).
		Report("map with time.Time keys are easy to misuse")
}
