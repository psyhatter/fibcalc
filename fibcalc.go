// Package fibcalc implements the calculation of
// the Fibonacci number by raising the matrix to
// a power optimized enough to calculate large
// Fibonacci numbers.
package fibcalc

import (
	"math/big"
	"sync"
)

// Uint64 the best way to calculate Fibonacci
// numbers for small n and understand how to
// calculate a number by quickly raising the
// matrix to a power. Function takes an natural
// number [0,93] (uint64 limit).
func Uint64(n uint8) uint64 {
	if n < 2 {
		return uint64(n)
	}

	// It is based on the identity:
	// | 1 1 |^n = | Fn+1 Fn   |
	// | 1 0 |     | Fn   Fn-1 |

	// Temporary matrix for exponentiation
	// A = | a b |
	//     | c d |
	var a, b, c, d uint64 = 1, 1, 1, 0

	// Results vector
	// R = | rc rd |
	var rc, rd uint64 = 0, 1

	for n != 1 {
		// If the n is odd
		if n&1 != 0 {
			// Multiply the vector R by the matrix A
			rc, rd = rc*a+rd*c, rc*b+rd*d
		}

		// Multiply the matrix A by itself
		a, b, c, d = a*a+b*c, a*b+b*d, c*a+d*c, c*b+d*d

		// Halve the n
		n >>= 1
	}

	return rc*a + rd*c
}

// Sequential quite effective way to calculate Fibonacci
// numbers for natural n large 128, using sequential
// calculations, but to calculate really large numbers
// it’s more efficient to use Concurrent function. See
// benchmarks for details.
func Sequential(n uint64) *big.Int {
	// Optimization for numbers that can be obtained within uint64
	if n < 94 {
		return uint64calc(n)
	}

	var (
		// Temporary matrix for exponentiation
		// A = | a b | = | 1 1 |
		//     | c d |   | 1 0 |
		a, b = big.NewInt(1), big.NewInt(1)
		c, d = big.NewInt(1), big.NewInt(0)

		// Results vector
		// R = | rc rd | = | 0 1 |
		rc, rd = big.NewInt(0), big.NewInt(1)

		// Temporary variables for calculations
		// Fewer number of temporary variables can be dispensed with,
		// but for n > 128, calculations thus become more efficient
		tempA, tempB        = &big.Int{}, &big.Int{}
		copyA, copyB, copyC = &big.Int{}, &big.Int{}, &big.Int{}
	)

	for n != 1 {
		// If the n is odd
		if n&1 != 0 {
			// Temporary copy for calculations
			copyC.Set(rc)

			// Multiply the vector R by the matrix A
			// rc, rd = rc*a+rd*c, rc*b+rd*d
			rc.Add(tempA.Mul(rc, a), tempB.Mul(rd, c))
			rd.Add(tempA.Mul(rd, d), tempB.Mul(copyC, b))
		}

		// Temporary copy for calculations
		copyA.Set(a)
		copyB.Set(b)
		copyC.Set(c)

		// a, b, c, d = a*a+b*c, a*b+b*d, c*a+d*c, c*b+d*d
		a.Add(tempA.Mul(copyA, a), tempB.Mul(copyB, copyC))
		b.Add(tempA.Mul(copyA, copyB), tempB.Mul(copyB, d))
		c.Add(tempA.Mul(copyC, copyA), tempB.Mul(d, copyC))
		// This will save a few allocations for n > 2048
		copyA.Set(d)
		d.Add(tempA.Mul(copyC, copyB), tempB.Mul(d, copyA))

		// Halve the n
		n >>= 1
	}

	// rc*a+rd*c
	return rc.Add(tempA.Mul(rc, a), tempB.Mul(rd, c))
}

// Concurrent is the best way to calculate very large
// Fibonacci numbers using concurrent computing. See
// benchmarks for details.
func Concurrent(n uint64) *big.Int {
	// Optimization for numbers that can be obtained within uint64
	if n < 94 {
		return uint64calc(n)
	}

	var (
		// Temporary matrix for exponentiation
		// A = | a b | = | 1 1 |
		//     | c d |   | 1 0 |
		a, b = big.NewInt(1), big.NewInt(1)
		c, d = big.NewInt(1), big.NewInt(0)

		// Results vector
		// R = | rc rd | = | 0 1 |
		rc, rd = big.NewInt(0), big.NewInt(1)

		// Temporary variables for calculations
		aa, ab, bd = &big.Int{}, &big.Int{}, &big.Int{}
		ca, cb     = &big.Int{}, &big.Int{}
		dc, dd     = &big.Int{}, &big.Int{}

		wg sync.WaitGroup
	)

	for n != 1 {
		// If the n is odd
		if n&1 != 0 {
			// Preliminary calculations
			wg.Add(4)
			go func() { ca.Mul(rc, a); wg.Done() }()
			go func() { dc.Mul(rd, c); wg.Done() }()
			go func() { cb.Mul(rc, b); wg.Done() }()
			go func() { dd.Mul(rd, d); wg.Done() }()
			wg.Wait()

			// Multiply the vector R by the matrix A
			// rc, rd = rc*a+rd*c, rc*b+rd*d
			wg.Add(2)
			go func() { rc.Add(ca, dc); wg.Done() }()
			go func() { rd.Add(cb, dd); wg.Done() }()
			wg.Wait()
		}

		// Preliminary calculations
		wg.Add(7)
		go func() { aa.Mul(a, a); wg.Done() }()
		go func() { cb.Mul(b, c); wg.Done() }()
		go func() { ab.Mul(a, b); wg.Done() }()
		go func() { bd.Mul(b, d); wg.Done() }()
		go func() { ca.Mul(c, a); wg.Done() }()
		go func() { dc.Mul(d, c); wg.Done() }()
		go func() { dd.Mul(d, d); wg.Done() }()
		wg.Wait()

		// a, b, c, d = a*a+b*c, a*b+b*d, c*a+d*c, c*b+d*d
		wg.Add(4)
		go func() { a.Add(aa, cb); wg.Done() }()
		go func() { b.Add(ab, bd); wg.Done() }()
		go func() { c.Add(ca, dc); wg.Done() }()
		go func() { d.Add(cb, dd); wg.Done() }()
		wg.Wait()

		// Halve the n
		n >>= 1
	}

	//rc*a+rd*c
	return d.Add(ca.Mul(rc, a), dc.Mul(rd, c))
}

// uint64calc big.Int wrapper for the Uint64 function
func uint64calc(n uint64) *big.Int {
	result := &big.Int{}
	return result.SetUint64(Uint64(uint8(n)))
}
