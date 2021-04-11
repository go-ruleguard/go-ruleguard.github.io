package gorules

import "github.com/quasilyte/go-ruleguard/dsl"

func imageRect(m dsl.Matcher) {
	m.Match(`image.Rectangle{Max: $max, Min: $min}`).
		Report(`Min field is mentioned before Max field`).
		Suggest(`image.Rectangle{Min: $min, Max: $max}`)

	m.Match(
		`image.Rectangle{image.Pt($x0, $y0), image.Pt($x1, $y1)}`,
		`image.Rectangle{Min: image.Pt($x0, $y0), Max: image.Pt($x1, $y1)}`,
	).Report(`could use image.Rect() helper function`).
		Suggest(`image.Rect($x0, $y0, $x1, $y1`)
}
