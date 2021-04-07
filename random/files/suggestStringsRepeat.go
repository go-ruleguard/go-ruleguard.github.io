package gorules

import "github.com/quasilyte/go-ruleguard/dsl"

func suggestStringsRepeat(m dsl.Matcher) {
	m.Match(`$s := make([]string, $n); for $i := range $s { $s[$i] = $x }`,
		`$s := make([]string, $n); for $i := 0; $i < len($s); $i++ { $s[$i] = $x }`).
		Suggest(`strings.Repeat($x, $i)`)
}
