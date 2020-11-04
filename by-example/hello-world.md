# Hello world

```go
package gorules

import "github.com/quasilyte/go-ruleguard/dsl/fluent"

func helloWorld(m fluent.Matcher) {
	m.Match(`fmt.Println("Hello, world")`)
}
```

```go
package main

import "fmt"

func main() {
	fmt.Println("Hello, world")
}
```
