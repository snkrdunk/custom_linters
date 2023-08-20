package main

import (
	"github.com/snkrdunk/empty_err_checker"
	"golang.org/x/tools/go/analysis/multichecker"
)

func main() {
	multichecker.Main(
		empty_err_checker.Analyzer,
	)
}
