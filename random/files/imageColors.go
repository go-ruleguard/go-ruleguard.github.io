package gorules

import "github.com/quasilyte/go-ruleguard/dsl"

func imageColors(m dsl.Matcher) {
	m.Match(`color.Gray16{}`, `color.Gray16{0}`, `color.Gray16{$_: 0}`).
		Suggest(`color.Black`)

	m.Match(`color.Gray16{$x}`, `color.Gray16{$_: $x}`).
		Where(m["x"].Value.Int() == 0xffff).
		Suggest(`color.White`)

	m.Match(`color.Alpha16{}`, `color.Alpha16{0}`, `color.Alpha16{$_: 0}`).
		Suggest(`color.Transparent`)

	m.Match(`color.Alpha16{$x}`, `color.Alpha16{$_: $x}`).
		Where(m["x"].Value.Int() == 0xffff).
		Suggest(`color.Opaque`)
}
