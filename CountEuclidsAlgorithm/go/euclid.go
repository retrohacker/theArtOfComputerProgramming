/*
	Package euclid allows for finding the greatest common denominator
	between two numbers.
*/
package main

// Function euclid finds the greates common denominator between two unsigned
// If either value is 0 the result will be 0.
func euclid(input1,input2 uint) (uint,int) {
	var remainder uint
	count:=0
	if(input1==0||input2==0) {return 0,count}
	for {
		count++
		remainder = input1%input2
		if(remainder==0) {break}
		input1 = input2
		input2 = remainder
	}
	return input2,count
}
