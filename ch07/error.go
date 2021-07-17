package main

import  "fmt"

func main(){
	
	n, err := fmt.Println("Hello Srikanth")

	if(err!=nil){
		fmt.Println("Error : ",err)
	}
	fmt.Println("Lenght of Println : ",n)

}

