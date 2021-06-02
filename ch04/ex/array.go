package main
import ("fmt")

func main(){
	arr  := [5] int  { 3,4,5,6}

	for i,v := range arr{
		fmt.Println(i,v)}
		fmt.Printf("%T\n", arr)
	}
