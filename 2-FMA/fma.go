package main

import (
	"fmt"
)

const grav float32 = 9.81 // m s^-2

func forceEq(mass, acc float32) float32 {
	return mass * acc
}

func massEq(force, acc float32) float32 {
	return force / acc
}

func accEq(force, mass float32) float32 {
	return force / mass
}

func main() {
	mass := 72 // kg
	weight := forceEq(float32(mass), grav)
	fmt.Println(weight, "N")
	fmt.Println(massEq(weight, grav), "kg")
	fmt.Println(accEq(weight, float32(mass)), "m s^-2")
}
