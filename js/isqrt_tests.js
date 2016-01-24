'use strict';

describe('isqrt', function() {
  var bigIntegerEquality = function(first, second) {
    if (first instanceof BigInteger && second instanceof BigInteger) {
      return first.equals(second);
    }
  };

  beforeEach(function() {
    jasmine.addCustomEqualityTester(bigIntegerEquality);
  });

  it('isqrt with all 8-bit numbers', function() {
    for (var n = 0; n < (1 << 8); n++) {
      var expectedR = Math.floor(Math.sqrt(n));
      var r = isqrt(new BigInteger(n.toString()));
      expect(r).toEqual(new BigInteger(expectedR.toString()));
    }
  });
});
