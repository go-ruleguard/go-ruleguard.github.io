# [Ruleguard by example](https://go-ruleguard.github.io/by-example/): Type filters

```go
package gorules

import "github.com/quasilyte/go-ruleguard/dsl"

// Type filters make it possible to reject matches based on the submatch expr types.

func copyStringToBytes(m dsl.Matcher) {
	// It's not necessary to do a string->[]byte conversion for the copy() argument
	// when copying a string into []byte. In this rule, we require that $s is string-typed.
	m.Match(`copy($b, []byte($s))`).
		Where(m["s"].Type.Is(`string`)).
		Suggest(`copy($b, $s)`)
}

func sprintStringer(m dsl.Matcher) {
	// Instead of calling fmt.Sprint on a value that implements fmt.Stringer,
	// we can call its String() method directly. In order to report such cases,
	// it's necessary to assert that $x is something that implements that interface.
	m.Match(`fmt.Sprint($x)`).
		Where(m["x"].Type.Implements(`fmt.Stringer`)).
		Suggest(`$x.String()`)
}
```

<pre style="color: white; background-color: black">
$ ruleguard -c 0 -rules rules.go main.go
main.go:14:2: suggestion: copy(b, s)
14		copy(b, []byte(s)) // Excessive conversion
main.go:17:10: suggestion: myStringer{}.String()
17		println(fmt.Sprint(myStringer{}))
main.go:18:10: suggestion: stringer.String()
18		println(fmt.Sprint(stringer))
</pre>

<details><summary>main.go</summary>

```go
package main

import "fmt"

type myStringer struct{}
func (myStringer) String() string { return "example" }

func main() {
	var s string
	var b []byte
	var stringer myStringer

	copy(b, []byte(s)) // Excessive conversion
	copy(b, s)         // Neat, no excessive conversion here

	println(fmt.Sprint(myStringer{}))
	println(fmt.Sprint(stringer))
	println(stringer.String()) // Good
}
```

</details>

<hr>

**Notes**:

* It's possible to use third-party (external) types in filters, but they need to be imported; we'll get to it
* By default, you can use builtin types (e.g. `int16`) and qualified types from the stdlib (e.g. `io.Reader`)

<table><tr>
<td><a href="index">To index</a></td>
<td><a href="type-patterns">Next: Type patterns</a></td>
<td><a href="https://github.com/go-ruleguard/go-ruleguard.github.io/edit/master/by-example/type-filters.md">Edit this page</a></td>
</tr></table>
