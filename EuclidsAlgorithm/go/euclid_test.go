package main

//Import Libraries
import "testing"

var flagTests = []struct {
	input1 int
	input2 int
	output int
}{
	{20,100,20},
	{100,20,20},
	{8,15,1},
	{15,8,1},
	{21,133,7},
	{133,21,7},
	{36,4352,4},
	{4352,36,4},
}

func TestEuclidsAlgorithm(t *testing.T) {
	for i,tt := range flagTests {
		result:=euclid(tt.input1,tt.input2)
		if(result!=tt.output) {
			t.Error("Index ",i," FAILED. Input:",tt.input1," ",tt.input2,
			" Expected Output:",tt.output," Result:",result)
		}
	}
}
