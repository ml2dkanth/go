package main

import ("fmt"
	"log"
)

type myError struct {
	str string
	err error
}
func main(){

	var name string
	var count int

	fmt.Print("Enter your Name :")
	fmt.Scan(&name)
	fmt.Print("Number :")
	fmt.Scan(&count)

	_, err := myFunc(count)

	if err != nil{
		log.Println(err)
	}
}

func (merr myError) Error() string {
	return fmt.Sprintf("CustomError: %v %v",merr.str, merr.err)
}

func myFunc(ct int) (int, error){

	if ct==2  {
		str := fmt.Errorf("Error GIVEN: %v", ct)
		return 0, myError{" SRIKERR ",str}
	}
	return 33,nil
}


