package euclidext;

/**
 * A structure desinged to support Euclid's Extended Algorithm.
 */
class Result {
	public int input1;
	public int input2;
	public int coeff1;
	public int coeff2;
	public int gcd;	
	public boolean errorFlag;
	public String err;

	public Result(int input1, int input2) {
		this.input1 = input1;
		this.input2 = input2;
		this.errorFlag = false;
		this.err = null;
	}

	public Result(int input1, int input2, int coeff1, int coeff2, int gcd, boolean errorFlag, String err) {
		this.input1=input1;
		this.input2=input2;
		this.coeff1=coeff1;
		this.coeff2=coeff2;
		this.gcd=gcd;
		this.errorFlag=errorFlag;
		this.err=err;
	}

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
