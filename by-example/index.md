# Ruleguard by example

[Ruleguard](https://github.com/quasilyte/go-ruleguard) is a tool that makes it easy to write custom [Go](golang.org/) static analysis.

[Ruleguard by example](https://go-ruleguard.github.io/by-example/) is a hands-on introduction that helps you to learn how to write and run custom linting rules.

<hr>

1. [Hello world](hello-world) - a minimal working example to start out journey
1. [Submatches](submatches) - exploring named and unnamed pattern submatch variables
1. [Multi-rule groups](multi-rule-groups) - putting several rules inside one function, making them short-circuit
1. [Autofixable rules](autofixable-rules) - using [`Suggest()`](https://pkg.go.dev/github.com/quasilyte/go-ruleguard/dsl/fluent#Matcher.Suggest) to enable quickfix actions
1. [Optional submatches](optional-submatches) - usign `*` variable modifier to match variadic or optional parts of the pattern
1. Match filters - using [`Where()`](https://pkg.go.dev/github.com/quasilyte/go-ruleguard/dsl/fluent#Matcher.Where) to skip the unrelated matches
1. Type patterns - applying advanced type filters
1. Underlying types - understanding the underlying types in Go and how can we handle them in the Ruleguard
1. External types - using [`Import()`](https://pkg.go.dev/github.com/quasilyte/go-ruleguard/dsl/fluent#Matcher.Import) to bind qualified type names
1. File predicates - applying file-scoped filters
1. Text filters - matching "literal" variables (by their name) and more
1. Custom locations - using [`At()`](https://pkg.go.dev/github.com/quasilyte/go-ruleguard/dsl/fluent#Matcher.At) to adjust the match location
