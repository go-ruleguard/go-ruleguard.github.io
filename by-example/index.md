# Ruleguard by example

[Ruleguard](https://github.com/quasilyte/go-ruleguard) is a tool that makes it easy to write custom [Go](golang.org/) static analysis.

[Ruleguard by example](https://go-ruleguard.github.io/by-example/) is a hands-on introduction that helps you to learn how to write and run custom linting rules.

<hr>

1. [Hello world](hello-world) - a minimal working example to start out journey
1. [Submatches](submatches) - exploring named and unnamed pattern submatch variables
1. [Multi-rule groups](multi-rule-groups) - putting several rules inside one function, making them short-circuit
1. [Autofixable rules](autofixable-rules) - using [`Suggest()`](https://pkg.go.dev/github.com/quasilyte/go-ruleguard/dsl/fluent#Matcher.Suggest) to enable quickfix actions
1. [Optional submatches](optional-submatches) - usign `*` variable modifier to match variadic or optional parts of the pattern
