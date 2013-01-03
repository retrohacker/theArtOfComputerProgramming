package main

func euclid(m,n uint) (uint) {
	var r uint
	if(m==0||n==0) {return 0}
	for {
		r = m%n
		if(r==0) {break}
		m = n
		n = r
	}
	return n
}
