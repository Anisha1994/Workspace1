package main

import (
	"fmt"
)

type CarDetails interface {
	CarDetails()
}

type Vehicle struct {
	List [3]CarDetails
}

type Car struct {
	Sound string
}

type Owner struct {
	Owner string
}

type Insurance struct {
	Insurance string
}


func main() {
	vehicle := new(Vehicle)
	vehicle.List[0] = &Car{"Car: BMW"}
	vehicle.List[1] = &Owner{"Onwer : San Shah"}
	vehicle.List[2] = &Insurance{"Insurance: Remote service"}


	for _, hornSounder := range vehicle.List {
		hornSounder.CarDetails()
	}
}

func (car *Car) CarDetails() {
	fmt.Println(car.Sound)
}

func (owner *Owner) CarDetails() {
	fmt.Println(owner.Owner)
}

func (insurance *Insurance) CarDetails() {
	fmt.Println(insurance.Insurance)
}


func CarD(cardetails CarDetails) {
	cardetails.CarDetails()
}