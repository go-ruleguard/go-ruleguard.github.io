# [Ruleguard by example](https://go-ruleguard.github.io/by-example/): Hello world

```go
package gorules

// We always need this import.
import "github.com/quasilyte/go-ruleguard/dsl/fluent"

// helloWorld is a *named rules group*.
// It has only 1 simple rule inside and its name is "helloWorld".
// The *rules group* signature is fixed.
func helloWorld(m fluent.Matcher) {
	// A chain that starts with Match() and ends with Report() (or Suggest) call called a *rule*.
	// Therefore, a minimal *rule* consists of Match+Report call.
	//
	// Match matches the AST using the gogrep pattern string;
	// Report prints the specified message when this rule is matched.
	m.Match(`fmt.Println("Hello, world")`).Report("found hello world")
}
```

<details><summary>main.go</summary>

```go
package main

import "fmt"

func main() {
	fmt.Println("Hello, world")
}
```

</details>

<pre style="color: white; background-color: black">
$ ruleguard -c 0 -rules rules.go main.go
main.go:6:2: found hello world
6		fmt.Println("Hello, world")
</pre>

**Notes**:

* `-c 0` argument makes the `ruleguard` print **one** context line in it's output (line 6 in this example)
* `-rules` argument accepts a `ruleguard` rules source (always `rules.go` in out examples)
* [`Match()`](https://pkg.go.dev/github.com/quasilyte/go-ruleguard/dsl/fluent#Matcher.Match) argument is a [github.com/mvdan/gogrep](https://github.com/mvdan/gogrep) pattern string
* Ruleguard rules file **must** start with `package gorules`; it can't be other package
* The [`dsl/fluent`](https://pkg.go.dev/github.com/quasilyte/go-ruleguard/dsl/fluent) package documentation is a common source of the most answers

**Next**: [Submatches](submatches)
