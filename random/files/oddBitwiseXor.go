package gorules

import "github.com/quasilyte/go-ruleguard/dsl"

func oddBitwiseXor(m dsl.Matcher) {
	m.Match(
		"2 ^ $_",
		"10 ^ $_",
	).Report("caret (^) is not exponentiation; it's a bitwise xor")
}
