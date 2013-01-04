package euclidext;

/**
 * An implementation of Euclid's Extend Algortihm for finding the GCD.
 *
 * This class is used to find A*M+B*N=D for any input M and N where B is the
 * greatest common denominator. To use this class, instantiate an object of
 * type EuclidsExtended supplying input1 and input2. The constructor sets
 * up the class to run. To run the algorithm call .execute() on the object.
 * To retrieve the results call .getResult().
 */
public class EuclidsExtended {

	public static final String NonPos = "Non-Positive Input";
	private Result variables;

	/**
	 * Instantiate the class.
	 * The result will be of the form A*M+B*N=D. Input1 corresponds to
	 * M and input2 corresponds to N.
	 */
	public EuclidsExtended(int input1, int input2) {
		variables = new Result(input1,input2);
		
		// Ensure the input is valid
		if(input1<=0||input2<=0){
			variables.setError(EuclidsExtended.NonPos);
		}
	}

	/**
	 * This function runs the actual algorithm on the input supplied in the
	 * constructor.
	 */
	public boolean execute() {
		// If the input is invalid, don't bother with processing
		if(variables.errorFlag) {return false;}

		// Initialize Variables
		int aNot=1;
		variables.coeff2=1;
		variables.coeff1=0;
		int bNot=0;
		int c=variables.input1;
		variables.gcd=variables.input2;
		// Main loop	
		while(true) {
			// Is remainder zero?
			int r = c%variables.gcd;
			if(r==0) {return true;}
			int q = c/variables.gcd;
			c=variables.gcd;
			variables.gcd=r;
			int t = aNot;
			aNot = variables.coeff1;
			variables.coeff1 = t-q*variables.coeff1;
			t = bNot;
			bNot = variables.coeff2;
			variables.coeff2 = t-q*variables.coeff2;
		}
	}

	/**
	 * Get the coefficient for input1.
	 * This is used in constructing the form A*M+B*N=D where A is this
	 * returned coefficent and M is input1.
	 *
	 * Call this function after calling execute()
	 */
	public int getCoeff1() {
		return variables.coeff1;
	}

	/**
	 * Get the coefficient for input2.
	 * This is used in constructing the form A*M+B*N=D where B is this
	 * returned coefficent and N is input1.
	 *
	 * Call this function after calling execute()
	 */
	public int getCoeff2() {
		return variables.coeff2;
	}

	/**
	 * Get the greatest common denominator.
	 * This is used in constructing the form A*M+B*N=D where D is this
	 * returned value. The greatest common denominator is the largest
	 * number which evenly devides input1 (M) and input2 (N).
	 *
	 * Call this function after calling execute()
	 */
	public int getGCD() {
		return variables.gcd;
	}

	/**
	 * Get a datastructure containing the result of the algorithm.
	 *
	 * Call this function after calling execute()
	 */
	public Result getResult() {
		return variables;
	}
}
