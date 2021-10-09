package gorules

import "github.com/quasilyte/go-ruleguard/dsl"

func contextWithValue(m dsl.Matcher) {
	// https://utcc.utoronto.ca/~cks/space/blog/programming/GoContextValueMistake

	m.Match(`context.WithValue($*_)`).
		Where(m["$$"].Node.Parent().Is(`ExprStmt`)).
		Report(`context.WithValue result should not be ignored`)
}
