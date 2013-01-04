package euclid;

public class Euclid {
	
	/**
	 * A function for finding the greatest common denominator between
	 * two inputs.
	 *
	 * @return
	 * Returns the Greatest Common Divisor or -1 if the input is invalid
	 * (invalid inputs are any value less than 0). This function returns
	 * 0 if either input is equal to 0.
	 */
	public static int GCD(int input1, int input2) {
		if(input1==0||input2==0) {return 0;}
		if(input1<0||input2<0) {return -1;}

		int remainder;
		while((remainder=input1%input2)!=0) {
			input1 = input2;
			input2 = remainder;
		} 

		return input2;
	}
}
