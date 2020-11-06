# [Ruleguard by example](https://go-ruleguard.github.io/by-example/): Constexpr evaluation

```go
package gorules

import "github.com/quasilyte/go-ruleguard/dsl/fluent"

// It's possible to compute the expression value if operands are constant expressions:
// literals, named constants and some other special kinds of constexprs.

func replaceAll(m fluent.Matcher) {
	// Do you remember our older strings.Replace -> strings.ReplaceAll rules?
	// We matched $n argument as -1 literal.
	//
	// This is not exhaustive as any $n that is less than 0 will work identically.
	// So what we really want to do is to tell whether it's less than 0.

	m.Match(`strings.Replace($s, $old, $replacement, $n)`).
		Where(m["n"].Value.Int() < 0).
		Suggest(`strings.ReplaceAll($s, $old, $replacement)`)

	m.Match(`bytes.Replace($s, $old, $replacement, $n)`).
		Where(m["n"].Value.Int() < 0).
		Suggest(`bytes.ReplaceAll($s, $old, $replacement)`)
}
```

<pre style="color: white; background-color: black">
$ ruleguard -c 0 -rules rules.go main.go
main.go:14:6: suggestion: strings.ReplaceAll(s, toReplace, replacement)
14		_ = strings.Replace(s, toReplace, replacement, -1)
main.go:15:6: suggestion: strings.ReplaceAll(s, toReplace, replacement)
15		_ = strings.Replace(s, toReplace, replacement, -10)
main.go:16:6: suggestion: strings.ReplaceAll(s, " ", "")
16		_ = strings.Replace(s, " ", "", 10-11) // Const-folded and matched
main.go:17:6: suggestion: strings.ReplaceAll(s, "/", `\`)
17		_ = strings.Replace(s, "/", `\`, negativeConst)
main.go:18:6: suggestion: bytes.ReplaceAll([]byte(s), []byte("?"), nil)
18		_ = bytes.Replace([]byte(s), []byte("?"), nil, negativeConst*2)
</pre>

<details><summary>main.go</summary>

```go
package main

import (
	"bytes"
	"strings"
)

func main() {
	var s, toReplace, replacement string
	const negativeConst = -2

	_ = strings.Replace(s, toReplace, replacement, 1) // Not matched

	_ = strings.Replace(s, toReplace, replacement, -1)
	_ = strings.Replace(s, toReplace, replacement, -10)
	_ = strings.Replace(s, " ", "", 10-11) // Const-folded and matched
	_ = strings.Replace(s, "/", `\`, negativeConst)
	_ = bytes.Replace([]byte(s), []byte("?"), nil, negativeConst*2)
}
```

</details>

<hr>

**Notes**:

* If expression is not a constexpr, `Value.Int()` filter will fail and reject the match

<table><tr>
<td><a href="index">To index</a></td>
<td><a href="custom-locations">Next: Custom locations</a></td>
<td><a href="https://github.com/go-ruleguard/go-ruleguard.github.io/edit/master/by-example/constexpr-evaluation.md">Edit this page</a></td>
</tr></table>
