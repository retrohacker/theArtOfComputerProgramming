package main

import "flag"

func main() {
	maxVal:=flag.Int("c",1,"Count. This argument is used to determine the maximum value you want to search to for prime numbers. Do not confuse this with the number of primes found")
	newLine:=flag.Bool("n",false,"New Line. Should there be line breaks between each number?")
	flag.Parse()
	SieveOfEratosthenes(*maxVal,*newLine)
}
