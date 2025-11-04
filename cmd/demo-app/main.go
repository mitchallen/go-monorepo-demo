// Author: Mitch Allen
// File: main.go

package main

import (
	"flag"
	"fmt"
	"os"

	"github.com/mitchallen/go-monorepo-demo/pkg/alpha"
	"github.com/mitchallen/go-monorepo-demo/pkg/beta"
	"github.com/mitchallen/go-monorepo-demo/pkg/shared"
)

func main() {
	logger := shared.NewLogger("demo-app")

	flips := flag.Int("flips", 100, "number of coin flips to simulate")
	mode := flag.String("mode", "analyze", "mode: hello, analyze, or compare")
	flag.Parse()

	logger.Info(fmt.Sprintf("Starting demo app in %s mode", *mode))

	switch *mode {
	case "hello":
		runHelloMode(logger)
	case "analyze":
		runAnalyzeMode(logger, *flips)
	case "compare":
		runCompareMode(logger)
	default:
		logger.Error(fmt.Sprintf("Unknown mode: %s", *mode))
		flag.Usage()
		os.Exit(1)
	}

	logger.Info("Demo app completed")
}

func runHelloMode(logger *shared.Logger) {
	logger.Info("Running hello mode")
	alpha.Hello()
	beta.Hello()
}

func runAnalyzeMode(logger *shared.Logger, flips int) {
	logger.Info(fmt.Sprintf("Running analyze mode with %d flips", flips))
	analysis := beta.AnalyzeCoinFlips(flips)

	fmt.Println("\n=== Coin Flip Analysis ===")
	fmt.Printf("Total Flips: %d\n", analysis["total"])
	fmt.Printf("Heads: %d\n", analysis["heads"])
	fmt.Printf("Tails: %d\n", analysis["tails"])
	fmt.Printf("Max: %d\n", analysis["max"])
	fmt.Printf("Min: %d\n", analysis["min"])
}

func runCompareMode(logger *shared.Logger) {
	logger.Info("Running compare mode")
	trials := []int{50, 100, 150}

	fmt.Println("\n=== Comparing Multiple Flip Series ===")
	for _, trial := range trials {
		fmt.Printf("\nTrial with %d flips:\n", trial)
		analysis := beta.AnalyzeCoinFlips(trial)
		fmt.Printf("  Heads: %d, Tails: %d\n", analysis["heads"], analysis["tails"])
	}

	results := beta.CompareFlipSeries(trials)
	fmt.Printf("\nOverall Statistics:\n")
	fmt.Printf("  Max Heads: %d\n", results["max_heads"])
	fmt.Printf("  Total Heads: %d\n", results["total"])
}
