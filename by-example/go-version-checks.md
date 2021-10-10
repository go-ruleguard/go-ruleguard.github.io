# [Ruleguard by example](https://go-ruleguard.github.io/by-example/): Go version checks

```go
package gorules

import (
	"github.com/quasilyte/go-ruleguard/dsl"
)

// Some rules require particular Go features to be available.
// For example, `ioutil` package was deprecated in 1.16.
// We can't suggest `ioutil` alternatives unless the target
// Go version is at least 1.16. It would be annoying to
// get the warning that is not actionable for you. 

func ioutilDeprecated(m dsl.Matcher) {
	// GreaterEqThan performs a "version >= x" check.
	m.Match(`ioutil.ReadAll($r)`).
		Where(m.GoVersion().GreaterEqThan("1.16")).
		Suggest(`io.ReadAll($r)`).
		Report(`ioutil.ReadAll is deprecated, use io.ReadAll instead`)
}
```

<pre style="color: white; background-color: black">
$ ruleguard -rules rules.go main.go 
main.go:9:1: ioutilDeprecated: ioutil.ReadAll is deprecated, use io.ReadAll instead
</pre>

<details><summary>main.go</summary>

```go
package main

import (
	"io"
	"io/ioutil"
)

func example(r io.Reader) {
	ioutil.ReadAll(r)
}
```

</details>

<hr>

**Notes**:

* The Go version is configured explicitly by the user
* There are more version operations, so you can perform `==`, `<`, `>` and so on
* Use `&&` in combination with [`GoVersion`](https://pkg.go.dev/github.com/quasilyte/go-ruleguard/dsl#GoVersion) methods to create a version range constraint

<table><tr>
<td><a href="index">To index</a></td>
<td><a href="go-version-checks">Next: ?</a></td>
<td><a href="https://github.com/go-ruleguard/go-ruleguard.github.io/edit/master/by-example/go-version-checks.md">Edit this page</a></td>
</tr></table>
