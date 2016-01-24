To use, put this in your .html file:

```
<!-- For BigInteger (see https://github.com/jasondavies/jsbn ). Any
class that follows the same interface will work. -->
<script src="jsbn.js"></script>
<script src="jsbn2.js"></script>

<!-- For isqrt(). -->
<script src="isqrt.js"></script>

<!-- For iroot(). -->
<script src="iroot.js"></script>
```

Then you can do:

```
var r1 = isqrt(new BigInteger('1000'));
var r2 = iroot(new BigInteger('1000'), 3);
```
