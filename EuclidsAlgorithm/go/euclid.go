/*
	Package euclid allows for finding the greatest common denominator
	between two numbers.
*/
package euclid

// Function euclid finds the greates common denominator between two unsigned
// If either value is 0 the result will be 0.
func euclid(input1,input2 uint) (uint) {
	var remainder uint
	if(input1==0||input2==0) {return 0}
	for {
		remainder = input1%input2
		if(remainder==0) {break}
		input1 = input2
		input2 = remainder
	}
	return input2
}
