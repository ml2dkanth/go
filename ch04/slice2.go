package main
import ("fmt")

func main(){
	// arr := type{values} //composite
	// group together VALUES of same type

	arr := [] int {3,4,5,6,43}

	fmt.Println(arr)
	for i, v := range arr {
		fmt.Print("Index ", i)
		fmt.Println("  Value ",v)
	}

}
