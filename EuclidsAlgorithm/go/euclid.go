package main

func euclid(m,n uint) (uint) {
	var r uint
	//Make sure m > n, if not swap
	if(n>m) {r=m;m=n;n=r}
	if(m==0||n==0) {return 0}
	for m!=n {
		r = m%n
		if(r==0) {break}
		m = n
		n = r
	}
	return n
}
