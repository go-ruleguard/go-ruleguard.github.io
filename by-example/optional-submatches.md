# [Ruleguard by example](https://go-ruleguard.github.io/by-example/): Optional submatches

```go
package gorules

import "github.com/quasilyte/go-ruleguard/dsl"

// The special `*` modifier can be used to match 0 or more nodes.

func nilValReturn(m dsl.Matcher) {
	// $*_ is used to make init clause inside if statement optional,
	// so our pattern can match all if statements.
	m.Match(`if $*_; $x == nil { return $x }`).
		Report(`returned $x value is always nil`)
}

func fprintStdout(m dsl.Matcher) {
	// `*` can be used with a name other than "_".
	// This is especially useful for variadic functions matching.
	// The named $* submatch can be used in both Report() and Suggest() templates.
	m.Match(`fmt.Fprint(os.Stdout, $*args)`).
		Suggest(`fmt.Print($args)`)
}
```

<pre style="color: white; background-color: black">
$ ruleguard -c 1 -rules rules.go main.go
main.go:9:2: suggestion: fmt.Print("Hello", "world!")
8	func main() {
9		fmt.Fprint(os.Stdout, "Hello", "world!")
10		foo()
main.go:15:2: returned err value is always nil
14		err := os.Chdir("/example1")
15		if err == nil { // No init clause
16			return err
main.go:18:2: returned err value is always nil
17		}
18		if err := os.Chdir("/example2"); err == nil {
19			return err
</pre>

<details><summary>main.go</summary>

```go
package main

import (
	"fmt"
	"os"
)

func main() {
	fmt.Fprint(os.Stdout, "Hello", "world!")
	foo()
}

func foo() error {
	err := os.Chdir("/example1")
	if err == nil { // No init clause
		return err
	}
	if err := os.Chdir("/example2"); err == nil {
		return err
	}
	if err == nil {
		return nil // Doesn't match
	}
	return nil
}
```

</details>

<hr>

**Notes**:

* `-c 1` argument makes the `ruleguard` print **three** context lines in its output
* A single `gorules` file can contain multiple rule groups; today we defined `dupArg` and `nilValReturn` there
* When using a `$*` variable in message templates, they're spelled without `*` (e.g. `$args`, not `$*args`)

<table><tr>
<td><a href="index">To index</a></td>
<td><a href="filters">Next: Filters</a></td>
<td><a href="https://github.com/go-ruleguard/go-ruleguard.github.io/edit/master/by-example/optional-submatches.md">Edit this page</a></td>
</tr></table>
