package main

import (
	"fmt"
)

const C1 = 2.4
var D1 string = "First program" //Static Type declaration
var F1, G1, H1 = "Second", 222, 333.990

func main() {
	calculate()
}

func calculate() {
	E1 := 3 //Dynamic type declaration
	fmt.Println("Value of C1:", C1)
	fmt.Printf("Constant variable type is: %T\n", C1)
	fmt.Println("Value of D1:", D1)
	fmt.Printf("Dynamic variable type is: %T\n",D1)
	fmt.Println("Value of E1:", E1)
	fmt.Printf("Dynamic variable type is: %T\n",E1)
	fmt.Println("Value of F1:", F1)
	fmt.Printf("Dynamic variable type is: %T\n",F1)
	fmt.Println("Value of G1:", G1)
	fmt.Printf("Dynamic variable type is: %T\n",G1)
	fmt.Println("Value of H1:", H1)
	fmt.Printf("Dynamic variable type is: %T\n",H1)

}