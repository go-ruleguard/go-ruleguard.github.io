# [Ruleguard by example](https://go-ruleguard.github.io/by-example/): Autofixable rules

```go
package gorules

import "github.com/quasilyte/go-ruleguard/dsl/fluent"

func stringsSimplify(m fluent.Matcher) {
	// Some issues have simple fixes that can be expressed as
	// a replacement pattern. Rules can use Suggest() function
	// to add a quickfix action for such issues.
	m.Match(`strings.Replace($s, $old, $new, -1)`).
		Report(`this Replace() call can be simplified`).
		Suggest(`strings.ReplaceAll($s, $old, $new)`)

	// Suggest() can be used without Report().
	// It'll print the suggested template to the user.
	m.Match(`strings.Count($s1, $s2) == 0`).
		Suggest(`!strings.Contains($s1, $s2)`)
}
```

<pre style="color: white; background-color: black">
$ ruleguard -c 0 -rules rules.go main.go
main.go:8:6: this Replace() call can be simplified
8		_ = strings.Replace(s, "old", "new", -1)
main.go:11:5: suggestion: !strings.Contains(s, "#golang")
11		if strings.Count(s, "#golang") == 0 {
</pre>

If we'll run `ruleguard` with the `-fix` flag, quickfix actions will be applied inplace.

<pre style="color: white; background-color: black">
$ ruleguard -fix -c 0 -rules rules.go main.go
</pre>

```diff
+ strings.ReplaceAll(s, "old", "new")
- strings.Replace(s, "old", "new", -1)

+ if !strings.Contains(s, "#golang") {
- if strings.Count(s, "#golang") == 0 {
```

<details><summary>main.go</summary>

```go
package main

import "strings"

func main() {
	var s string

	_ = strings.Replace(s, "old", "new", -1)
	_ = strings.Replace(s, "old", "new", 10) // Not matched

	if strings.Count(s, "#golang") == 0 {
		println("no golang tags?")
	}
}
```

</details>

<hr>

**Notes**:

* Combine [`Report()`](https://pkg.go.dev/github.com/quasilyte/go-ruleguard/dsl/fluent#Matcher.Report) and [`Suggest()`](https://pkg.go.dev/github.com/quasilyte/go-ruleguard/dsl/fluent#Matcher.Suggest) if suggested replacement is not self-explanatory

**Next**: TODO
