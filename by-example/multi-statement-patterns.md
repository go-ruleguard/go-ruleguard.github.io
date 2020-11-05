# [Ruleguard by example](https://go-ruleguard.github.io/by-example/): Multi-statement patterns [:pencil:](https://github.com/go-ruleguard/go-ruleguard.github.io/blob/master/by-example/multi-statement-patterns.md)

```go
package gorules

import "github.com/quasilyte/go-ruleguard/dsl/fluent"

// Sometimes you want to match a statement that is **followed** or **preceded** by another statement.
// For this case we'll need a multi-statement patterns that gogrep does recognize.

func boolLiteralInExpr(m fluent.Matcher) {
	// If we want to specify $x statement that is followed by $y, we use `$x; $y` notation.
	// Here we match three statements:
	// 1. $x assigned to a temporary $tmp
	// 2. $x is re-assigned to $y
	// 3. $y is re-assigned to $tmp
	m.Match(`$tmp := $x; $x = $y; $y = $tmp`).
		Suggest(`$x, $y = $y, $x`)
}
```

<pre style="color: white; background-color: black">
$ ruleguard -c 0 -rules rules.go main.go
main.go:9:2: suggestion: obj1.value, obj2.value = obj2.value, obj1.value
9		tmp := obj1.value
</pre>

<details><summary>main.go</summary>

```go
package main

func main() {
	var obj1, obj2 struct {
		value int
	}

	// Bad swapping choice:
	tmp := obj1.value
	obj1.value = obj2.value
	obj2.value = tmp

	// This is the right way to do it:
	obj1.value, obj2.value = obj2.value, obj1.value
}
```

</details>

<hr>

**Notes**:

* Quickfix **can** replace a statement list with a single statement as we did in this example

**Next**: [Autofixable rules](autofixable-rules)
