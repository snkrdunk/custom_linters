# custom_linters

# Features

custom_linters is a tool to run multiple linters at once using [the multichecker package](https://pkg.go.dev/golang.org/x/tools/go/analysis/multichecker).

# Installation

```console
go install github.com/snkrdunk/custom_linters/cmd/custom_linters@latest
```

# How to add analyzer

## 1. Add linter module to go.mod

```console
go get github.com/path/to/module
```

**If you want to add a private linter module, set GOPRIVATE to the repository or account path before executing `go get` command**
```console
export GOPRIVATE=github.com/snkrdunk
```
In this example, you can get private module from github.com/snkrdunk.

## 2. Add analyzer to cmd/custom_linters/custom_linters.go

```diff
 package main

 import (
+       "path/to/your/linter/your_linter_pkg"

        "github.com/snkrdunk/empty_err_checker"
        "golang.org/x/tools/go/analysis/multichecker"
 )
 
 func main() {
        multichecker.Main(
                empty_err_checker.Analyzer,
+               your_linter_pkg.Analyzer,
        )
 }
```

## 3. Add analyzer plugin/main.go

```diff
 package main

 import (
+       "path/to/your/linter/your_linter_pkg"

        "github.com/snkrdunk/empty_err_checker"
        "golang.org/x/tools/go/analysis"
 )

 func (analyzerPlugin) GetAnalyzers() []*analysis.Analyzer {
        return []*analysis.Analyzer{
                empty_err_checker.Analyzer,
+               your_linter_pkg.Analyzer,
        }
 }
```
