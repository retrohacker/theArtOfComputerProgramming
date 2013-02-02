package main

import "fmt"

func main() {
	smallest:=9999;
	sm:=1;
	sn:=1;
	largest:=0;
	lm:=1;
	ln:=1;
	for m:=1;m<=10;m++ {
		for n:=1;n<m;n++ {
			_,count:=euclid(uint(m),uint(n))
			if(count<smallest) {smallest=count;sm=m;sn=n;}
			if(count>largest) {largest=count;lm=m;ln=n;}
		}
	}
	fmt.Println("Largest: ",largest," [",lm,",",ln,"] Smallest: ",smallest,"[",sm,",",sn,"]");
}
