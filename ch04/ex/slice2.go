package main
import ("fmt")

func main(){
	//Remove [4] to [] to become a slice
	arr  := [] int  { 3,4,5,6,7}

	fmt.Println(arr[0:2])
	fmt.Println(arr[:])

	}
