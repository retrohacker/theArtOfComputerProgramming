/**
 * Euclid's algorithm can be used to find the greatest commmon denominator
 * between to input values. This function will return 0 if either input is
 * 0 or -1 if either input is negative.
 */
function Euclid(input1,input2) {
	if(input1==0||input2==0){return 0;}
	if(input1<0||input2<0){return -1;}
	while(true) {
		remainder=input1%input2;
		if(remainder==0) {return input2;}
		input1 = input2;
		input2 = remainder;
	}
}
