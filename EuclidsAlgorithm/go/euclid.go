package main

func euclid(m,n int) (int) {
	var r int
	//Make sure m > n, if not swap
	if(n>m) {r=m;m=n;n=r}
	for m!=n {
		r = m%n
		if(r==0) {break}
		m = n
		n = r
	}
	return n
}
