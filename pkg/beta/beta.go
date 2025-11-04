// Author: Mitch Allen
// File: beta.go

package beta

import (
	"fmt"

	"github.com/mitchallen/go-monorepo-demo/pkg/alpha"
	"github.com/mitchallen/go-monorepo-demo/pkg/shared"
)

// Hello prints a greeting from beta and calls alpha
func Hello() {
	logger := shared.NewLogger("beta")
	logger.Info("Hello from beta!")
	alpha.Hello()
}

// AnalyzeCoinFlips runs coin flips and analyzes the results
func AnalyzeCoinFlips(limit int) map[string]int {
	logger := shared.NewLogger("beta")
	logger.Info(fmt.Sprintf("Analyzing %d coin flips", limit))

	results := alpha.CoinCount(limit)

	analysis := map[string]int{
		"heads": results[true],
		"tails": results[false],
		"max":   shared.Max(results[true], results[false]),
		"min":   shared.Min(results[true], results[false]),
		"total": limit,
	}

	logger.Info(fmt.Sprintf("Results - Heads: %d, Tails: %d", analysis["heads"], analysis["tails"]))

	return analysis
}

// CompareFlipSeries runs multiple flip series and returns statistics
func CompareFlipSeries(trials []int) map[string]int {
	totals := make([]int, len(trials))

	for i, trial := range trials {
		results := alpha.CoinCount(trial)
		totals[i] = results[true]
	}

	return map[string]int{
		"max_heads": shared.Max(totals[0], shared.Max(totals[1], totals[len(totals)-1])),
		"total":     shared.Sum(totals),
	}
}
