package gorules

import "github.com/quasilyte/go-ruleguard/dsl"

func underlyingError(m dsl.Matcher) {
	/* See https://twitter.com/dvyukov/status/1174698980208513024 */
	m.Match(`type $x error`).
		Report(`error as underlying type is probably a mistake`).
		Suggest(`type $x struct { error }`)
}
