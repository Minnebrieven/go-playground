package main

import (
	"errors"
	"fmt"
)

type Car struct {
	FuelIn float64
}

func (c *Car) setFuel(fuel float64) {
	c.FuelIn = fuel
}

func (c Car) EstimatedMile() float64 {
	const FuelConsumed float64 = 1.5
	return c.FuelIn / FuelConsumed
}

func main() {
	var car1 Car
	fmt.Print("\nEstimated Fuel Mileage. \nEnter how many fuel (liters) your car have = ")
	fmt.Scan(&car1.FuelIn)
	if car1.FuelIn <= 0 {
		car1.setFuel(0)
		fmt.Println(errors.New("fuel can't be 0 or less and should be number"))
	} else {
		fmt.Printf("\nFuel in Car : %+v L\nEstimated Mile can Reach : %.2f Mile\n", car1.FuelIn, car1.EstimatedMile())
	}

}
