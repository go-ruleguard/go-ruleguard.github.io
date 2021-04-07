package gorules

import "github.com/quasilyte/go-ruleguard/dsl"

func sloppyErrorCheck(m dsl.Matcher) {
	// Checking error without assigning it to a variably,
	// like in `if f() != nil`.
	m.Match(`$err != nil`,
		`$err == nil`).
		Where(!m["err"].Pure && m["err"].Type.Is(`error`)).
		Report(`assign $err to err and then do a nil check`)
}
