package main

import ("fmt")

func main(){
	var i,j int=0,0
	var start,end int=65,90
	for i=start;i<=end;i++ {
		fmt.Println("Loop \t",i)
		for j=1;j<=2;j++ {
		fmt.Printf("Unicode %#U\n",i)
		fmt.Printf("Binnary %b\n",i)
		fmt.Println("")
	}
	}
}


