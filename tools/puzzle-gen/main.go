package main

import (
	"encoding/json"
	"fmt"
	"io"
	"math/rand/v2"
	"os"
	"strings"

	"project/tools/puzzlelib"
)

type CLI struct {
	Seed uint64 `help:"RNG seed for reproducibility" default:"1"`
}

func main() {
	var cli CLI
	// Minimal parsing: --seed N is optional, message comes from stdin.
	for i := 1; i < len(os.Args); i++ {
		if os.Args[i] == "--seed" && i+1 < len(os.Args) {
			var s uint64
			_, _ = fmt.Sscanf(os.Args[i+1], "%d", &s)
			cli.Seed = s
			i++
		}
	}

	msgBytes, err := io.ReadAll(os.Stdin)
	if err != nil {
		_, _ = fmt.Fprintln(os.Stderr, "error:", err)
		os.Exit(1)
	}
	message := string(msgBytes)
	message = strings.TrimRight(message, "\n")

	rng := rand.New(rand.NewPCG(cli.Seed, cli.Seed^0x9e3779b97f4a7c15))
	start, err := puzzlelib.Generate(message, rng)
	if err != nil {
		fmt.Fprintln(os.Stderr, "error:", err)
		os.Exit(1)
	}

	out, err := json.MarshalIndent(start, "", "  ")
	if err != nil {
		fmt.Fprintln(os.Stderr, "error:", err)
		os.Exit(1)
	}
	_, _ = os.Stdout.Write(out)
}
