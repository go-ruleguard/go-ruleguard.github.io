package gorules

import "github.com/quasilyte/go-ruleguard/dsl"

func emptySliceMapRange(m dsl.Matcher) {
	m.Match(`if len($xs) != 0 { for range $xs { $*_ } }`,
		`if len($xs) != 0 { for $i := range $xs { $*_ } }`,
		`if len($xs) != 0 { for _, $x := range $xs { $*_ } }`,
		`if len($xs) != 0 { for _, $x = range $xs { $*_ } }`,
		`if $xs != nil { for range $xs { $*_ } }`,
		`if $xs != nil { for $i := range $xs { $*_ } }`,
		`if $xs != nil { for _, $x := range $xs { $*_ } }`,
		`if $xs != nil { for _, $x = range $xs { $*_ } }`).
		Report(`check on $xs is redundant, empty/nil slices and maps can be safely iterated`)
}
