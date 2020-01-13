package fibonacci

// Fibonacci implements the Fibonacci number search using the fast method
// raising a matrix to a power. It is based on the identity:
// | 1 1 |^n = | Fn+1 Fn   |
// | 1 0 |     | Fn   Fn-1 |
// Function takes an integer [1,93] (uint64 limit)
func Fibonacci(n int) uint64 {
	// Temporary matrix for exponentiation
	// A = | a b |
	//     | c d |
	var a, b, c, d uint64 = 1, 1, 1, 0

	// Results vector
	// R = | rc rd |
	var rc, rd uint64 = 0, 1

	for n != 0 {
		// If the degree is odd
		if n&1 == 1 {
			// Multiply the vector R by the matrix A
			rc, rd = rc*a+rd*c, rc*b+rd*d
		}

		// Multiply the matrix A by itself
		a, b, c, d = a*a+b*c, a*b+b*d, c*a+d*c, c*b+d*d

		// Halve the degree
		n >>= 1
	}

	return rc
}
