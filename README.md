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
|[massimo-marino](https://github.com/massimo-marino/go-fibonacci)|7323484400|220671040|4556|
|[T-PWK](https://github.com/T-PWK/go-fibonacci)|7317491700|220671120|4555|
|[sevlyar](https://github.com/sevlyar/fibonacci)|211470120|7720432|994|
|[visualfc](https://github.com/visualfc/fibutil)|135291675|8177616|1254|
|fibcalc.Sequential|92109933|2669724|227|
|fibcalc.Concurrent|60074050|4921615|500|

## Comparison of sequential implementation and concurrent
### H3 Quad-core Cortex-A7
![](https://github.com/psyhatter/fibcalc/blob/master/graphs/H3%20Quad-core%20Cortex-A7/0-60000.PNG?raw=true)
![](https://github.com/psyhatter/fibcalc/blob/master/graphs/H3%20Quad-core%20Cortex-A7/0-15000000.PNG?raw=true)

### Intel(R) Celeron(R) CPU 1005M @ 1.90GHz
![](https://github.com/psyhatter/fibcalc/blob/master/graphs/Intel(R)%20Celeron(R)%20CPU%201005M%20@%201.90GHz/0-60000.PNG?raw=true)
![](https://github.com/psyhatter/fibcalc/blob/master/graphs/Intel(R)%20Celeron(R)%20CPU%201005M%20@%201.90GHz/0-15000000.PNG?raw=true)

### Intel(R) Core(TM) i5-7500 CPU @ 3.40GHz
![](https://github.com/psyhatter/fibcalc/blob/master/graphs/Intel(R)%20Core(TM)%20i5-7500%20CPU%20@%203.40GHz/0-125000.PNG?raw=true)
![](https://github.com/psyhatter/fibcalc/blob/master/graphs/Intel(R)%20Core(TM)%20i5-7500%20CPU%20@%203.40GHz/0-15000000.PNG?raw=true)

### AMD Ryzen 5 3600 6-Core Processor
![](https://github.com/psyhatter/fibcalc/blob/master/graphs/AMD%20Ryzen%205%203600%206-Core%20Processor/0-125000.PNG?raw=true)
![](https://github.com/psyhatter/fibcalc/blob/master/graphs/AMD%20Ryzen%205%203600%206-Core%20Processor/0-15000000.PNG?raw=true)