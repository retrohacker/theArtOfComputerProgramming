/*
	Package euclid allows for finding the greatest common denominator
	between two numbers.
*/
package euclid

// Function euclid finds the greates common denominator between two unsigned
// integers m and n. If either value is 0 the result will be 0.
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
