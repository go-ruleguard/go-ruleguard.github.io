# [Ruleguard by example](https://go-ruleguard.github.io/by-example/): Deadcode detection

```go
package gorules

import "github.com/quasilyte/go-ruleguard/dsl"

// Some checks may give false positives if they're
// executed in the context of the unreachable code.
// For example, you may want to implement a rule
// that checks for shift operation overflow.
// It can give false positives for a code that does
// if-based branching based on the arch-dependent things;
// one branch is always true, another one is always false.
// The always false branch should generally be ignored
// by such rule.

func shiftOverflow(m dsl.Matcher) {
	// Deadcode() predicate returns true if the matched
	// pattern is located on the unreachable code path.
	// The negation of that predicate reports reachable
	// code paths.
	m.Match(`$x << $n`).
		Where(!m["x"].Const &&
			m["x"].Type.Size == 1 &&
			m["n"].Value.Int() >= 8 &&
			!m.Deadcode()).
		Report(`$x (8 bits) too small for shift of $n`)
}
```

<pre style="color: white; background-color: black">
$ ruleguard -c 0 -rules rules.go main.go
</pre>

<details><summary>main.go</summary>

```go
package main

import "unsafe"

func main() {
	var x int
	var result int
	if unsafe.Sizeof(int(0)) == 4 {
		// 32-bit platform.
		result = x << 31
	} else {
		// 64-bit platform.
		// This would give a dull warning when
		// targeting a 32-bit platform
		result = x << 63
	}
	println(result)
}
```

</details>

<hr>

**Notes**:

* Deadcode detection is very simple, it only works for compile-time solveable conditions
* Note: there are no warnings for this example in both `32-bit` and `64-bit` modes

<table><tr>
<td><a href="index">To index</a></td>
<td><a href="local-functions">Next: Local functions</a></td>
<td><a href="https://github.com/go-ruleguard/go-ruleguard.github.io/edit/master/by-example/deadcode-detection.md">Edit this page</a></td>
</tr></table>
