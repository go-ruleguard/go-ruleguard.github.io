package gorules

import "github.com/quasilyte/go-ruleguard/dsl"

func switchConst(m dsl.Matcher) {
	m.Match("switch $x { $*_ }", "switch $*_; $x { $*_ }").
		Where(m["x"].Const && !m["x"].Text.Matches(`^runtime\.`)).
		Report("suspicious switch over a constant tag")
}
