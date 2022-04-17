# sqlconv

[![CI](https://github.com/dochang/sqlconv/actions/workflows/ci.yml/badge.svg)](https://github.com/dochang/sqlconv/actions/workflows/ci.yml)
[![codecov](https://codecov.io/gh/dochang/sqlconv/branch/master/graph/badge.svg?token=qmz7MW4X1k)](https://codecov.io/gh/dochang/sqlconv)
[![Go Reference](https://pkg.go.dev/badge/github.com/dochang/sqlconv.svg)](https://pkg.go.dev/github.com/dochang/sqlconv)

`sqlconv` is a Go library converts between Go values and SQL values.

## Installation

```sh
go get github.com/dochang/sqlconv
```

## Usage

Currently, `sqlconv` provides a interface `ScannerConverter`, which is used to
convert SQL values to Go values.

```go
type ScannerConverter interface {
	ScanConvert(dest, src any) error
}
```

`ScanConvert` converts SQL value `src` to Go value, then writes the Go value
into `dest`.

A predefined variable `DefaultScannerConverter` converts SQL values based on
the rules defined by the function `"database/sql".convertAssign`.
