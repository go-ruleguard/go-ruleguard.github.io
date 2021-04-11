package gorules

import "github.com/quasilyte/go-ruleguard/dsl"

func imagePt(m dsl.Matcher) {
	m.Match(`image.Pt(0, 0)`, `image.Point{0, 0}`).
		Report(`zero point should be written as image.Point{}`).
		Suggest(`image.Point{}`)

	m.Match(`image.Point{$x, $y}`).
		Report(`could use image.Pt() helper function`).
		Suggest(`image.Pt($x, $y`)
}
