/**
 * Author: Mitch Allen
 * File: alpha_test.go
 */

package alpha

import (
	"math"
	"testing"
)

func TestHello(t *testing.T) {
	// Just ensure it doesn't panic
	Hello()
}

func TestCoinCount(t *testing.T) {
	tests := []struct {
		name      string
		limit     int
		threshold float64
	}{
		{"small sample", 10, 0.0},
		{"medium sample", 100, 0.3},
		{"large sample", 1000, 0.4},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			expectedThreshold := int(math.Round(float64(tt.limit) * tt.threshold))
			got := CoinCount(tt.limit)

			// Verify total count is correct
			if got[true]+got[false] != tt.limit {
				t.Errorf("CoinCount(%d): total count = %d, want %d",
					tt.limit, got[true]+got[false], tt.limit)
			}

			// Verify we get at least some results (threshold test)
			if got[true] < expectedThreshold {
				t.Errorf("CoinCount(%d) = %v (heads below threshold: %d)", tt.limit, got, expectedThreshold)
			}

			if got[false] < expectedThreshold {
				t.Errorf("CoinCount(%d) = %v (tails below threshold: %d)", tt.limit, got, expectedThreshold)
			}
		})
	}
}

func BenchmarkCoinCount(b *testing.B) {
	for i := 0; i < b.N; i++ {
		CoinCount(100)
	}
}
