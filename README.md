# Project Euler 612 Inspired Micro-Optimization

This doesn't solve [612](https://projecteuler.net/problem=612). It's intended
to demonstrate that implementing decimal math isn't so scary and can be useful
for performance as well as currency. The decimal version had a lot more bugs,
though, and isn't that much faster. Point only sort of made.

There's one place (indicated in comments) where I can think of a further
optimization. Also, I didn't go so far as to start looking at generated
assembly and trying to find unnecessary bounds checks and such.

|places|seconds|
|------|-------|
|5     |  32   |
|4     |   0.62|

```
$ go test -bench .
goos: darwin
goarch: amd64
cpu: VirtualApple @ 2.50GHz
BenchmarkDecimal-8         65095         18364 ns/op
BenchmarkBinary-8          41398         28948 ns/op
PASS
ok      _/Users/sk/project-euler-612    3.056s
$ 
```
