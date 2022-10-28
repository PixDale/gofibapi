package main

type FibResult struct {
	Input    uint64 `json:"input"`
	Result   uint64 `json:"result"`
	Duration string `json:"duration"`
}
type FibResponse struct {
	Done bool       `json:"done"`
	Fib  *FibResult `json:"fib,omitempty"`
}

type AllResponse struct {
	Queries []FibResult `json:"queries"`
}
