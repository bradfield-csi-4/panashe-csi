Initial benchmark results:

```
❯ go test -bench=.
OK: 4 passed
goos: darwin
goarch: arm64
pkg: iguazu-db
BenchmarkDB/Put-8                 596074             53066 ns/op
BenchmarkDB/Get-8               10665912               100.2 ns/op
BenchmarkDB/RangeScan-8              100          33661052 ns/op
BenchmarkDB/Delete-8               29480             63186 ns/op
PASS
ok      iguazu-db       65.163s
```
