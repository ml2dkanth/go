package main

import (
	"fmt"
"io"
"os"
"strings"
)

func main() {

	myfile, err := os.Create("GoSample.txt")

	if (err!=nil){
		fmt.Println("Error :",err)
	}
	defer myfile.Close()

	refFile := strings.NewReader("Srikanth Deva")

	io.Copy(myfile, refFile)
}

