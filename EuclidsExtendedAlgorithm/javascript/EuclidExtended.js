/**
 * Euclid's Extended algorithm can be used to solve A*M+B*N=D for any input
 * M and N where D is the greatest common denominator. The returned value
 * will be in the form Number[3] with [A,B,D]. If either input is 0 then
 * the result will be [0,0,0]. If either input is negative, the result is
 * [0,0,-1]. If the input is either [negative,0] or [0,negative] the
 * result will be [0,0,0].
 */
function EuclidExtended(input1,input2) {
	//Check for 0 or Negative input
	if(input1==0||input2==0) {
		return [0,0,0];
	}
	if(input1<0||input2<0) {
		return [0,0,-1]
	}

	coeff1=0;
	conot1=1;
	coeff2=1;
	conot2=0;
	gcdnot=input1;
	gcd=input2;

	while(true) {
		//Is the remainder 0?
		remainder=gcdnot%gcd;	
		if(remainder==0) {break}
		quotient=(gcdnot-remainder)/gcd;

		//Swap values
		gcdnot=gcd;
		gcd=remainder;
		temp=conot1;
		conot1=coeff1;
		coeff1=temp-(coeff1*quotient);
		temp=conot2;
		conot2=coeff2;
		coeff2=temp-(coeff2*quotient);
	}
	return [coeff1,coeff2,gcd]; 
}
