package gorules

import "github.com/quasilyte/go-ruleguard/dsl"

func suggestStringsContains(m dsl.Matcher) {
	m.Match(`strings.Count($s1, $s2) > 0`,
		`strings.Count($s1, $s2) >= 1`).
		Suggest(`strings.Contains($s1, $s2)`)

	m.Match(`strings.Count($s1, $s2) == 0`).
		Suggest(`!strings.Contains($s1, $s2)`)
}
