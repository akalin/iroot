package iroot

import "math/big"

// IRoot returns the greatest number x such that x^p <= n. n must be
// non-negative, and p must be positive.
//
// See https://www.akalin.com/computing-iroot for an analysis.
func IRoot(n *big.Int, p uint) *big.Int {
	s := n.Sign()
	if s < 0 {
		panic("negative radicand")
	}
	if p == 0 {
		panic("zero degree")
	}
	if s == 0 {
		return &big.Int{}
	}

	b := uint(n.BitLen())
	if p >= b {
		return big.NewInt(1)
	}

	// x = 2^ceil(Bits(n)/p)
	var x big.Int
	x.Lsh(big.NewInt(1), (b+(p-1))/p)
	pMinusOne := big.NewInt(int64(p - 1))
	pBig := big.NewInt(int64(p))
	for {
		// y = floor(((p-1)x + floor(n/x^(p-1)))/p)
		var y big.Int
		y.Exp(&x, pMinusOne, nil)
		y.Div(n, &y)
		var t big.Int
		t.Mul(pMinusOne, &x)
		y.Add(&y, &t)
		y.Div(&y, pBig)

		if y.Cmp(&x) >= 0 {
			return &x
		}
		x = y
	}
}
