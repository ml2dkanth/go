package main
import ("fmt")

func main(){
// Multiple slide example
veg:=[]string {"Tamato","Potato","Brinjal","Carrot"}
fruits :=[] string{"Apple", "Banana","Orange","Pear"}

fmt.Println("Vegetables : ", veg)
fmt.Println("Fruits: ", fruits)

// MultiDimentional slice

green :=[][] string {veg, fruits}
fmt.Println (green)

}

