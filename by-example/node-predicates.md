# [Ruleguard by example](https://go-ruleguard.github.io/by-example/): Node predicates

```go
package gorules

import "github.com/quasilyte/go-ruleguard/dsl"

// When ruleguard inspects the program and tries to match your rules,
// it walks over the AST created from the target files source code.
// The AST consists of the nodes that are defined in "go/ast" package.
//
// These node types are not connected to the "expression types" directly,
// but they're useful when you want to assert the shape of the match.

func formatLiteral(m dsl.Matcher) {
	// The following rule tries to find the formatting calls
	// that use non-literal arguments for the formatting string.
	// So, variables, named constants and everything else is a no-no.
	m.Match(`fmt.Sprintf($format, $*_)`,
		`fmt.Fprintf($_, $format, $*_)`,
		`fmt.Printf($format, $*_)`).
		Where(!m["format"].Node.Is(`BasicLit`)).
		Report("non-literal formatting string is used")
}
```

<pre style="color: white; background-color: black">
$ ruleguard -c 0 -rules rules.go main.go
main.go:9:6: non-literal formatting string is used
9		_ = fmt.Sprintf(formatStringVar, 1)
main.go:10:6: non-literal formatting string is used
10		_ = fmt.Sprintf(formatStringConst, 2)
main.go:13:2: non-literal formatting string is used
13		fmt.Printf("%"+"d", -5)
</pre>

<details><summary>main.go</summary>

```go
package main

import "fmt"

func main() {
	formatStringVar := "%d"
	const formatStringConst = "%d"

	_ = fmt.Sprintf(formatStringVar, 1)
	_ = fmt.Sprintf(formatStringConst, 2)
	_ = fmt.Sprintf("%s: %d", "the answer is", 42) // OK

	fmt.Printf("%"+"d", -5)

	fmt.Printf("%d", 10) // OK
}
```

</details>

<hr>

**Notes**:

* [`Is()`](https://pkg.go.dev/github.com/quasilyte/go-ruleguard/dsl#MatchedNode.Is) argument is a [`go/ast`](https://golang.org/pkg/go/ast/) type name
* You can use [`m["$$"].Parent()`](https://pkg.go.dev/github.com/quasilyte/go-ruleguard/dsl#MatchedNode.Parent) to add filters based on the AST context

<table><tr>
<td><a href="index">To index</a></td>
<td><a href="file-predicates">Next: File predicates</a></td>
<td><a href="https://github.com/go-ruleguard/go-ruleguard.github.io/edit/master/by-example/node-predicates.md">Edit this page</a></td>
</tr></table>
