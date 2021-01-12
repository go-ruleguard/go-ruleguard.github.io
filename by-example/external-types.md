# [Ruleguard by example](https://go-ruleguard.github.io/by-example/): External types

```go
package gorules

import "github.com/quasilyte/go-ruleguard/dsl"

// Imagine that you're using the sqlx package.
// You have runQuery() and runQueryContext() wrappers to exec SQL queries.
//
// What you would like to do is to suggest runQueryContext() whether it's possible.
//
// Now we need to check if the runQuery() argument is capable of ExecContext() and
// if it is, propose the runQueryContext() as an alternative.

func execContext(m dsl.Matcher) {
	// First, we need to tell ruleguard what do we mean when we say "sqlx" inside
	// the type filters. Import() works for the current group scope.
	m.Import("github.com/jmoiron/sqlx")

	// Now it's as simple as it was before: check whether $e implements the
	// interface from the imported package.
	m.Match(`runQuery($e, $*_)`).
		Where(m["e"].Type.Implements(`sqlx.ExecerContext`)).
		Report(`prefer runQueryContext`)
}
```

<pre style="color: white; background-color: black">
$ ruleguard -c 0 -rules rules.go main.go
main.go:13:2: prefer runQueryContext
13		runQuery(&db, "SELECT 9000")
</pre>

<details><summary>main.go</summary>

```go
package main

import (
	"context"

	"github.com/jmoiron/sqlx"
)

func main() {
	var db sqlx.DB
	var execer sqlx.Execer // Doesn't implement sqlx.ExecerContext

	runQuery(&db, "SELECT 9000")
	runQuery(execer, "SELECT 9000")
}

func runQuery(e sqlx.Execer, rest ...interface{}) {}

func runQueryContext(ctx context.Context, e sqlx.ExecerContext, rest ...interface{}) {}
```

</details>

<hr>

**Notes**:

* [`Import()`](https://pkg.go.dev/github.com/quasilyte/go-ruleguard/dsl#Matcher.Import) must successfully find the specified packages during the ruleguard init phase
* Packages with the last package path part that is different from their package name are not supported yet

<table><tr>
<td><a href="index">To index</a></td>
<td><a href="underlying-types">Next: Underlying types</a></td>
<td><a href="https://github.com/go-ruleguard/go-ruleguard.github.io/edit/master/by-example/external-types.md">Edit this page</a></td>
</tr></table>
