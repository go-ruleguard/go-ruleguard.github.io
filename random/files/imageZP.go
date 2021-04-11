package gorules

import "github.com/quasilyte/go-ruleguard/dsl"

func imageZP(m dsl.Matcher) {
	m.Match(`image.ZP`).
		Report(`image.ZP is deprecated, use image.Point{} instead`).
		Suggest(`image.Point{}`)
}
