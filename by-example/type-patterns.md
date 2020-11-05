# [Ruleguard by example](https://go-ruleguard.github.io/by-example/): Type patterns

```go
package gorules

import "github.com/quasilyte/go-ruleguard/dsl/fluent"

// Sometimes you can't describe the exact type using the normal syntax.
//
// Suppose that you're trying to find array-typed value dereferencing.
// There can be unlimited ways to write array type: not only you need to spell
// its element type, it should also have a size specified as a part of the type.
//
// Fear not, we have a solution!

func unslice(m fluent.Matcher) {
	// We can safely omit the [:] for strings and slices (but not arrays!).
	// The `[]$_` type pattern accepts any slice types: []int, []T, [][]string and so on.
	m.Match(`$s[:]`).
		Where(m["s"].Type.Is(`[]$_`)).
		Suggest(`$s`).
		Report(`$s[:] is a no-op, can be written as $s`)
}

func underef(m fluent.Matcher) {
	// Arrays are automatically dereferenced during the indexing, so explicit * is not needed.
	// The `*[$_]$_` type pattern accepts any pointer-to-array types: *[5]int, [16][]T and so on.
	m.Match(`(*$arr)[$i]`).
		Where(m["arr"].Type.Is(`*[$_]$_`)).
		Suggest(`$arr[$i]`).
		Report(`explicit dereference is redundant, can be written as $arr[$i]`)
}
```

<pre style="color: white; background-color: black">
$ ruleguard -c 0 -rules rules.go main.go
main.go:10:10: b[:] is a no-op, can be written as b
10		println(b[:])
main.go:16:10: explicit dereference is redundant, can be written as intArr5ptr[0]
16		println((*intArr5ptr)[0])
main.go:17:10: explicit dereference is redundant, can be written as stringArr2Ptr[0]
17		println((*stringArr2Ptr)[0])
</pre>

<details><summary>main.go</summary>

```go
package main

type myStringer struct{}

func (myStringer) String() string { return "example" }

func main() {
	var b []byte

	println(b[:])

	var intArr5 [5]int
	intArr5ptr := &intArr5
	stringArr2Ptr := new([2]string)

	println((*intArr5ptr)[0])
	println((*stringArr2Ptr)[0])
}
```

</details>

<hr>

**Notes**:

* It's possible to use named variables inside type patterns; repeated names match identical types
* `map[$k]$v` matches any map
* `map[$t]$t` matches a map where a key and value types are the same
* `struct{$*_}` matches any struct
* `struct{$x; $*_}` matches a struct that has $x-typed first field
* `struct{$*_; $x; $*_}` matches a struct that contains $x-typed field
* `struct{$*_; $x}` matches a struct that has $x-typed last field
* It's impossible to enumerate everything here, but you should get the general

<table><tr>
<td><a href="index">To index</a></td>
<td><a href="external-types">Next: External types</a></td>
<td><a href="https://github.com/go-ruleguard/go-ruleguard.github.io/edit/master/by-example/type-patterns.md">Edit this page</a></td>
</tr></table>
