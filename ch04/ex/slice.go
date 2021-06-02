package main
import ("fmt")

func main(){
	//Remove [4] to [] to become a slice
	arr  := [] int  { 3,4,5,6}

	for i,v := range arr{
		fmt.Println(i,v)}
		fmt.Printf("%T\n", arr)
	}
