package gorules

import "github.com/quasilyte/go-ruleguard/dsl"

func suggestStringsSuffixPrefix(m dsl.Matcher) {
	m.Match(`len($s) >= len($x) && $s[:len($x)] == $x`).
		Suggest(`strings.HasPrefix($s, $x)`)

	m.Match(`len($s) >= len($x) && $s[len($s)-len($x):] == $x`).
		Suggest(`strings.HasSuffix($s, $x)`)
}
