function Euclid(m,n) {
	if(m<=0||n<=0){return -1;}
	var r
	while(true) {
		r=m%n;
		if(r==0) {return n;}
		m = n;
		n = r;
	}
}
