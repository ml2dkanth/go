package main
import ("fmt")

func main(){
	// arr := type{values} //composite
	// group together VALUES of same type

	arr := [] int {3,4,5,6,43}

	fmt.Println(arr)
	for i, v := range arr {
		fmt.Printf("Index %d  Value %d \n",i, v)
	}

}
