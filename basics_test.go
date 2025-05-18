package main

import (
	"fmt"
	"strings"
	"testing"
)

func mapEq[T1, T2 comparable](m1 map[T1]T2, m2 map[T1]T2) bool {
	if len(m1) != len(m2) {
		return false
	}
	for k, v := range m1 {
		if w, ok := m2[k]; !ok || v != w {
			return false
		}
	}
	return true
}

func sliceEq[T comparable](s1 []T, s2 []T) bool {
	if len(s1) != len(s2) {
		return false
	}
	for i := 0; i < len(s1); i++ {
		if s1[i] != s2[i] {
			return false
		}
	}
	return true
}

func TestHello(t *testing.T) {
	const expectedPrefix = "Hi, my name is "
	res := Hello()
	if strings.HasPrefix(res, expectedPrefix) == false {
		t.Fatalf("Does not start with prefix \"%s\"", expectedPrefix)
	}
	rest := strings.TrimPrefix(res, expectedPrefix)
	if len(rest) == 0 {
		t.Fatalf("Expected something else after \"%s\"", expectedPrefix)
	}
	fmt.Printf("Hey there %s!\n", rest)
}

func TestZeroInt(t *testing.T) {
	if ZeroInt() != 0 {
		t.Fatalf("Incorrect integer zero value")
	}
}

func TestZeroFloat(t *testing.T) {
	if ZeroFloat() != 0.0 {
		t.Fatalf("Incorrect float zero value")
	}
}

func TestZeroBool(t *testing.T) {
	if ZeroBool() != false {
		t.Fatalf("Incorrect boolean zero value")
	}
}

func TestZeroString(t *testing.T) {
	if ZeroString() != "" {
		t.Fatalf("Incorrect string zero value")
	}
}

func TestIntToString(t *testing.T) {
	if IntToString(123) != "123" {
		t.Fatalf("Failed to convert 123 to string")
	}
	if IntToString(-54321) != "-54321" {
		t.Fatalf("Failed to convert -54321 to string")
	}
	if IntToString(0) != "0" {
		t.Fatalf("Failed to convert 0 to string")
	}
}

func TestStringToInt(t *testing.T) {
	{
		v, err := StringToInt("123")
		if err != nil {
			t.Fatalf("error converting 123 to int - %s", err.Error())
		}
		if v != 123 {
			t.Fatalf("incorrect value, expected 123, got %d", v)
		}
	}
	{
		v, err := StringToInt("-7832")
		if err != nil {
			t.Fatalf("error converting -7832 to int - %s", err.Error())
		}
		if v != -7832 {
			t.Fatalf("incorrect value, expected -7832, got %d", v)
		}
	}
	{
		v, err := StringToInt("0")
		if err != nil {
			t.Fatalf("error converting 0 to int - %s", err.Error())
		}
		if v != 0 {
			t.Fatalf("incorrect value, expected 0, got %d", v)
		}
	}
}

func TestMultipleReturn(t *testing.T) {
	if n1, n2 := MultipleReturn(42); n1 != 21 || n2 != 1_764 {
		t.Fatalf("Incorrect values for 42")
	}
	if n1, n2 := MultipleReturn(41); n1 != 40 || n2 != 6 {
		t.Fatalf("Incorrect values for 41")
	}
}

func TestMakeMap(t *testing.T) {
	m := MakeMap()
	if len(m) != 2 {
		t.Fatalf("Map needs to have exactly two key-value pairs")
	}
	if elem, present := m["sagan"]; !present || elem != 30023 {
		t.Fatalf("Map does not have right 'carl' key-value")
	}
	if elem, present := m["mendeleev"]; !present || elem != 101 {
		t.Fatalf("Map does not have right 'mendeleev' key-value")
	}
}

func TestMapClear(t *testing.T) {
	m := map[uint64]any{
		123:    "foobar",
		0xfcde: 3.1415,
		999:    "baz",
	}
	MapClear(m)
	if len(m) != 0 {
		t.Fatalf("Map is not cleared. Still has %d elements", len(m))
	}
}

func TestMapManipulate(t *testing.T) {
	m := map[int]string{}
	mapExpected := map[int]string{}
	MapManipulate(m)
	if !mapEq(m, mapExpected) {
		t.Fatalf("Expected empty map")
	}

	m[2000] = "asdf"
	mapExpected = map[int]string{
		2000: "asdf",
		4000: "XYZ",
	}
	MapManipulate(m)
	if !mapEq(m, mapExpected) {
		t.Fatalf("Incorrect map")
	}

	m[100] = "delete me"
	MapManipulate(m)
	if !mapEq(m, mapExpected) {
		t.Fatalf("Incorrect map")
	}
}

func TestForLoop(t *testing.T) {
	ans := ForLoop()
	if ans != (1249875*1000)+((1<<3)*9*5*5*5*5*5*5*11*101) {
		t.Fatalf("Incorrect answer")
	}
}

func TestWhileLoop(t *testing.T) {
	ans := WhileLoop()
	if ans != 3*43 {
		t.Fatalf("Incorrect answer")
	}
}

func TestForRangeArray(t *testing.T) {
	if ForRangeArray([]int{}) != 0 {
		t.Fatalf("Input [] expected 0")
	}
	if ForRangeArray([]int{1, 1, 1}) != 3 {
		t.Fatalf("Input [1,1,1] expected 3")
	}
	if ForRangeArray([]int{7, 7, 7}) != 24 {
		t.Fatalf("Input [7,7,7] expected 24")
	}
}

func TestForRangeMap(t *testing.T) {
	{
		m := map[string]int{}
		if _, err := ForRangeMap(m); err == nil {
			t.Fatalf("Expected error")
		}
	}
	{
		m := map[string]int{
			"can": 100,
		}
		if _, err := ForRangeMap(m); err == nil {
			t.Fatalf("Expected error")
		}
	}
	{
		m := map[string]int{
			"can":  100,
			"bean": 200,
		}
		val, err := ForRangeMap(m)
		if err != nil {
			t.Fatalf("Expected error to be nil")
		}
		if val != 200 {
			t.Fatalf("Expected minimum of 200")
		}
	}
	{
		m := map[string]int{
			"can":    100,
			"bean":   200,
			"become": 50,
		}
		val, err := ForRangeMap(m)
		if err != nil {
			t.Fatalf("Expected error to be nil")
		}
		if val != 50 {
			t.Fatalf("Expected minimum of 50")
		}
	}
}

func TestCreateArray(t *testing.T) {
	{
		res := CreateArray(12345)
		if res[0] != 123450 {
			t.Fatalf("Wrong answer for zeroth element")
		}
		if res[1] != "12345" {
			t.Fatalf("Wrong answer for first element")
		}
		if res[2] != 57 {
			t.Fatalf("Wrong answer for second element")
		}
	}
	{
		res := CreateArray(5468672)
		if res[0] != 54686720 {
			t.Fatalf("Wrong answer for zeroth element")
		}
		if res[1] != "5468672" {
			t.Fatalf("Wrong answer for first element")
		}
		if res[2] != 0 {
			t.Fatalf("Wrong answer for second element")
		}
	}
}

func TestCreateSlice(t *testing.T) {
	{
		res := CreateSlice(579)
		if !sliceEq(res, []int{579, 289, 144, 72, 36, 18, 9, 4, 2, 1}) {
			t.Fatalf("Unexpected result for 579, got %#v", res)
		}
	}
	{
		res := CreateSlice(-10)
		if !sliceEq(res, []int{}) {
			t.Fatalf("Unexpected result for -10, got %#v", res)
		}
	}
	{
		res := CreateSlice(5789102934)
		if !sliceEq(res, []int{5789102934, 2894551467, 1447275733, 723637866, 361818933, 180909466, 90454733, 45227366, 22613683, 11306841, 5653420, 2826710, 1413355, 706677, 353338, 176669, 88334, 44167, 22083, 11041, 5520, 2760, 1380, 690, 345, 172, 86, 43, 21, 10, 5, 2, 1}) {
			t.Fatalf("Unexpected result for 5789102934, got %#v", res)
		}
	}
}

func TestMakeSlice(t *testing.T) {
	test := func(s []string, expectedLength int) {
		if len(s) != expectedLength {
			t.Fatalf("Wrong length")
		}
		for _, elem := range s {
			if elem != "" {
				t.Fatalf("String is not empty string")
			}
		}
	}
	test(MakeSlice(10), 10)
	test(MakeSlice(100), 100)
	test(MakeSlice(1000), 1000)
}

func TestSliceMerge(t *testing.T) {
	{
		res := SliceMerge([]int{}, []int{})
		if !sliceEq(res, []int{}) {
			t.Fatalf("Unexpected result for [], [] -> %#v", res)
		}
	}
	{
		res := SliceMerge([]int{5}, []int{})
		if !sliceEq(res, []int{5}) {
			t.Fatalf("Unexpected result for [5], [] -> %#v", res)
		}
	}
	{
		res := SliceMerge([]int{5}, []int{1, 10})
		if !sliceEq(res, []int{1, 5, 10}) {
			t.Fatalf("Unexpected result for [5], [1, 10] -> %#v", res)
		}
	}
	{
		res := SliceMerge([]int{5}, []int{1, 10})
		if !sliceEq(res, []int{1, 5, 10}) {
			t.Fatalf("Unexpected result for [5], [1, 10] -> %#v", res)
		}
	}
	{
		res := SliceMerge([]int{6, 10, 20, 50}, []int{-20, 2, 3, 4, 44, 60})
		if !sliceEq(res, []int{-20, 2, 3, 4, 6, 10, 20, 44, 50, 60}) {
			t.Fatalf("Unexpected result for [5], [1, 10] -> %#v", res)
		}
	}
}

func TestInitStruct(t *testing.T) {
	s1, s2 := InitStruct()
	{
		if s1.A != "" {
			t.Fatalf("Unexpected value")
		}
		if s1.B != 0 {
			t.Fatalf("Unexpected value")
		}
		if s1.C != 0 {
			t.Fatalf("Unexpected value")
		}
		if s1.D != 0 {
			t.Fatalf("Unexpected value")
		}
		if s1.E != 0 {
			t.Fatalf("Unexpected value")
		}
		if s1.F != "" {
			t.Fatalf("Unexpected value")
		}
		if s1.G != false {
			t.Fatalf("Unexpected value")
		}
		if s1.H != "" {
			t.Fatalf("Unexpected value")
		}
		if s1.I != 0 {
			t.Fatalf("Unexpected value")
		}
		if s1.J != 0.0 {
			t.Fatalf("Unexpected value")
		}
		if s1.K != "" {
			t.Fatalf("Unexpected value")
		}
		if s1.L != "" {
			t.Fatalf("Unexpected value")
		}
		if s1.M != 0 {
			t.Fatalf("Unexpected value")
		}
		if s1.N != 0.0 {
			t.Fatalf("Unexpected value")
		}
		if s1.O != 0 {
			t.Fatalf("Unexpected value")
		}
		if s1.P != "" {
			t.Fatalf("Unexpected value")
		}
		if s1.Q != "" {
			t.Fatalf("Unexpected value")
		}
		if s1.R != 0 {
			t.Fatalf("Unexpected value")
		}
		if s1.S != false {
			t.Fatalf("Unexpected value")
		}
		if s1.T != 0 {
			t.Fatalf("Unexpected value")
		}
		if s1.U != "" {
			t.Fatalf("Unexpected value")
		}
		if s1.V != 0 {
			t.Fatalf("Unexpected value")
		}
		if s1.W != 0 {
			t.Fatalf("Unexpected value")
		}
		if s1.X != 0.0 {
			t.Fatalf("Unexpected value")
		}
		if s1.Y != 0 {
			t.Fatalf("Unexpected value")
		}
		if s1.Z != "" {
			t.Fatalf("Unexpected value")
		}
	}
	{
		if s2.A != "" {
			t.Fatalf("Unexpected value")
		}
		if s2.B != 0 {
			t.Fatalf("Unexpected value")
		}
		if s2.C != 0 {
			t.Fatalf("Unexpected value")
		}
		if s2.D != 42 {
			t.Fatalf("Unexpected value")
		}
		if s2.E != 0 {
			t.Fatalf("Unexpected value")
		}
		if s2.F != "" {
			t.Fatalf("Unexpected value")
		}
		if s2.G != false {
			t.Fatalf("Unexpected value")
		}
		if s2.H != "" {
			t.Fatalf("Unexpected value")
		}
		if s2.I != 0 {
			t.Fatalf("Unexpected value")
		}
		if s2.J != 0.0 {
			t.Fatalf("Unexpected value")
		}
		if s2.K != "" {
			t.Fatalf("Unexpected value")
		}
		if s2.L != "" {
			t.Fatalf("Unexpected value")
		}
		if s2.M != 0 {
			t.Fatalf("Unexpected value")
		}
		if s2.N != 0.0 {
			t.Fatalf("Unexpected value")
		}
		if s2.O != 0 {
			t.Fatalf("Unexpected value")
		}
		if s2.P != "" {
			t.Fatalf("Unexpected value")
		}
		if s2.Q != "qqq" {
			t.Fatalf("Unexpected value")
		}
		if s2.R != 0 {
			t.Fatalf("Unexpected value")
		}
		if s2.S != false {
			t.Fatalf("Unexpected value")
		}
		if s2.T != 0 {
			t.Fatalf("Unexpected value")
		}
		if s2.U != "" {
			t.Fatalf("Unexpected value")
		}
		if s2.V != 0 {
			t.Fatalf("Unexpected value")
		}
		if s2.W != 0 {
			t.Fatalf("Unexpected value")
		}
		if s2.X != 3.1415 {
			t.Fatalf("Unexpected value")
		}
		if s2.Y != 0 {
			t.Fatalf("Unexpected value")
		}
		if s2.Z != "" {
			t.Fatalf("Unexpected value")
		}
	}
}

func TestInitTree(t *testing.T) {
	tree := InitTree()
	if tree.Val != 100 {
		t.Fatalf("tree head doesn't have value 100")
	}
	if tree.Left == nil {
		t.Fatalf("tree head doesn't have left branch")
	}
	if tree.Right == nil {
		t.Fatalf("tree head doesn't have right branch")
	}
	if tree.Left.Val != "W" {
		t.Fatal("left branch doesn't have value \"W\"")
	}
	if tree.Left.Left != nil {
		t.Fatalf("left branch has extra node")
	}
	if tree.Left.Right != nil {
		t.Fatalf("left branch has extra node")
	}
	if tree.Right.Val != 8.7 {
		t.Fatal("right branch doesn't have value 8.7")
	}
	if tree.Right.Left != nil {
		t.Fatalf("right branch has extra node")
	}
	if tree.Right.Right != nil {
		t.Fatalf("right branch has extra node")
	}
}

func TestMakeEnum(t *testing.T) {
	enums := MakeEnum()
	if len(enums) != 11 {
		t.Fatalf("Expected 11 enum values")
	}
	if enums[0] != 1 {
		t.Fatalf("Expected enum 0 to be 1")
	}
	if enums[1] != 2 {
		t.Fatalf("Expected enum 1 to be 2")
	}
	if enums[2] != 4 {
		t.Fatalf("Expected enum 2 to be 4")
	}
	if enums[3] != 8 {
		t.Fatalf("Expected enum 3 to be 8")
	}
	if enums[4] != 16 {
		t.Fatalf("Expected enum 4 to be 16")
	}
	if enums[5] != 32 {
		t.Fatalf("Expected enum 5 to be 32")
	}
	if enums[6] != 64 {
		t.Fatalf("Expected enum 6 to be 64")
	}
	if enums[7] != 128 {
		t.Fatalf("Expected enum 7 to be 128")
	}
	if enums[8] != 256 {
		t.Fatalf("Expected enum 8 to be 256")
	}
	if enums[9] != 512 {
		t.Fatalf("Expected enum 9 to be 512")
	}
	if enums[10] != 1024 {
		t.Fatalf("Expected enum 10 to be 1024")
	}
}

func TestHandleCommand(t *testing.T) {
	if HandleCommand(IDENTITY, 10) != 10 {
		t.Fatalf("command IDENTITY 10 expected 10")
	}
	if HandleCommand(UNUSED_1, 10) != 10 {
		t.Fatalf("command UNUSED_1 10 expected 10")
	}
	if HandleCommand(UNUSED_2, 10) != 10 {
		t.Fatalf("command UNUSED_2 10 expected 10")
	}
	if HandleCommand(UNUSED_3, 10) != 10 {
		t.Fatalf("command UNUSED_3 10 expected 10")
	}
	if HandleCommand(SQUARE, 25) != 625 {
		t.Fatalf("command SQUARE 25 expected 625")
	}
	if HandleCommand(SQUARE, -8) != 64 {
		t.Fatalf("command SQUARE -8 expected 64")
	}
	if HandleCommand(HALVE, 30) != 15 {
		t.Fatalf("command HALVE 30 expected 15")
	}
	if HandleCommand(HALVE, -10) != -5 {
		t.Fatalf("command HALVE -10 expected -5")
	}
	if HandleCommand(CLAMP_0_100, 15) != 15 {
		t.Fatalf("command CLAMP_0_100 15 expected 15")
	}
	if HandleCommand(CLAMP_0_100, -50) != 0 {
		t.Fatalf("command CLAMP_0_100 -50 expected 0")
	}
	if HandleCommand(CLAMP_0_100, 999) != 100 {
		t.Fatalf("command CLAMP_0_100 999 expected 100")
	}
	if HandleCommand(TRANSFORM, 15) != 5 {
		t.Fatalf("command TRANSFORM 15 expected 5")
	}
	if HandleCommand(TRANSFORM, 62) != 124 {
		t.Fatalf("command TRANSFORM 62 expected 124")
	}
	if HandleCommand(TRANSFORM, 65) != 65 {
		t.Fatalf("command TRANSFORM 65 expected 65")
	}
}

func TestHandleTypes(t *testing.T) {
	if HandleTypes(64) != 8 {
		t.Fatalf("Expected integer 64 -> 8")
	}
	if HandleTypes(99) != 9 {
		t.Fatalf("Expected integer 99 -> 9")
	}
	if HandleTypes("Hanna") != "Hello Hanna" {
		t.Fatalf("Expected string Hanna -> Hello Hanna")
	}
	if HandleTypes(false) != true {
		t.Fatalf("Expected bool false -> true")
	}
	if HandleTypes(true) != false {
		t.Fatalf("Expected bool true -> false")
	}
	if _, ok := HandleTypes(10 + 5i).(error); !ok {
		t.Fatalf("Expected complex -> error")
	}
}
