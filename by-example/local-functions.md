# [Ruleguard by example](https://go-ruleguard.github.io/by-example/): Local functions

```go
package gorules

import "github.com/quasilyte/go-ruleguard/dsl"

// Sometimes we find ourselves repeating the parts of Where()
// clause in several rules.
// Other times we use some conditions that may be
// more clear if we would give them a better name.
// In normal Go, we use functions for that.

func exposedMutex(m dsl.Matcher) {
	// isExported reports whether bound matcher var
	// represents exported Go identifier.
	isExported := func(v dsl.Var) bool {
		return v.Text.Matches(`^\p{Lu}`)
	}

	// isExported() is more readable than `^\p{Lu}`
	// regular expression itself.

	m.Match(`type $name struct { $*_; sync.Mutex; $*_ }`).
		Where(isExported(m["name"])).
		Report("do not embed sync.Mutex")

	m.Match(`type $name struct { $*_; sync.RWMutex; $*_ }`).
		Where(isExported(m["name"])).
		Report("don not embed sync.RWMutex")
}
```

<pre style="color: white; background-color: black">
$ ruleguard -rules rules.go main.go
main.go:5:1: exposedMutex: do not embed sync.Mutex (rules.go:21)
</pre>

<details><summary>main.go</summary>

```go
package main

import "sync"

type EmbedsMutex struct {
	key int
	sync.Mutex
}

type unexportedEmbedsMutex struct {
	sync.Mutex
}

func main() {}
```

</details>

<hr>

**Notes**:

* Local predicate functions can have any params, but they should return `bool`
* You can use captured variables inside local functions

<table><tr>
<td><a href="index">To index</a></td>
<td><a href="?">Next: ?</a></td>
<td><a href="https://github.com/go-ruleguard/go-ruleguard.github.io/edit/master/by-example/local-functions.md">Edit this page</a></td>
</tr></table>
