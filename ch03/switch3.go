package main
import ("fmt");

func main(){
	switch {
	case false:
		fmt.Println("Will Not Print1")
	case true:
		fmt.Println("Will Print2")
		fallthrough
	case (2==2):
		fmt.Println("Will Print3")
		fallthrough
	case (3==5):
		fmt.Println("Will Not Print4")
	case (5==5):
		fmt.Println("Will Not Print5")
	default :
		fmt.Println("Default Print5")

	}
}

