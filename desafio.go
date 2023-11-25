package main

import "fmt"

var tempKelvin = 373.0

func main() {
	tempCelsius := tempKelvin - 273

	fmt.Printf("A temperatura de ebulição em °C é: %g", tempCelsius)
}
