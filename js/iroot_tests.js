'use strict';

describe('iroot', function() {
  // From isqrt_tests.js.
  beforeEach(function() {
    jasmine.addCustomEqualityTester(bigIntegerEquality);
  });

  it('iroot with all 8-bit numbers and 2', function() {
    for (var n = 0; n < (1 << 8); n++) {
      var expectedR = Math.floor(Math.sqrt(n));
      var r = iroot(new BigInteger(n.toString()), 2);
      expect(r.toString()).toEqual(expectedR.toString());
    }
  });

  it('iroot with all 8-bit numbers and 4', function() {
    for (var n = 0; n < (1 << 8); n++) {
      var expectedR = Math.floor(Math.sqrt(Math.sqrt(n)));
      var r = iroot(new BigInteger(n.toString()), 4);
      if (r.toString() != expectedR.toString()) {
        console.log(n, r.toString(), expectedR.toString());
      }
      expect(r.toString()).toEqual(expectedR.toString());
    }
  });

  it('iroot with large p', function() {
    var p = 100;
    var n = BigInteger.ONE.shiftLeft(p-1);
    var r = iroot(n, p);
    expect(r.toString()).toEqual('1');
  });
});
