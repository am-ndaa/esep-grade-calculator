# ESEP Grade Calculator

This library is a simple utility to calculate your grades.

Runtime requirements:
Golang 1.24+

Run unit tests:
```
cd grade_calculator/
go test ./...
```

Run unit tests with code coverage:
```
cd grade_calculator/
go test -coverprofile=coverage.txt ./...
```

File Summary:
```
go.mod: contains the path and the go version
grade_calculator.go: implements the grade calculator
grade_calculator_test.go: contains tests