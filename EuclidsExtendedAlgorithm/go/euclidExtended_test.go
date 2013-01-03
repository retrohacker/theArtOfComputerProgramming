package main

import "testing"
import "errors"

var TestCases = []struct{
	input1 int;
	input2 int;
	output1 int;
	output2 int;
	output3 int;
	err error;
}{
	{1769,551,5,-16,29,nil},
	{-1769,551,0,0,0,errors.New(NonPos)},
	{1769,-551,0,0,0,errors.New(NonPos)},
	{-1769,-551,0,0,0,errors.New(NonPos)},
}

func TestEuclidsExtendedAlgorithm(t *testing.T) {
	for i,tt := range TestCases {
		result1,result2,result3,err:=EuclidExtended(tt.input1,tt.input2)
		if result1!=tt.output1||result2!=tt.output2||result3!=result3 {
			t.Error("Index ",i," FAILED. Input:",tt.input1," ",tt.input2," Output: ",result1," ",result2," ",result3," ",err,"\n")
		} else if  ( err!=nil&&tt.err!=nil ) && ( err.Error()!=tt.err.Error() ) {
			t.Error("Index ",i," RETURNED WRONG ERROR. Input:",tt.input1," ",tt.input2," Error Returned:",err)
		} else if err==nil&&tt.err!=nil {
			t.Error("Index ",i, " DIDNT RETURN EXPECTED ERROR. Input:",tt.input1," ",tt.input2," Expected Error:",tt.err," Returned Error:",err)
		} else if err!=nil&&tt.err==nil {
			t.Error("Index ",i, " RETURNED UNEXPECTED ERROR. Input:",tt.input1," ",tt.input2," Expected Error:",tt.err," Returned Error:",err)
		}
	}
}

