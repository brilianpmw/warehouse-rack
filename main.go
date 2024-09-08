package main

import (
	"os"
	"warehouse_rack/shell"
)

func main() {
	// Process the input file commands only
	if len(os.Args) > 1 && os.Args[1] != "" {
		shell.ProcessFile(os.Args[1])
		return
	}
	// Start interactive shell
	shell.Start()
}
