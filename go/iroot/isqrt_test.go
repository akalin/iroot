package iroot

import (
	"math"
	"math/big"
	"testing"
)

// TestISqrt16Bit tests ISqrt() with all 16-bit numbers.
func TestISqrt16Bit(t *testing.T) {
	for n := 0; n <= math.MaxInt16; n++ {
		expectedR := int64(math.Sqrt(float64(n)))
		r := ISqrt(big.NewInt(int64(n)))
		if r.Cmp(big.NewInt(expectedR)) != 0 {
			t.Errorf("For n=%d, expected r=%d, got r=%s",
				n, expectedR, r)
		}
	}
}
