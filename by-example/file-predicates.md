# [Ruleguard by example](https://go-ruleguard.github.io/by-example/): File predicates

```go
package gorules

import "github.com/quasilyte/go-ruleguard/dsl"

// The "path/filepath" package forwards path separators so if
// the file already uses filepath-related API it might be
// a good idea to reduce the direct "os" package dependency.
//
// In some cases it helps to remove the "os" package import completely.

func osFilepath(m dsl.Matcher) {
	// The File() method returns an object that can be used
	// to add file-related filters.
	// Here we require a file to import the "path/filepath" package.

	m.Match(`os.PathSeparator`).
		Where(m.File().Imports("path/filepath")).
		Suggest(`filepath.Separator`)

	m.Match(`os.PathListSeparator`).
		Where(m.File().Imports("path/filepath")).
		Suggest(`filepath.ListSeparator`)
}
```

<pre style="color: white; background-color: black">
$ ruleguard -c 0 -rules rules.go main.go
main.go:9:10: suggestion: filepath.Separator
9		println(os.PathSeparator)
main.go:10:10: suggestion: filepath.ListSeparator
10		println(os.PathListSeparator)
</pre>

<details><summary>main.go</summary>

```go
package main

import (
	"os"
	"path/filepath"
)

func main() {
	println(os.PathSeparator)
	println(os.PathListSeparator)

	// Good:
	println(filepath.Separator)
	println(filepath.ListSeparator)
}
```

</details>

<hr>

**Notes**:

* [`File.Name`](https://godoc.org/github.com/quasilyte/go-ruleguard/dsl#File) can be used to check the current file name
* [`File.PkgPath`](https://godoc.org/github.com/quasilyte/go-ruleguard/dsl#File) can be used to check the current file package path

<table><tr>
<td><a href="index">To index</a></td>
<td><a href="text-filters">Next: Text filters</a></td>
<td><a href="https://github.com/go-ruleguard/go-ruleguard.github.io/edit/master/by-example/file-predicates.md">Edit this page</a></td>
</tr></table>
