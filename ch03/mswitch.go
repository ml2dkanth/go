package main
import ("fmt");

func main(){
	myvalue := "M"
	//myvalue := "sri"
	//myvalue := "x"

	switch myvalue {
	case "sri", "kanth" ,"rao":
			fmt.Println("Its my name")
	case "M":
		fmt.Println("Will Print2")
		fallthrough
	case "Q":
		fmt.Println("Will Print3")
		fallthrough
	case "R":
		fmt.Println("Will Not Print4")
	case "X":
		fmt.Println("Will Not Print5")
	default :
		fmt.Println("Default Print5")

	}
}

