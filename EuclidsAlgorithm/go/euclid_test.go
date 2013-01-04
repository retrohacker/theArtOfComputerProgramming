package euclid

import "testing"

// A table based approach to unit testing
var testCases = []struct {
	input1 uint
	input2 uint
	output uint
}{
	{20,100,20},
	{100,20,20},
	{8,15,1},
	{15,8,1},
	{0,15,0},
	{8,0,0},
	{21,133,7},
	{133,21,7},
	{36,4352,4},
	{4352,36,4},
}

//A table based test for Euclid's algorithm.
func TestEuclidsAlgorithm(t *testing.T) {
	for i,testCase := range testCases {
		result:=euclid(testCase.input1,testCase.input2)
		if(result!=testCase.output) {
			t.Error("Index ",i," FAILED. Input:",testCase.input1," ",testCase.input2,
			" Expected Output:",testCase.output," Result:",result)
		}
	}
}
