# [Ruleguard by example](https://go-ruleguard.github.io/by-example/): Multi-rule groups

```go
package gorules

import "github.com/quasilyte/go-ruleguard/dsl"

func assignOp(m dsl.Matcher) {
	// This group defines 2 rules.
	//
	// Rules are executed in the order they're described, until
	// the first successfull match.
	//
	// This short-circuit logic allows us to write overlapping rules
	// without conflicts: only one of them will be executed.
	m.Match(`$x = $x + 1`).Report(`can rewrite as $x++`)
	m.Match(`$x = $x + $y`).Report(`can rewrite as $x += $y`)
}
```

<pre style="color: white; background-color: black">
$ ruleguard -c 0 -rules rules.go main.go
main.go:7:2: can rewrite as v1 += v2
7		v1 = v1 + v2
main.go:8:2: can rewrite as v1++
8		v1 = v1 + 1
</pre>

<details><summary>main.go</summary>

```go
package main

func main() {
    v1 := 0
    v2 := 10

    v1 = v1 + v2
    v1 = v1 + 1
}
```

</details>

<hr>

**Notes**:

* If you don't want rules to short-circuit, place them in separate functions

<table><tr>
<td><a href="index">To index</a></td>
<td><a href="multi-pattern-rules">Next: Multi-pattern rules</a></td>
<td><a href="https://github.com/go-ruleguard/go-ruleguard.github.io/edit/master/by-example/multi-rule-groups.md">Edit this page</a></td>
</tr></table>
