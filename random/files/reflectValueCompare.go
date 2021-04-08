package gorules

import "github.com/quasilyte/go-ruleguard/dsl"

func reflectValueCompare(m dsl.Matcher) {
	m.Match(`reflect.DeepEqual($x, $y)`).
		Where(m["x"].Type.Is(`reflect.Value`) &&
			m["y"].Type.Is(`reflect.Value`) &&
			!m["x"].Node.Is(`CompositeLit`) &&
			!m["y"].Node.Is(`CompositeLit`)).
		Report("avoid using reflect.DeepEqual with reflect.Value").
		Suggest("reflect.DeepEqual($x.Interface(), $y.Interface()")

	m.Match(`$x == $y`, `$x != $y`).
		Where(m["x"].Type.Is(`reflect.Value`) &&
			m["y"].Type.Is(`reflect.Value`) &&
			!m["x"].Node.Is(`CompositeLit`) &&
			!m["y"].Node.Is(`CompositeLit`)).
		Report("avoid using == and != with reflect.Value")
}
