# [Ruleguard by example](https://go-ruleguard.github.io/by-example/): Text filters

```go
package gorules

import "github.com/quasilyte/go-ruleguard/dsl"

// Normally, all matching is done on the AST and types level.
// But sometimes it's necessary (or just easier) to treat the
// submatches as text.

func printFmt(m dsl.Matcher) {
	// Here we scan a $s string for %s, %d and %v.
	m.Match(`fmt.Println($s, $*_)`,
		`fmt.Print($s, $*_)`).
		Where(m["s"].Text.Matches(`%[sdv]`)).
		Report("found formatting directive in non-formatting call")
}

func wrongErrChecked(m dsl.Matcher) {
	// Here we check for both types and names.
	// Note that we can use `==` and `!=` with Text.
	m.Match("if $*_, $err0 := $*_; $err1 != nil { $*_ }").
		Where(m["err0"].Text == "err" && m["err0"].Type.Is("error") &&
			m["err1"].Text != "err" && m["err1"].Type.Is("error")).
		Report("maybe wrong err in error check")
}
```

<pre style="color: white; background-color: black">
$ ruleguard -c 0 -rules rules.go main.go
main.go:8:2: found formatting directive in non-formatting call
8		fmt.Println("error: %s", msg)
main.go:9:2: found formatting directive in non-formatting call
9		fmt.Print("Hello, %s!", "ruleguard")
main.go:14:2: maybe wrong err in error check
14		if v, err := f(); err2 != nil {
</pre>

<details><summary>main.go</summary>

```go
package main

import "fmt"

func main() {
	var msg string

	fmt.Println("error: %s", msg)
	fmt.Print("Hello, %s!", "ruleguard")
	fmt.Println("no formatting directives")

	var err2 error

	if v, err := f(); err2 != nil {
		fmt.Println(v, err)
	}
	if v, err2 := f(); err2 != nil {
		fmt.Println(v, err2)
	}
	if v, err := f(); err != nil {
		fmt.Println(v, err)
	}
}

func f() (int, error) { return 0, nil }
```

</details>

<hr>

**Notes**:

* If ruleguard lacks some feature and you're solving this by using text filters, consider doing a [feature request](https://github.com/quasilyte/go-ruleguard/issues/new)
* You can use `Text` on a right-hand-side (e.g. `m["x"].Text != m["y"].Text`)

<table><tr>
<td><a href="index">To index</a></td>
<td><a href="ast-filters">Next: AST filters</a></td>
<td><a href="https://github.com/go-ruleguard/go-ruleguard.github.io/edit/master/by-example/text-filters.md">Edit this page</a></td>
</tr></table>
