package main

import "fmt"

type car struct {
	brand string
	features  string
}

func main() {
	var cartype = map[string]string{
		"\nBrand\t\t ": "Features\n",
		"Rolls Royce\t ": "Ultimate luxury sedan\n",
		"Fiat (EU)\t ": "Italian Volkswagen\n",
		"Lamborghini\t ": "Trendier than Ferrari\n",
		"Aston Martin\t ": "Comfortable GT, James Bond\n",
		"Toyota\t\t ": "Reliable\n",
		"Jaguar\t\t ": "Unique alternative to German cars\n",
	}
	fmt.Println(cartype)
}