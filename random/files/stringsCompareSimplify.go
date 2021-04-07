package gorules

import "github.com/quasilyte/go-ruleguard/dsl"

func stringsCompareSimplify(m dsl.Matcher) {
	m.Match(`strings.Compare($s1, $s2) == 0`).
		Suggest(`$s1 == $s2`)
	m.Match(`strings.Compare($s1, $s2) < 0`,
		`strings.Compare($s1, $s2) == -1`).
		Suggest(`$s1 < $s2`)
	m.Match(`strings.Compare($s1, $s2) > 0`,
		`strings.Compare($s1, $s2) == 1`).
		Suggest(`$s1 > $s2`)
}
