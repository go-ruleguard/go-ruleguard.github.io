package gorules

import "github.com/quasilyte/go-ruleguard/dsl"

func errorVarPrefix(m dsl.Matcher) {
	// For all global vars, require Err
	// prefix for variables of type error.
	m.Match(`var $v = $_`).
		Where(m["$$"].Node.Parent().Is(`File`) &&
			m["v"].Type.Implements(`error`) &&
			!m["v"].Text.Matches(`^Err`)).
		Report(`error vars should be prefixed with Err`)
}
