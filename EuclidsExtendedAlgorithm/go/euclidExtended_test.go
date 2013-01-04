package euclidext

import "testing"
import "errors"

// We are using a table driven test
var TestCases = []struct{
	input1 int;
	input2 int;
	output1 int;
	output2 int;
	output3 int;
	err error;
}{
	{1769,551,5,-16,29,nil},
	{551,1769,-16,5,29,nil},
	{-1769,551,0,0,0,errors.New(ErrNonPos)},
	{1769,-551,0,0,0,errors.New(ErrNonPos)},
	{-1769,-551,0,0,0,errors.New(ErrNonPos)},
	{120,23,-9,47,1,nil},
	{23,120,47,-9,1,nil},
	{-120,23,0,0,0,errors.New(ErrNonPos)},
	{120,-23,0,0,0,errors.New(ErrNonPos)},
	{-120,-23,0,0,0,errors.New(ErrNonPos)},
	{57,46,21,-26,1,nil},
	{46,57,-26,21,1,nil},
	{-57,46,0,0,0,errors.New(ErrNonPos)},
	{57,-46,0,0,0,errors.New(ErrNonPos)},
	{-57,-46,0,0,0,errors.New(ErrNonPos)},
}

func TestEuclidsExtendedAlgorithm(t *testing.T) {
	for i,tt := range TestCases {
		result1,result2,result3,err:=EuclidExtended(tt.input1,tt.input2)
		if err==nil&&tt.err!=nil {
			// If the function did not return an error message when it was expected to.
			t.Error("Index ",i, " DIDNT RETURN EXPECTED ERROR. Input:",tt.input1," ",tt.input2," Expected Error:",tt.err," Returned Error:",err,"\n")
		} else if err!=nil&&tt.err==nil {
			// If the function returned an error message when it was NOT expected to.
			t.Error("Index ",i, " RETURNED UNEXPECTED ERROR. Input:",tt.input1," ",tt.input2," Expected Error:",tt.err," Returned Error:",err,"\n")
		} else if result1!=tt.output1||result2!=tt.output2||result3!=result3 {
			// If the function returned unexpected results.
			t.Error("Index ",i," FAILED. Input:",tt.input1," ",tt.input2," Output: ",result1," ",result2," ",result3," ",err,"\n")
		} else if  ( err!=nil&&tt.err!=nil ) && ( err.Error()!=tt.err.Error() ) {
			//If the function returned an error message like expected, but it was the wrong error.
			t.Error("Index ",i," RETURNED WRONG ERROR. Input:",tt.input1," ",tt.input2," Error Returned:",err,"\n")
		}
	}
}
