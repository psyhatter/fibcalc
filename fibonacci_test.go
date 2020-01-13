package fibonacci

import "testing"

func TestFibonacci(t *testing.T) {
	type args struct {
		n int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{"10", struct{ n int }{n: 10}, 55},
		{"20", struct{ n int }{n: 20}, 6765},
		{"30", struct{ n int }{n: 30}, 832040},
		{"40", struct{ n int }{n: 40}, 102334155},
		{"50", struct{ n int }{n: 50}, 12586269025},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := Fibonacci(tt.args.n); got != tt.want {
				t.Errorf("Fibonacci() = %v, want %v", got, tt.want)
			}
		})
	}
}

func BenchmarkFibonacci10(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Fibonacci(10)
	}
}

func BenchmarkFibonacci20(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Fibonacci(20)
	}
}

func BenchmarkFibonacci30(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Fibonacci(30)
	}
}

func BenchmarkFibonacci40(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Fibonacci(40)
	}
}

func BenchmarkFibonacci50(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Fibonacci(50)
	}
}
