package main

import ("fmt")

type ml2dkanth int
var x ml2dkanth 
var y int
func main() {
	fmt.Println(x)
	fmt.Printf("Type of x is %T\n", x)
	x=42
	fmt.Println("X is ",x)
	
	y = int(x)
	fmt.Println("Y is ",y);
	fmt.Printf("Type of y is %T\n", y)
}

