package main

import (
	"fmt"
	"io"
	"os"

	"project/tools/puzzlelib"
)

func main() {
	b, err := io.ReadAll(os.Stdin)
	if err != nil {
		fmt.Fprintln(os.Stderr, "error:", err)
		os.Exit(1)
	}

	res, err := puzzlelib.SolveFromStartTextTrace(string(b), os.Stderr)
	if err != nil {
		fmt.Fprintln(os.Stderr, "error:", err)
		os.Exit(1)
	}

	_, _ = fmt.Fprintf(os.Stderr, "Solved %d steps.\n", res.StepsSolved)
	fmt.Print(res.FinalText)
}
