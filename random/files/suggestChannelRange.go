package gorules

import "github.com/quasilyte/go-ruleguard/dsl"

func suggestChannelRange(m dsl.Matcher) {
	// Type check of $ch is not strictly needed, since
	// Go would not permit having non-chan type in the select case clause.
	m.Match(`for { select { case $_ := <-$ch: $*_ } }`).
		Report(`can use for range over $ch`)
}
