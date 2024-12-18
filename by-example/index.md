# Ruleguard by example

[Ruleguard](https://github.com/quasilyte/go-ruleguard) is a tool that makes it easy to write custom [Go](golang.org/) static analysis.

[Ruleguard by example](https://go-ruleguard.github.io/by-example/) is a hands-on introduction that helps you to learn how to write and run custom linting rules.

<hr>

1. [Hello world](hello-world) - a minimal working example to start out journey
1. [Submatches](submatches) - exploring named and unnamed pattern submatch variables
1. [Multi-rule groups](multi-rule-groups) - putting several rules inside one function, making them short-circuit
1. [Multi-pattern rules](multi-pattern-rules) - putting several patterns into a single rule
1. [Multi-statement patterns](multi-statement-patterns) - putting several statements into a single pattern
1. [Autofixable rules](autofixable-rules) - using [`Suggest()`](https://pkg.go.dev/github.com/quasilyte/go-ruleguard/dsl#Matcher.Suggest) to enable quickfix actions
1. [Optional submatches](optional-submatches) - usign `*` variable modifier to match variadic or optional parts of the pattern
1. [Filters](filters) - using [`Where()`](https://pkg.go.dev/github.com/quasilyte/go-ruleguard/dsl#Matcher.Where) to reject the unwanted matches
1. [Type filters](type-filters) - applying type-based constraints to your rules
1. [Type patterns](type-patterns) - mastering the advanced type pattern matching techniques
1. [External types](external-types) - using [`Import()`](https://pkg.go.dev/github.com/quasilyte/go-ruleguard/dsl#Matcher.Import) to bind qualified type names
1. [Underlying types](underlying-types) - understanding the underlying types in Go and how can we handle them in the Ruleguard
1. [Node predicates](node-predicates) - asserting the submatch nodes AST types
1. [File predicates](file-predicates) - applying file-scoped filters
1. [Text filters](text-filters) - matching "literal" variables (by their name) and more
1. [Nested AST filters](ast-filters) - matching an AST pattern inside other AST
1. [Constexpr evaluation](constexpr-evaluation) - writing filters based on the expression values
1. [Custom locations](custom-locations) - using [`At()`](https://pkg.go.dev/github.com/quasilyte/go-ruleguard/dsl#Matcher.At) to adjust the match location
1. [Custom filters](custom-filters) - writing own filters for the `Where()` clause
1. [Matching comments](matching-comments) - finding issues inside code comments
1. [Go version checks](go-version-checks) - limiting the rules to the specified Go version
1. [Deadcode detection](deadcode-detection) - handling the code reachability
1. [Local functions](local-functions) - reusing filter expressions in multiple `Where()` clauses

<hr>

<p align="center">
  <img src="ruleguard_poster.png">
</p>
