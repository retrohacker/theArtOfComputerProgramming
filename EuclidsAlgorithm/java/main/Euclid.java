package main;

public class Euclid {
	
	public static int GCD(int m, int n) {
		//Ensure m is greater than n otherwise swap
		if(m<n){int t=m;m=n;n=t;}
		int r;
		while(true) {
			r = m%n;
			if(r==0) {return n;}
			m = n;
			n = r;
		}
	}
}
