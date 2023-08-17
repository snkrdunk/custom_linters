package main

import (
	"github.com/toshiki-otaka/argument_checker"
	"golang.org/x/tools/go/analysis/multichecker"
)

func main() {
	multichecker.Main(
		argument_checker.Analyzer,
	)
}
