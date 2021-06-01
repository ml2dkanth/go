package main
import ("fmt")

func main(){
	// arr := type{values} //composite
	// group together VALUES of same type

	arr := [] int {0,1,2,3,4,5}
	fmt.Println("Arr ",arr)
	
	arr2 := [] int {11,12}
	fmt.Println("Arr2 ",arr2)
	
	fmt.Println("Append ",append(arr,arr2...))
	//fmt.Println("Append3 " , append(,arr2...))
	//fmt.Println("Append4 " , append(,...))

}
