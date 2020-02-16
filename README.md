# fibcalc
Package fibcalc implements the calculation of the Fibonacci number by raising the matrix to a power optimized enough to calculate large Fibonacci numbers.

## Getting Started
### Installing
```
go get github.com/psyhatter/fibcalc
```

### Usage
```go
package main

import (
	"github.com/psyhatter/fibcalc"
	"fmt"
	"math/big"
)

func main() {
	var (
		n1 uint64 = 1 << 10
		n2 uint64 = 1 << 20

		// Fermat prime https://oeis.org/A019434
		mod = big.NewInt(1<<(1<<4) + 1)

		// Temporary variable for calculations
		tmp = &big.Int{}
	)

	tmp.Mod(fibcalc.Sequential(n1), mod)
	fmt.Printf("F(%d) mod %d = %d\n", n1, mod, tmp)

	tmp.Mod(fibcalc.Concurrent(n2), mod)
	fmt.Printf("F(%d) mod %d = %d\n", n2, mod, tmp)
}
```

### And the output is:
```
F(1024) mod 65537 = 26749
F(1048576) mod 65537 = 49949
```

# Benchmark
## Comparison with other implementations (n = 1048575)
|implementation|ns/op|B/op|allocs/op|
|---|---|---|---|
|[massimo-marino](https://github.com/massimo-marino/go-fibonacci)|7296502900|220671040|4556|
|[T-PWK](https://github.com/T-PWK/go-fibonacci)|7260525000|220671024|4554|
|[visualfc](https://github.com/visualfc/fibutil)|134541862|8177616|1254|
|fibcalc.Sequential|92443066|2669724|227|
|fibcalc.Concurrent|58408427|4921236|499|