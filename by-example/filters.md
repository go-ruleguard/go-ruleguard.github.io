# [Ruleguard by example](https://go-ruleguard.github.io/by-example/): Filters [:pencil:](https://github.com/go-ruleguard/go-ruleguard.github.io/blob/master/by-example/filters.md)

```go
package gorules

import "github.com/quasilyte/go-ruleguard/dsl/fluent"

// Suppose that we want to report the duplicated left and right operands of binary operations.
//
// But if the operand has some side effects, this rule can cause false positives:
// `f() && f()` can make sense (although it's not the best piece of code).
//
// This is where *filters* come to the rescue.

func dupSubExpr(m fluent.Matcher) {
	// All filters are written as a Where() argument.
	// In our case, we need to assert that $x is "pure".
	// It can be achieved by checking the m["x"] member Pure field.
	m.Match(`$x || $x`,
		`$x && $x`,
		`$x | $x`,
		`$x & $x`).
		Where(m["x"].Pure).
		Report(`suspicious identical LHS and RHS`)
}
```

<pre style="color: white; background-color: black">
$ ruleguard -c 0 -rules rules.go main.go
main.go:9:5: suspicious identical LHS and RHS
9		if cond && cond {
main.go:12:5: suspicious identical LHS and RHS
12		if obj.conditions[2] && obj.conditions[2] {
</pre>

<details><summary>main.go</summary>

```go
package main

func main() {
	var cond bool
	var obj struct {
		conditions []bool
	}

	if cond && cond {
	}

	if obj.conditions[2] && obj.conditions[2] {
	}

	if f() && f() { // $x is not pure
	}
}

func f() bool { return false }
```

</details>

<hr>

**Notes**:

* [`Where()`](https://pkg.go.dev/github.com/quasilyte/go-ruleguard/dsl/fluent#Matcher.Where) argument is just a normal Go boolean expression
* To combine several filter conditions, `&&` can be used, like in `m["x"].Pure && m["y"].Const`
* To get negated conditions, `!` can be used, like in `!m["x"].Const`
* If match is rejected due to the filter, the rest pattern alternatives from [`Match()`](https://pkg.go.dev/github.com/quasilyte/go-ruleguard/dsl/fluent#Matcher.Match) (if any) will not be checked

**Next**: TODO
