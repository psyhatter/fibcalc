package fibcalc_test

import (
	"fmt"

	"github.com/psyhatter/fibcalc"
)

func ExampleUint64() {
	var arg uint8
	for ; arg < 7; arg++ {
		fmt.Printf("F(%d) = %d\n", arg, fibcalc.Uint64(arg))
	}

	// Output:
	// F(0) = 0
	// F(1) = 1
	// F(2) = 1
	// F(3) = 2
	// F(4) = 3
	// F(5) = 5
	// F(6) = 8
}

func ExampleSequential() {
	var arg uint64
	for ; arg < 7; arg++ {
		fmt.Printf("F(%d) = %d\n", arg, fibcalc.Sequential(arg))
	}

	// Output:
	// F(0) = 0
	// F(1) = 1
	// F(2) = 1
	// F(3) = 2
	// F(4) = 3
	// F(5) = 5
	// F(6) = 8
}

func ExampleConcurrent() {
	var arg uint64
	for ; arg < 7; arg++ {
		fmt.Printf("F(%d) = %d\n", arg, fibcalc.Concurrent(arg))
	}

	// Output:
	// F(0) = 0
	// F(1) = 1
	// F(2) = 1
	// F(3) = 2
	// F(4) = 3
	// F(5) = 5
	// F(6) = 8
}
