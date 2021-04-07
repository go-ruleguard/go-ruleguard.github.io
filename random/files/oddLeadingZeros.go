package gorules

import "github.com/quasilyte/go-ruleguard/dsl"

func oddLeadingZeros(m dsl.Matcher) {
	m.Match(
		"64 - bits.LeadingZeros64($_)",
		"32 - bits.LeadingZeros32($_)",
		"16 - bits.LeadingZeros16($_)",
		"8 - bits.LeadingZeros8($_)",
	).Report("odd math/bits expression; use bits.Len*() instead?")
}
