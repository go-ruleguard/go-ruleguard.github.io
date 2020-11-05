# [Ruleguard by example](https://go-ruleguard.github.io/by-example/): Multi-pattern rules

```go
package gorules

import "github.com/quasilyte/go-ruleguard/dsl/fluent"

func boolLiteralInExpr(m fluent.Matcher) {
	// The following rule includes 4 patterns: 2 for each operator and bool value.
	// Every pattern in this list is called "an alternative".
	// Alternatives are tried out one by one, until we have a match.
	m.Match(`$x == true`,
		`$x != true`,
		`$x == false`,
		`$x != false`).
		Report(`omit bool literal in expression`)
}
```

<pre style="color: white; background-color: black">
$ ruleguard -c 0 -rules rules.go main.go
main.go:7:10: omit bool literal in expression
7		println(cond == true)
main.go:8:10: omit bool literal in expression
8		println(cond == false)
main.go:9:10: omit bool literal in expression
9		println(cond != true)
main.go:10:10: omit bool literal in expression
10		println(cond != false)
</pre>

<details><summary>main.go</summary>

```go
package main

func main() {
	var cond bool
	var boolVar bool

	println(cond == true)
	println(cond == false)
	println(cond != true)
	println(cond != false)

	// No warnings for these:
	println(cond == boolVar)
	println(cond == boolVar)
	println(cond != boolVar)
	println(cond != boolVar)
}
```

</details>

<hr>

**Notes**:

* Use alternations to avoid the unwanted rules repetition
* Use the same submatch variable names in all alternatives

<table><tr>
<td><a href="index">To index</a></td>
<td><a href="multi-statement-patterns">Next: Multi-statement patterns</a></td>
<td><a href="https://github.com/go-ruleguard/go-ruleguard.github.io/edit/master/by-example/multi-pattern-rules.md">Edit this page</a></td>
</tr></table>
