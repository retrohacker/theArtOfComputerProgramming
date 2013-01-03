package main

import "testing"

var TestCases = []struct{
	input1 int;
	input2 int;
	output1 int;
	output2 int;
	output3 int;
	errorMsg error;
}{
	{1769,551,5,-16,29,nil},
}

func TestEuclidsExtendedAlgorithm(t *testing.T) {
	for i,tt := range TestCases {
		result1,result2,result3,errorMsg:=EuclidExtended(tt.input1,tt.input2)
		if(result1!=tt.output1||result2!=tt.output2||result3!=result3||errorMsg!=tt.errorMsg) {
			t.Error("Index ",i," FAILED. Input:",tt.input1," ",tt.input2," Output: ",result1," ",result2," ",result3," ",errorMsg,"\n")
		}
	}
}

