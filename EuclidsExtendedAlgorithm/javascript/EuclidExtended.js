function EuclidExtended(input1,input2) {
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
