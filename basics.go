package main

import (
	"fmt"
	"math"
)

// TODO: return a string starting with "Hi, my name is "
func Hello() string {
	return "TODO"
}

// TODO: return zero value for uint64 WITHOUT USING ANY LITERALS
func ZeroInt() uint64 {
	return math.MaxInt
}

// TODO: return zero value for float32 WITHOUT USING ANY LITERALS
func ZeroFloat() float32 {
	return float32(math.MaxInt)
}

// TODO: return zero value for boolean WITHOUT USING ANY LITERALS**without** using the literal directly
func ZeroBool() bool {
	return true
}

// TODO: return zero value for string WITHOUT USING ANY LITERALS**without** using the literal directly
func ZeroString() string {
	return "TODO"
}

// TODO: convert integer to base-10 string value. For example 123 -> "123"
func IntToString(n int) string {
	return "TODO"
}

// TODO: convert a base-10 string to its integer value. If not possible then return an error.
func StringToInt(s string) (int, error) {
	return 0, fmt.Errorf("TODO")
}

// TODO: if n is even, return n divided by 2 and n squared, otherwise return n minus one and n modulo 7
func MultipleReturn(n int) (int, int) {
	return 0, 0
}

// TODO: construct a map that has the following key-value pairs:
//
//	"sagan": 30023
//	"mendeleev": 101
func MakeMap() map[string]int {
	return nil
}

// TODO: remove all key-value pairs from the input map
func MapClear(m map[uint64]any) {}

// TODO: manipulate the map according to the following rules:
//
//  1. if the key 2000 is present add the KV pair (4000, "XYZ")
//  2. if the key 100 is present, delete it and the associated value
func MapManipulate(m map[int]string) {}

// TODO: return the sum of odd number in the range 1,000 to 100,000
func ForLoop() int {
	return 0
}

// TODO: starting with the number 5231 call the WhileLoopDoWork() function, and continue calling
// this function with the return value until reaching the number 1. Return the amount of calls to
// WhileLoopDoWork() performed.
// Hint: do not try to predict what the WhileLoopDoWork() function does.
func WhileLoop() int {
	return 0
}

// TODO: use range() to iterate over an array. Return the total calculated in the following manner:
//
//	Each element is added to the total
//	If an element is divisible by 7, add its index to the total
func ForRangeArray(arr []int) int {
	return 0
}

// TODO: use range() to iterate over a map. Return the minimum value of all the keys that have a
// prefix of "be". If no KV pairs satisfy this condition then return a non-nil error.
// Hint: you can use math.MaxInt. The map values will always be less than math.MaxInt.
func ForRangeMap(m map[string]int) (int, error) {
	return 0, fmt.Errorf("TODO")
}

// TODO: return an array with the following items:
//
//  0. n multiplied by 10
//  1. the decimal string value of n. Hint: we did this before.
//  2. the lower 9 bits of n
func CreateArray(n int) [3]any {
	return [3]any{nil, nil, nil}
}

// TODO: return a slice constructed in the following manner: while n is positive add n to the end of
// the slice and then reduce n by half.
// For example, n=579 should produce [579, 289, 144, 72, 36, 18, 9, 4, 2, 1]
func CreateSlice(n int) []int {
	return []int{}
}

// TODO: make a slice of strings with length n. Each string should be the zero value ""
// Hint: use make()
func MakeSlice(n int) []string {
	return []string{}
}

// TODO: implement the merge step of merge sort. Specifically, left and right are guaranteed to be
// sorted, return a sorted slice containing all elements from left and right.
// left and right are of arbitrary length.
// Hint: use append(foo, bar...) to concatenate slices
func SliceMerge(left []int, right []int) []int {
	return []int{}
}

// TODO: return two LargeStructs initialized in the following manner:
// first struct: all fields are zero initialized
// second struct: all fields are zero initialized except for:
//   - D should be 42
//   - Q should be "qqq"
//   - X should be 3.1415
func InitStruct() (LargeStruct, LargeStruct) {
	return LargeStruct{}, LargeStruct{}
}

// TODO: initialize a binary tree that has the following structure:
//
//		 100
//		/   \
//	  "W"   8.7
//
// Hint: use new() to create pointer to BinTreeNode struct
func InitTree() BinTreeNode {
	return BinTreeNode{}
}

type EnumTestT int

const (
	EnumTestA EnumTestT = 0 // TODO
	EnumTestB
	EnumTestC
	EnumTestD
	EnumTestE
	EnumTestF
	EnumTestG
	EnumTestH
	EnumTestI
	EnumTestJ
	EnumTestK
)

// TODO: fill out the above "enum" with the following values: 1, 2, 4, 8, 16, 32, 64, ...
// Hint: use 'iota' to get the right number
// Hint: use implicit repetition to only specify EnumTestA, leaving the rest blank
func MakeEnum() []EnumTestT {
	return []EnumTestT{EnumTestA, EnumTestB, EnumTestC, EnumTestD, EnumTestE, EnumTestF, EnumTestG, EnumTestH, EnumTestI, EnumTestJ, EnumTestK}
}

type Command int

const (
	IDENTITY Command = iota // returns value unmodified
	SQUARE                  // squares the value
	HALVE                   // divide by 2
	// ensure the value is in the range 0 to 100. If the value is less than 0 return 0, if the value is greater than 100 return 100
	CLAMP_0_100
	TRANSFORM // if the value is divisible by 3, return the value divided by 3. Otherwise if even return the value multiplied by 2, otherwise return the value unmodified.
	UNUSED_1  // reserved for future use, return value unmodified
	UNUSED_2  // reserved for future use, return value unmodified
	UNUSED_3  // reserved for future use, return value unmodified
)

// TODO: implement the logic from the Command enums according to their comments.  If `cmd` is not a
// value command then follow the IDENTITY command.
// Hint: USE A SWITCH STATEMENT
// Hint: for TRANSFORM try using a "naked" switch
func HandleCommand(cmd Command, value int) int {
	return 0
}

// TODO: respond different based on the type of `input`:
// if input is `int`, return the square root of the number as an int
// if input is `string`, return "Hello ${input}"
// if input is `bool`, return the inverted boolean value
// if input is anything else, return an appropriate error
// Hint: use either type assertions or type switch
func HandleTypes(input any) any {
	return fmt.Errorf("TODO")
}
