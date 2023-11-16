# Learn GO with tests

This repository contains the code related to the exercises proposed at [Learn GO with tests](https://quii.gitbook.io/learn-go-with-tests/)

## Requirements

- [Golang](https://go.dev/doc/install)

## How to operate

- Run all tests:

```shell
go test ./... [-cover]
```

- Run module tests:

```shell
go test ./module-folder [-cover]
```

- Run benchmark: (only `BenchmarkRepeat` in `iteration` module right now)

```shell
 go test ./iteration -benchmem -run=^$ -bench ^BenchmarkRepeat$ 
```
