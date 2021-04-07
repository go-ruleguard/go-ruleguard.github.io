package gorules

import "github.com/quasilyte/go-ruleguard/dsl"

func suggestStringConcat(m dsl.Matcher) {
	m.Match(`fmt.Sprintf("%s%s", $a, $b)`).
		Where(m["a"].Type.Is(`string`) && m["b"].Type.Is(`string`)).
		Suggest(`$a+$b`)

	m.Match(`strings.Join([]string{$a, $b}, "")`).
		Suggest(`$a+$b`)

	m.Match(`strings.Join([]string{$a, $b}, $glue)`).
		Suggest(`$a+$glue+$b`)
}
