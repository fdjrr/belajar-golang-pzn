# Golang Unit Test

Slide : https://docs.google.com/presentation/d/1XxMEaA-JsPHr9BUw2oIOPlEL_psI3EaUFUpuvdlDB_Q/edit?usp=sharing

Source Code : https://github.com/ProgrammerZamanNow/belajar-golang-unit-test

# Cara Menjalankan Unit Test

```bash
$ go test -v // semua test
$ go test -v -run nama_function 
$ go test -v ./...
$ go test -v -run nama_function/subtest // subtest
$ go test -v -run /subtest // subtest
```

# Cara Menjalankan Benchmark

```bash
$ go test -v -bench=. // semua test dan benchmark di running
$ go test -v -run=NotMatchUnitTest -bench=. // semua benchmark di running, terkecuali test
$ go test -v -run=NotMatchUnitTest -bench=BenchmarkTest // hanya benchmark yang di running
$ go test -v -bench=. ./...  // semua benchmark yang ada di module
$ go test -v -run=NotMatchUnitTest -bench=BenchmarkHelloWorldSub/Fadjrir // sub benchmark
```