package gorules

import "github.com/quasilyte/go-ruleguard/dsl"

func dupBranchBody(m dsl.Matcher) {
	m.Match("if $*_ { $body } else { $body }").
		Report("identical if and else bodies")
}
