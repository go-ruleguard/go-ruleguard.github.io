package gorules

import "github.com/quasilyte/go-ruleguard/dsl"

func oddArgsOrder(m dsl.Matcher) {
	m.Match(`strings.HasPrefix($s1, $s2)`).
		Where(m["s1"].Const && !m["s2"].Const).
		Suggest(`strings.HasPrefix($s2, $s1)`)

	m.Match(`strings.HasSuffix($s1, $s2)`).
		Where(m["s1"].Const && !m["s2"].Const).
		Suggest(`strings.HasPrefix($s2, $s1)`)

	m.Match(`strings.Contains($s1, $s2)`).
		Where(m["s1"].Const && !m["s2"].Const).
		Suggest(`strings.Contains($s2, $s1)`)
}
