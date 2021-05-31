package main
import ("fmt");

func main(){
	switch {
	case false:
		fmt.Println("Will Not Print1")
	case true:
		fmt.Println("Will Print2")
	case (2==2):
		fmt.Println("Will Print3")
	case (3==5):
		fmt.Println("Will Not Print4")

	}
}

