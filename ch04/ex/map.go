package main

import ("fmt")

func main(){

	myMap := map[string] []string {
		"Veg" : [] string {"Potato", "Tamoto","Brinjal", "Carror"},
		"Fruits" : [] string {"Apple", "Banana","Orange", "Pear"},
		"Nuts" : [] string { "NN", "OO", "BB"},
	}
	fmt.Println(myMap)

	//Addition of element
	myMap["Cars"] = [] string {"Elantra", "Santro", "Creta"}

	fmt.Println(myMap)

	//Deletion of element
	delete(myMap,"Nuts")

	fmt.Println(myMap)

}


