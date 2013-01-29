package main

import "fmt"

func SieveOfEratosthenes(maxValue int, newLine bool) {
	//Keep track of how many primes we have found
	primes:=int(0)

	//We make a sieve starting at 2 and going to maxValue
	//To take advantage of the bool's zero value, false will mean that
	//a number could potentially be prime and true means that a number has
	//been ruled out from being prime.
	sieve:=make([]bool,maxValue-1)

	for i,NotPrime := range sieve {
		if !NotPrime {
			index:=i+2;
			output(index,newLine)
			primes++
			for j:=2*index-2;j<len(sieve);j+=index {
				sieve[j]=true; //sieve[j] is not prime
			}
		}
	}

	fmt.Print("\n")
}

func output(i int, newLine bool) {
	//Handle output
	if(!newLine) {
		if(i!=2){fmt.Print(" ")}
	} else {
		if(i!=2){fmt.Print("\n")}
	}
	fmt.Print(i)
}
