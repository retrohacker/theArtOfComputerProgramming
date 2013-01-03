function Euclid(m,n) {
	if(m<n){t=m;m=n;n=t;}

	var r
	while(true) {
		r=m%n;
		if(r==0) {return n;}
		m = n;
		n = r;
	}
}
