package euclidext;

/**
 * A structure desinged to support Euclid's Extended Algorithm.
 * The data is in the form A*M+B*N=D where A=coeff1, M=input1, 
 * B=coeff2, N=input2, and D=gcd
 */
class Result {
	public int input1; //M
	public int input2; //N
	public int coeff1; //A
	public int coeff2; //B
	public int gcd; //D
	public boolean errorFlag; //Set to true if an error is encountered
	public String err; //The error that has been encountered

	/**
	 * Initialize Result with input data.
	 * When using this constructor it is assumed the rest of the data will be
	 * filled by an external function at a later point in time. If all the
	 * values for Result are known at the time of instantiation, use the
	 * other constructor.
	 */
	public Result(int input1, int input2) {
		this.input1 = input1;
		this.input2 = input2;
	}

	/**
	 * Initialize Result with all data.
	 * This constructor should be used when all values of the Result are already known.
	 * an example of when to use this type of constructor is when making unit tests.
	 */
	public Result(int input1, int input2, int coeff1, int coeff2, int gcd, boolean errorFlag, String err) {
		this.input1=input1;
		this.input2=input2;
		this.coeff1=coeff1;
		this.coeff2=coeff2;
		this.gcd=gcd;
		this.errorFlag=errorFlag;
		this.err=err;
	}

	/**
	 * An error has been encountered.
	 * This will set all generated values to 0 (coeff1, coeff2, and gcd) as
	 * well as set the errorFlag to true. The generated error message should
	 * be passed in as a parameter.
	 */
	public void setError(String err) {
		this.coeff1=0;
		this.coeff2=0;
		this.gcd=0;
		this.errorFlag=true;
		this.err=err;
	}

	@Override
	public boolean equals(Object obj) {
		if(!(obj instanceof Result)) {
			return false;
		}

		Result that = (Result) obj;

		if(
			this.coeff1==that.coeff1 &&
			this.coeff2==that.coeff2 &&
			this.gcd==that.gcd &&
			this.errorFlag==that.errorFlag &&
			this.err==that.err
		) {return true;}
		return false;
	}

	@Override
	public String toString() {
		if(errorFlag) {
			return "Error: "+err;
		}
		return coeff1+"*"+input1+"+"+coeff2+"*"+input2+"="+gcd;
	}	
}
