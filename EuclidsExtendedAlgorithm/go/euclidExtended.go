/*
	Package euclidext contains a simple implementation of Euclid's Extended
	Algorithm for finding a functions greatest common denominator. It is
	capable of solving the equation A*M+B*N=D for any input M and N where D
	is the greatest common denominator.
*/
package euclidext

import "errors"

// Used as an error message indicating EuclidExtended was supplied with an
// input outside the range the function can handle. 
const ErrNonPos = "Non Positive Input";

// Solves the equation A*M+B*N=D for any input M and N where D is the greatest
// common denominator. Returns A, B, D, and an error in the event of invalid
// input. If an error is encountered A, B, and D will be 0.
func EuclidExtended(m,n int) (a,b,d int, e error) {
	//Check for invalid input
	if(m<0||n<0) {return 0,0,0,errors.New(ErrNonPos)};

	//Initialize all variables
	a,b,d,e=0,1,n,error(nil)
	aNot,bNot,c:=1,0,m
	var r,q int;

	for {
		//Remainder zero?
		r=c%d;
		if(r==0) {break}
		q=c/d;

		//Swap values
		c,d=d,r
		temp:=aNot
		aNot,a=a,temp-q*a
		temp=bNot
		bNot,b=b,temp-q*b
	}
	return
}
