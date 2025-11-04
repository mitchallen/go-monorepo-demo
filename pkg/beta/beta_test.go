// Author: Mitch Allen
// File: beta_test.go

package beta

import (
	"testing"
)

func TestHello(t *testing.T) {
	// Just ensure it doesn't panic
	Hello()
}

func TestAnalyzeCoinFlips(t *testing.T) {
	limit := 100
	analysis := AnalyzeCoinFlips(limit)

	if analysis["total"] != limit {
		t.Errorf("Expected total %d, got %d", limit, analysis["total"])
	}

	if analysis["heads"]+analysis["tails"] != limit {
		t.Errorf("Heads + Tails should equal %d, got %d", limit, analysis["heads"]+analysis["tails"])
	}

	if analysis["max"] < analysis["min"] {
		t.Errorf("Max %d should be >= Min %d", analysis["max"], analysis["min"])
	}
}

func TestCompareFlipSeries(t *testing.T) {
	trials := []int{10, 20, 30}
	results := CompareFlipSeries(trials)

	if results["total"] < 0 || results["total"] > 60 {
		t.Errorf("Total should be between 0 and 60, got %d", results["total"])
	}

	if results["max_heads"] < 0 {
		t.Errorf("Max heads should be non-negative, got %d", results["max_heads"])
	}
}

func BenchmarkAnalyzeCoinFlips(b *testing.B) {
	for i := 0; i < b.N; i++ {
		AnalyzeCoinFlips(100)
	}
}
