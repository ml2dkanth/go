// Golang program to illustrate the usage of
// fmt.Scan() function

// Including the main package
package main

// Importing fmt
import (
    "fmt"
)

// Calling main
func main() {

    // Declaring some variables
    var name string
    var alphabet_count int

    // Calling Scan() function for
    // scanning and reading the input
    // texts given in standard input
    fmt.Print("Enter a String:")
    fmt.Scan(&name)
    fmt.Print("Word Count : ")
    fmt.Scan(&alphabet_count)

    // Printing the given texts
    fmt.Printf("The word %s containing %d number of alphabets.",
               name, alphabet_count)

}
