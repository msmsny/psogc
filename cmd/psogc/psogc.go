package main

import (
	"fmt"
	"os"

	cmd "github.com/msmsny/psogc/psogc/cmd/psogc"
)

func main() {
	if err := cmd.Execute(); err != nil {
		fmt.Fprintf(os.Stderr, "%+v\n", err)
		os.Exit(1)
	}
}
