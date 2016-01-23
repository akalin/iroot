package iroot

import "math/big"

// ISqrt returns the greatest number x such that x^2 <= n. n must be
// non-negative.
//
// See https://www.akalin.com/computing-isqrt for an analysis.
func ISqrt(n *big.Int) *big.Int {
	s := n.Sign()
	if s < 0 {
		panic("negative radicand")
	}
	if s == 0 {
		return &big.Int{}
	}

	// x = 2^ceil(Bits(n)/2)
	var x big.Int
	x.Lsh(big.NewInt(1), (uint(n.BitLen())+1)/2)
	for {
		// y = floor((x + floor(n/x))/2)
		var y big.Int
		y.Div(n, &x)
		y.Add(&y, &x)
		y.Rsh(&y, 1)

		if y.Cmp(&x) >= 0 {
			return &x
		}
		x = y
	}
}
