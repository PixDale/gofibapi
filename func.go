package main

import "time"

// fibonacciRecursive é a função que calcula o enésimo elemento da sequência fibonacci
func fibonacciRecursive(x uint64) uint64 {
	if x <= 1 {
		return x
	}
	return fibonacciRecursive(x-1) + fibonacciRecursive(x-2)
}

// func fibonacci(n uint64) uint64 {
// 	a := 0
// 	b := 1
// 	for i := 0; i < n; i++ {
// 		var temp uint64 = a
// 		a = b
// 		b = temp + a
// 	}
// 	return a
// }

func startFib(input uint64, returnCh chan FibResult) {
	start := time.Now()
	result := fibonacciRecursive(input)
	duration := time.Since(start)

	returnCh <- FibResult{
		Input:    input,
		Result:   result,
		Duration: duration.String(),
	}

}

func callFib(input uint64, timeout time.Duration) FibResponse {
	result, ok := SavedResponses.Load(input)
	if ok {
		return FibResponse{
			Done: true,
			Fib:  result.(*FibResult),
		}
	}

	resultCh := make(chan FibResult, 1)
	go startFib(input, resultCh)

	select {
	case <-time.After(timeout):
		return FibResponse{
			Done: false,
			Fib:  nil,
		}

	case result := <-resultCh:
		SavedResponses.LoadOrStore(input, result)
		return FibResponse{
			Done: true,
			Fib:  &result,
		}
	}
}
