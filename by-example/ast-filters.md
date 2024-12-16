# [Ruleguard by example](https://go-ruleguard.github.io/by-example/): AST filters

```go
package gorules

import "github.com/quasilyte/go-ruleguard/dsl"

func goRecover(m dsl.Matcher) {
	// This `recover()` matching happens on AST (node) level,
	// not on a text level.
	m.Match(`go func () { $*body }`).
		Where(m["body"].Contains(`recover()`)).
		Report("called recover() inside go stmt")
}
```

<pre style="color: white; background-color: black">
$ ruleguard -c 0 -rules rules.go main.go
main.go:7:9: called recover() inside go stmt
7			rv := recover()
</pre>

<details><summary>main.go</summary>

```go
package main

import "fmt"

func main() {
	go func() {
		rv := recover()
		panic(rv)
	}
}
```

</details>

<hr>

**Notes**:

* Named submatches, like `$a` will refer to top-level `Match` matches

<table><tr>
<td><a href="index">To index</a></td>
<td><a href="constexpr-evaluation">Next: Constexpr evaluation</a></td>
<td><a href="https://github.com/go-ruleguard/go-ruleguard.github.io/edit/master/by-example/ast-filters.md">Edit this page</a></td>
</tr></table>
