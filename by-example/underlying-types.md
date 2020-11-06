# [Ruleguard by example](https://go-ruleguard.github.io/by-example/): Underlying types

```go
package gorules

import "github.com/quasilyte/go-ruleguard/dsl/fluent"

// The binary.Write function from the "encoding/binary" package
// does a simple Write call if the data argument is []byte.
//
// In that case we may want to suggest calling the Write method
// directly, without extra binary.Write overhead.
//
// But what if the data is not exactly []byte, but some named
// type that is defined like this: `type myBytes []byte`.
// It's still bytes, right?

func replaceAll(m fluent.Matcher) {
	// When we consider myBytes type we need to remember that
	// it's *underlying* type is []byte. The underlying type
	// of []byte is []byte.
	//
	// This leads us to the Underlying() method of the ExprType.

	m.Match(`binary.Write($w, $_, $v)`).
		Where(m["v"].Type.Underlying().Is(`[]byte`)).
		Report(`consider doing $w.Write($v) instead`)
}
```

<pre style="color: white; background-color: black">
$ ruleguard -c 0 -rules rules.go main.go
main.go:19:2: consider doing buf.Write(b) instead
19		binary.Write(buf, binary.LittleEndian, b)
main.go:22:2: consider doing buf.Write(customBytes) instead
22		binary.Write(buf, binary.LittleEndian, customBytes)
main.go:23:2: consider doing buf.Write(myByteSlice2("123")) instead
23		binary.Write(buf, binary.LittleEndian, myByteSlice2("123"))
</pre>

<details><summary>main.go</summary>

```go
package main

import (
	"bytes"
	"encoding/binary"
)

// myByteSlice underlying type is []byte
type myByteSlice []byte

// myByteSlice2 underlying type is still []byte.
type myByteSlice2 myByteSlice

func main() {
	var b []byte
	var customBytes myByteSlice
	buf := &bytes.Buffer{}

	binary.Write(buf, binary.LittleEndian, b)

	// These 2 lines will not give any warnings if we don't use Underlying().
	binary.Write(buf, binary.LittleEndian, customBytes)
	binary.Write(buf, binary.LittleEndian, myByteSlice2("123"))

	binary.Write(buf, binary.LittleEndian, 14) // OK: not []byte
}
```

</details>

<hr>

**Notes**:

* Underlying type is never a **named type** (search the [Go spec](https://golang.org/ref/spec#Types) for "underlying type")

<table><tr>
<td><a href="index">To index</a></td>
<td><a href="TODO">Next: TODO</a></td>
<td><a href="https://github.com/go-ruleguard/go-ruleguard.github.io/edit/master/by-example/underlying-types.md">Edit this page</a></td>
</tr></table>
