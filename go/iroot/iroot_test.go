package iroot

import (
	"math"
	"math/big"
	"testing"
)

// TestIRoot16BitSqrt tests IRoot() with all 16-bit numbers and 2.
func TestIRoot16BitSqrt(t *testing.T) {
	for n := 0; n <= math.MaxInt16; n++ {
		expectedR := int64(math.Sqrt(float64(n)))
		r := IRoot(big.NewInt(int64(n)), 2)
		if r.Cmp(big.NewInt(expectedR)) != 0 {
			t.Errorf("For n=%d and p=2, expected r=%d, got r=%s",
				&n, expectedR, r)
		}
	}
}

// TestIRoot16BitCbrt tests IRoot() with all 16-bit numbers and 3.
func TestIRoot16BitCbrt(t *testing.T) {
	for n := 0; n <= math.MaxInt16; n++ {
		expectedR := int64(math.Cbrt(float64(n)))
		r := IRoot(big.NewInt(int64(n)), 3)
		if r.Cmp(big.NewInt(expectedR)) != 0 {
			t.Errorf("For n=%d and p=3, expected r=%d, got r=%s",
				&n, expectedR, r)
		}
	}
}

// TestIRootLargeP tests IRoot() with p >= n.BitLen().
func TestIRootLargeP(t *testing.T) {
	var p uint = 100
	var n big.Int
	n.Lsh(big.NewInt(1), p-1)
	r := IRoot(&n, p)
	if r.Cmp(big.NewInt(1)) != 0 {
		t.Errorf("For n=%s and p=%d, expected r=1, got r=%s",
			&n, p, r)
	}
}
