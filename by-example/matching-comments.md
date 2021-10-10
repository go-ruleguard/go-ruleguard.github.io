# [Ruleguard by example](https://go-ruleguard.github.io/by-example/): Matching comments

```go
package gorules

import (
	"github.com/quasilyte/go-ruleguard/dsl"
)

// Match() is great for matching AST nodes; comments are
// usually ignored when we speak about the structural matching.
// But what if you want to check the comments?
//
// MatchComment() method is designed just for that.
// You specify regexp patterns that are applied to all comment nodes.

func unknownGoDirective(m dsl.Matcher) {
	// Named capture groups populate the m map.
	// Here, we use a "tag" as a capture name and use m["tag"]
	// to add some filters. Since it's a normal match var,
	// it can be interpolated into Report() or Suggest() message.
	m.MatchComment(`//go:(?P<tag>\w+)`).
		Where(!m["tag"].Text.Matches(`noescape|noinline|nosplit|norace`)).
		Report(`using unknown go directive "$tag"`)
}
```

<pre style="color: white; background-color: black">
$ ruleguard -rules rules.go main.go 
main.go:3:1: unknownGoDirective: using unknown go directive "inline"
main.go:10:2: unknownGoDirective: using unknown go directive "nolint"
</pre>

<details><summary>main.go</summary>

```go
package main

//go:inline
func f() int { return 10 }

//go:noinline
func g() int { return 0 }

func main() {
	//go:nolint unused-call-result
	f()
}
```

</details>

<hr>

**Notes**:

* You can use all usual things, like `At()` and other match var related methods

<table><tr>
<td><a href="index">To index</a></td>
<td><a href="go-version-checks">Next: Go version checks</a></td>
<td><a href="https://github.com/go-ruleguard/go-ruleguard.github.io/edit/master/by-example/matching-comments.md">Edit this page</a></td>
</tr></table>
