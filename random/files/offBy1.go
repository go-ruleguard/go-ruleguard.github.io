package gorules

import "github.com/quasilyte/go-ruleguard/dsl"

func offBy1(m dsl.Matcher) {
	m.Match(`$s[len($s)]`).
		Where(m["s"].Type.Is(`[]$elem`) && m["s"].Pure).
		Report(`index expr always panics; maybe you wanted $s[len($s)-1]?`)
}
