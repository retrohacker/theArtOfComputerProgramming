package main

func euclid(m,n int) int {
	var r int
	for m!=n {
		r = m%n
		if(r==0) {break}
		m = n
		n = r
	}
	return n
}
