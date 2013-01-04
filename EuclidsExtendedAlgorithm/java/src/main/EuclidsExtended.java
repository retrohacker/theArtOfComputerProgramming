package main;

public class EuclidsExtended {

	public static final String NonPos = "Non-Positive Input";
	private Result variables;

	public EuclidsExtended(int input1, int input2) {
		variables = new Result(input1,input2);
		if(input1<=0||input2<=0){
			variables.setError(EuclidsExtended.NonPos);
		}
	}

	public boolean execute() {
		if(variables.errorFlag) {return false;}
		//Initialize Variables
		int aNot=1;
		variables.coeff2=1;
		variables.coeff1=0;
		int bNot=0;
		int c=variables.input1;
		variables.gcd=variables.input2;
		//Main loop	
		while(true) {
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

	public int getCoeff1() {
		return variables.coeff1;
	}

	public int getCoeff2() {
		return variables.coeff2;
	}

	public int getGCD() {
		return variables.gcd;
	}

	public Result getResult() {
		return variables;
	}

}
