package main

import (
	"encoding/hex"
	"fmt"
	"os"
)

func main() {
	reverse()

}

func reverse() string {
	rev := "This is a recurssion call function"
	if len(rev) == 0 {
		return "null"

}
	fmt.Println(rev)
	hx := hex.EncodeToString([]byte(rev))
	fmt.Println("String to Hex Golang example")
	fmt.Println(hx)
	fmt.Println(rev + " ==> " + hx)
	os.Exit(1)
	return reverse()


}
