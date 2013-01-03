package main

//Returns:
//a=Coefficient of m
//b=Coefficient of n
//d=Greatest Common Devisor
//such that a*m+b*n=d
//e=if invalid input is supplied
func EuclidExtended(m,n int) (a,b,d int, e error) {
	//Initialize all variables
	a,b,d,e=0,1,n,error(nil)
	aNot,bNot,c:=1,0,m
	var r,q int;

	for {
		//Remainder zero?
		r=c%d;
		if(r==0) {break}
		q=c/d;

		//Recycle
		c=d
		d=r
		t:=aNot
		aNot=a
		a=t-q*a
		t=bNot
		bNot=b
		b=t-q*b
	}
	return
}
