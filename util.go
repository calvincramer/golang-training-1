package main

// Collatz
func WhileLoopDoWork(n int) int {
	if n%2 == 0 {
		return n / 2
	}
	return (3 * n) + 1
}

type LargeStruct struct {
	A string
	B int64
	C uint16
	D int
	E uint
	F string
	G bool
	H string
	I int64
	J float32
	K string
	L string
	M uint16
	N float64
	O int
	P string
	Q string
	R int64
	S bool
	T uint16
	U string
	V int
	W uint
	X float32
	Y int
	Z string
}

type BinTreeNode struct {
	Val   any
	Left  *BinTreeNode
	Right *BinTreeNode
}
