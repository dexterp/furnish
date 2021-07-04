package main

import (
	"os"

	"github.com/dexterp/furnish/internal"
	"github.com/dexterp/furnish/internal/options"
)

func main() {
	opts := options.GetUsage(os.Args[1:], internal.Version)
	_ = opts
}
