package gorules

import "github.com/quasilyte/go-ruleguard/dsl"

func contextTodo(m dsl.Matcher) {
	m.Match(`context.TODO()`).
		Report(`might want to replace the context.TODO()`)
}
