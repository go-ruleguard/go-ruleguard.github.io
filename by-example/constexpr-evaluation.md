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
}

func badCond(m fluent.Matcher) {
	// If we use Value.Int() on the right-hand-side, we can achieve even more
	// and compare how to constexpr values compare to each other.

	m.Match(`$x < $a && $x > $b`).
		Where(m["a"].Value.Int() <= m["b"].Value.Int()).
		Report("the condition is always false because $a <= $b")

	m.Match(`$x > $a && $x < $b`).
		Where(m["a"].Value.Int() >= m["b"].Value.Int()).
		Report("the condition is always false because $a >= $b")
}
```

<pre style="color: white; background-color: black">
$ ruleguard -c 0 -rules rules.go main.go
main.go:11:6: suggestion: strings.ReplaceAll(s, toReplace, replacement)
11		_ = strings.Replace(s, toReplace, replacement, -1)
main.go:12:6: suggestion: strings.ReplaceAll(s, toReplace, replacement)
12		_ = strings.Replace(s, toReplace, replacement, -10)
main.go:13:6: suggestion: strings.ReplaceAll(s, " ", "")
13		_ = strings.Replace(s, " ", "", 10-11) // Const-folded and matched
main.go:14:6: suggestion: strings.ReplaceAll(s, "/", `\`)
14		_ = strings.Replace(s, "/", `\`, negativeConst)
main.go:20:5: the condition is always false because -10 <= 10
20		if x < -10 && x > 10 {
main.go:23:5: the condition is always false because ten+1 >= -10
23		if x > ten+1 && x < -10 {
</pre>

<details><summary>main.go</summary>

```go
package main

import "strings"

func main() {
	var s, toReplace, replacement string
	const negativeConst = -2

	_ = strings.Replace(s, toReplace, replacement, 1) // Not matched

	_ = strings.Replace(s, toReplace, replacement, -1)
	_ = strings.Replace(s, toReplace, replacement, -10)
	_ = strings.Replace(s, " ", "", 10-11) // Const-folded and matched
	_ = strings.Replace(s, "/", `\`, negativeConst)

	badCond(1, 2)
}

func badCond(x, y int) {
	if x < -10 && x > 10 {
	}
	const ten = 10
	if x > ten+1 && x < -10 {
	}

	// Don't know what value `y` have.
	if x < y && x > 10 {
	}
	if x < -10 && y > 10 {
	}
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
