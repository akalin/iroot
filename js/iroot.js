'use strict';

// iroot returns the greatest number x such that x^p <= n. The type of
// n must behave like BigInteger (e.g.,
// https://github.com/jasondavies/jsbn ), n must be non-negative, and
// p must be a positive integer.
//
// See https://www.akalin.com/computing-iroot for an analysis.
function iroot(n, p) {
  var s = n.signum();
  if (s < 0) {
    throw new Error('negative radicand');
  }
  if (p <= 0) {
    throw new Error('non-positive degree');
  }
  if (p !== (p|0)) {
    throw new Error('non-integral degree');
  }

  if (s == 0) {
    return n;
  }

  var b = n.bitLength();
  if (p >= b) {
    return n.constructor.ONE;
  }

  // x = 2^ceil(Bits(n)/p)
  var x = n.constructor.ONE.shiftLeft(Math.ceil(b/p));
  var pMinusOne = new n.constructor((p - 1).toString());
  var pBig = new n.constructor(p.toString());
  while (true) {
    // y = floor(((p-1)x + floor(n/x^(p-1)))/p)
    var y = pMinusOne.multiply(x).add(n.divide(x.pow(pMinusOne))).divide(pBig);
    if (y.compareTo(x) >= 0) {
      return x;
    }
    x = y;
  }
}
