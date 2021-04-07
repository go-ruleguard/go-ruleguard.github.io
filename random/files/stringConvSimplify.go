package gorules

import "github.com/quasilyte/go-ruleguard/dsl"

func stringConvSimplify(m dsl.Matcher) {
	m.Match(`copy($b, []byte($s))`).
		Where(m["s"].Type.Is(`string`)).
		Report(`can write copy($b, $s) without type conversion`)
}
