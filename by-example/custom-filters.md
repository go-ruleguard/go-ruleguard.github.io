# [Ruleguard by example](https://go-ruleguard.github.io/by-example/): Custom filters

```go
package gorules

import (
	"github.com/quasilyte/go-ruleguard/dsl"
	"github.com/quasilyte/go-ruleguard/dsl/types"
)

// If you need to apply a filter that can't be expressed via Matcher API,
// it's possible to use a lower-level mechanism and resort to `go/types`.

// implementsStringer is a custom filter function.
// It tries both T and *T to see whether a type implements `fmt.Stringer`.
func implementsStringer(ctx *dsl.VarFilterContext) bool {
    // ctx.Type mapped to m["x"].Type;
	// ctx.GetInterface finds the `types.Type` object for the `fmt.Stringer`.
	stringer := ctx.GetInterface(`fmt.Stringer`)
	return types.Implements(ctx.Type, stringer) ||
		types.Implements(types.NewPointer(ctx.Type), stringer)
}

func sprintStringer(m dsl.Matcher) {
	// If we used m["x"].Type.Implements(`fmt.Stringer`), then we
	// wouldn't get all results we wanted: if $x type implements
	// stringer by pointer, then non-pointer value will not "implement" it.
	// Our custom filter will try both pointer and non-pointer versions.
	m.Match(`fmt.Sprint($x)`).
		Where(m["x"].Filter(implementsStringer) && m["x"].Addressable).
		Report(`can use $x.String() directly`)
}
```

<pre style="color: white; background-color: black">
$ ruleguard -rules rules.go main.go
main.go:9:10: can use foo.String() directly
main.go:10:10: can use fooPtr.String() directly
</pre>

<details><summary>main.go</summary>

```go
package main

import "fmt"

func main() {
	fooPtr := &Foo{}
	foo := Foo{}

	println(fmt.Sprint(foo))
	println(fmt.Sprint(fooPtr))

	println(fmt.Sprint(0))    // Not fmt.Stringer
	println(fmt.Sprint(&foo)) // Not addressable
}

type Foo struct{}

func (*Foo) String() string { return "Foo" }
```

</details>

<hr>

**Notes**:

* Custom filter functions are bytecode-compiled when rule files are parsed
* If you get a function compilation error, try using a simpler way to express it
* `github.com/quasilyte/go-ruleguard/dsl/types` mimics `go/types` package
* `github.com/quasilyte/go-ruleguard/dsl/ast` mimics `go/ast` package

<table><tr>
<td><a href="index">To index</a></td>
<td><a href="?">Next: ?</a></td>
<td><a href="https://github.com/go-ruleguard/go-ruleguard.github.io/edit/master/by-example/custom-filters.md">Edit this page</a></td>
</tr></table>
