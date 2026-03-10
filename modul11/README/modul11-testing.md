# Modul 11 — Testing (Unit Test & Benchmark)

## Cara menjalankan test

Dari root project:

```bash
go test ./...
```

Menjalankan test pada satu package saja:

```bash
go test ./modul11/mathutil
go test ./modul10/calculator
```

## Table-driven test & subtest

Contoh ada di:
- `modul11/mathutil/mathutil_test.go` (`TestSum_TableDriven`, `t.Run`)
- `modul10/calculator/calculator_test.go` (`TestMul_TableDriven`)

## Benchmark

Benchmark ada di:
- `modul11/mathutil/mathutil_bench_test.go`

Jalankan benchmark:

```bash
go test -bench . ./modul11/mathutil
```

Tambahkan memori allocation info:

```bash
go test -bench . -benchmem ./modul11/mathutil
```

