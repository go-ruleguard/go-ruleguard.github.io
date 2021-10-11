package gorules

import "github.com/quasilyte/go-ruleguard/dsl"

// https://github.com/uber-go/guide/blob/master/style.md#channel-size-is-one-or-none
func uberChannelSize(m dsl.Matcher) {
	m.Match(`make(chan $_, $size)`).
		Where(m["size"].Value.Int() != 0 && m["size"].Value.Int() != 1).
		Report(`channels should have a size of one or be unbuffered`)
}
