# [Ruleguard by example](https://go-ruleguard.github.io/by-example/): Custom locations

```go
package gorules

import "github.com/quasilyte/go-ruleguard/dsl/fluent"

// When analyzer is integrated into the text editor (or IDE),
// it's very important to provide good issue locations so
// the editor can highlight the offending part correctly.

func mismatchingUnlock(m fluent.Matcher) {
	// By default, an entire match position is used as a location.
	// This can be changed by the At() method that binds the location
	// to the provided named submatch.
	//
	// In the rules below text editor would get mismatching method
	// name locations:
	//
	//   defer mu.RUnlock()
	//            ^^^^^^^

	m.Match(`$mu.Lock(); defer $mu.$unlock()`).
		Where(m["unlock"].Text == "RUnlock").
		At(m["unlock"]).
		Report(`maybe $mu.Unlock() was intended?`)

	m.Match(`$mu.RLock(); defer $mu.$unlock()`).
		Where(m["unlock"].Text == "Unlock").
		At(m["unlock"]).
		Report(`maybe $mu.RUnlock() was intended?`)
}
```

<pre style="color: white; background-color: black">
$ ruleguard -c 0 -rules rules.go main.go
main.go:14:11: maybe mu.Unlock() was intended?
14		defer mu.RUnlock()
main.go:19:11: maybe mu.RUnlock() was intended?
19		defer mu.Unlock()
</pre>

If we'll run `ruleguard` with the `-fix` flag, quickfix actions will be applied inplace.

<pre style="color: white; background-color: black">
$ ruleguard -fix -c 0 -rules rules.go main.go
</pre>

<details><summary>main.go</summary>

```go
package main

import "sync"

func main() {
	var mu sync.RWMutex
	f1(&mu)
	f2(&mu)
	f3(&mu)
}

func f1(mu *sync.RWMutex) {
	mu.Lock()
	defer mu.RUnlock()
}

func f2(mu *sync.RWMutex) {
	mu.RLock()
	defer mu.Unlock()
}

func f3(mu *sync.RWMutex) {
	mu.Lock()
	defer mu.Unlock()
}
```

</details>

<hr>

**Notes**:

* Custom locations are especially important for multi-statement rules

<table><tr>
<td><a href="index">To index</a></td>
<td><a href="?">Next: ?</a></td>
<td><a href="https://github.com/go-ruleguard/go-ruleguard.github.io/edit/master/by-example/custom-locations.md">Edit this page</a></td>
</tr></table>
