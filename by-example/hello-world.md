# Hello world

`main.go`:
```go
package gorules

import "github.com/quasilyte/go-ruleguard/dsl/fluent"

func helloWorld(m fluent.Matcher) {
	m.Match(`fmt.Println("Hello, world")`).Report("found hello world")
}
```

`rules.go`:
```go
package main

import "fmt"

func main() {
	fmt.Println("Hello, world")
}
```

```bash
$ ruleguard -rules rules.go main.go
main.go:6:2: found hello world
```
