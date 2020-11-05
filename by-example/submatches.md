# [Ruleguard by example](https://go-ruleguard.github.io/by-example/): Submatches

```go
package gorules

import "github.com/quasilyte/go-ruleguard/dsl/fluent"

func dupArg(m fluent.Matcher) {
	// The pattern string can contain *submatch variables*.
	//
	// $_ is unnamed submatch variable, it matches any kind of AST element.
	// $<name> (e.g. $x) is a named submatch variable; it also matches anything.
	// Repeated named variable requires all submatches to be identical, like in pattern matching.
	//
	// If named submatch is used in Report() argument string, it will be interpolated.
	m.Match(`strings.ReplaceAll($_, $x, $x)`).
		Report(`suspicious duplicated arg $x`)
}
```

<pre style="color: white; background-color: black">
$ ruleguard -c 0 -rules rules.go main.go
main.go:9:10: suspicious duplicated arg "_"
9		println(strings.ReplaceAll(s, "_", "_")) // $x="_"
main.go:11:10: suspicious duplicated arg part
11		println(strings.ReplaceAll(path+"/local/bin/", part, part)) // $x=part
</pre>

<details><summary>main.go</summary>

```go
package main

import "strings"

var path string

func main() {
	s := "Hello, world"
	println(strings.ReplaceAll(s, "_", "_")) // $x="_"
	part := "x"
	println(strings.ReplaceAll(path+"/local/bin/", part, part)) // $x=part

	println(strings.ReplaceAll(s, "_", "")) // Doesn't match
}
```

</details>

<hr>

**Notes**:

* `$_` can be repeated many times, it doesn't require submatches to be identical
* Variable names should be helpful, like `$len`; don't use `$x` everywhere

**Next**: TODO
