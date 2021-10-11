package gorules

import "github.com/quasilyte/go-ruleguard/dsl"

// https://github.com/uber-go/guide/blob/master/style.md#local-variable-declarations
func uberLocalVarDecl(m dsl.Matcher) {
	m.Match(`var $x = $y`).
		Where(!m["$$"].Node.Parent().Is(`File`)).
		Suggest(`$x := $y`).
		Report(`use := for local variables declaration`)

	m.Match(`var $x $_ = $y`).
		Where(!m["$$"].Node.Parent().Is(`File`)).
		Report(`use := for local variables declaration`)
}
