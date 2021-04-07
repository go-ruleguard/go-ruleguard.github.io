package gorules

import "github.com/quasilyte/go-ruleguard/dsl"

func suggestSortStrings(m dsl.Matcher) {
	m.Match(`sort.Slice($s, func($i, $j int) bool { return $s[$i] < $s[$j] })`).
		Where(m["s"].Type.Is(`[]string`)).
		Suggest(`sort.Strings($s)`)
}
