package main
import ("fmt")

func main(){
	// arr := type{values} //composite
	// group together VALUES of same type

	arr := [] int {0,1,2,3,4,5}

	fmt.Println(arr)
	fmt.Println("arr len ", len(arr))
	
	for i, v := range arr {
		fmt.Print("Index ", i)
		fmt.Println("  Value ",v)
	}
	
	fmt.Println(arr [:])
	fmt.Println(arr[2:4])

	arr2 := append(arr,8,9)
	fmt.Println("Append " , arr2)
}
