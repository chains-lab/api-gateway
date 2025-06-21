package main

import (
	"os"

	"github.com/chains-lab/api-gateway/internal/cli"
)

func main() {
	if !cli.Run(os.Args) {
		os.Exit(1)
	}
}
