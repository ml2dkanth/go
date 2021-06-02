package main
import ("fmt")

func main(){
	//map example
	myMap := map [string] int{
		"Door1": 301,
		"Door2": 401}

		fmt.Println("Print Map ",myMap)

		//access a element in map
		fmt.Println("Element in map ", myMap["Door1"])

		//access a element not in map
		fmt.Println("Element NOT in map ", myMap["Door3"])

		//Check exists with condition
		v, ok :=myMap["Door3"]
		fmt.Println(v)
		fmt.Println(ok)

		if v,ok := myMap["Door3"]; ok{
			fmt.Println("NOT Present in map ", v)
		}
		if v,ok := myMap["Door1"]; ok{
			fmt.Println("Present in map ", v)
		}

		//addition of element
		myMap["Door5"] = 501
		fmt.Println("Addition in map :", myMap)

		//Deletion of element
		delete(myMap,"Door1")
		fmt.Println("After Deletion", myMap)
}

