package main

import (
	"fmt"
	"os"

	"github.com/cetincetindag/dfmgr/cmd"
)

func main() {
	if err := cmd.Execute(); err != nil {
		fmt.Fprintf(os.Stderr, "Error: %s\n", err)
		os.Exit(1)
	}
} 