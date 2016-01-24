'use strict';

// isqrt returns the greatest number x such that x^2 <= n. The type of
// n must behave like BigInteger (e.g.,
// https://github.com/jasondavies/jsbn ), and n must be non-negative.
//
// See https://www.akalin.com/computing-isqrt for an analysis.
function isqrt(n) {
  var s = n.signum();
  if (s < 0) {
    throw new Error('negative radicand');
  }
  if (s == 0) {
    return n;
  }

  // x = 2^ceil(Bits(n)/2)
  var x = n.constructor.ONE.shiftLeft(Math.ceil(n.bitLength()/2));
  while (true) {
    // y = floor((x + floor(n/x))/2)
    var y = x.add(n.divide(x)).shiftRight(1);
    if (y.compareTo(x) >= 0) {
      return x;
    }
    x = y;
  }
}
