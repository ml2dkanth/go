package main

import (
	"fmt"
	"log"
)

type CustomMathError struct {
	lat  string
	long string
	err  error
}

func (n CustomMathError) Error() string {
	return fmt.Sprintf("CustomMathError math error occured: %v %v %v", n.lat, n.long, n.err)
}

func main() {
	_, err := sqrt(-10.23)
	if err != nil {
		log.Println(err)
	}
}

func sqrt(f float64) (float64, error) {
	if f < 0 {
		nme := fmt.Errorf("Error math redux: square root of negative number: %v", f)
		return 0, CustomMathError{"50.2289 N", "99.4656 W", nme}
	}
	return 42, nil
}

// see use of structs with error type in standard library:
//
// http://golang.org/pkg/net/#OpError
// http://golang.org/src/pkg/net/dial.go
// http://golang.org/src/pkg/net/net.go
//
// http://golang.org/src/pkg/encoding/json/decode.go
