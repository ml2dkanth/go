package main

import ("fmt")

type ml2dkanth int
var x ml2dkanth 

func main() {
	fmt.Println(x)
	fmt.Printf("Type %T\n", x)
	x=42
	fmt.Println(x)
}

