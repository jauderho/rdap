package main

import (
	"os"

	"github.com/jauderho/rdap"
)

func main() {
	exitCode := rdap.RunCLI(os.Args[1:], os.Stdout, os.Stderr, rdap.CLIOptions{})

	os.Exit(exitCode)
}
