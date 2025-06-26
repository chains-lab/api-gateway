package main

import (
	"os"

	"github.com/chains-lab/api-gateway/cmd/cli"
)

func main() {
	if !cli.Run(os.Args) {
		os.Exit(1)
	}
}
