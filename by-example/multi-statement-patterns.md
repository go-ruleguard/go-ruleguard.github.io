# [Ruleguard by example](https://go-ruleguard.github.io/by-example/): Multi-statement patterns

```go
package gorules

import "github.com/quasilyte/go-ruleguard/dsl"

// Sometimes you want to match a statement that is **followed** or **preceded** by another statement.
// For this case we'll need a multi-statement patterns that gogrep does recognize.

func useParallelAssignment(m dsl.Matcher) {
	// If we want to specify $x statement that is followed by $y, we use `$x; $y` notation.
	// Here we match three statements:
	// 1. $x assigned to a temporary $tmp
	// 2. $x is re-assigned to $y
	// 3. $y is re-assigned to $tmp
	m.Match(`$tmp := $x; $x = $y; $y = $tmp`).
		Report(`use parallel assignent: $x, $y = $y, $x`)
}
```

<pre style="color: white; background-color: black">
$ ruleguard -c 0 -rules rules.go main.go
main.go:9:2: use parallel assignent: obj1.value, obj2.value = obj2.value, obj1.value
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

* `$*_` can be used to describe "any number of statements before this node": `$*_; $x`
* `$*_` can be used to describe "any number of statements after this node": `$x; $*_`
* Multi-statement patterns are especially useful when matching blocks: `{$*_; $x; $*_}`


<table><tr>
<td><a href="index">To index</a></td>
<td><a href="autofixable-rules">Next: Autofixable rules</a></td>
<td><a href="https://github.com/go-ruleguard/go-ruleguard.github.io/edit/master/by-example/multi-statement-patterns.md">Edit this page</a></td>
</tr></table>
