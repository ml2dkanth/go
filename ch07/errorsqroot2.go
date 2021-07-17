package main

import (
	"errors"
	"fmt"
	"log"
)

var ErrInfo= errors.New("Info: square root of negative number")

func main() {
	fmt.Printf("Type of Error %T\n", ErrInfo)
	_, err := sqrt(-10)
	if err != nil {
		log.Fatalln(err)
	}
}

func sqrt(f float64) (float64, error) {
	if f < 0 {
		return 0, ErrInfo
	}
	return 42, nil
}

// see use of errors.New in standard library:
// http://golang.org/src/pkg/bufio/bufio.go
// http://golang.org/src/pkg/io/io.go
