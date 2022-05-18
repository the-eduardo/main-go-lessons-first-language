package main

import "fmt"

func fuelGauge(fuel int) {
	fmt.Println("Fuel left:", fuel)
}

func calculateFuel(planet string) int {
	var fuel int
	switch planet {
	case "Venus":
		fuel = 300000
	case "Mercury":
		fuel = 500000
	case "Mars":
		fuel = 700000
	default:
		fuel = 0
	}
	return (fuel)
}

/** TODO
- Add a variable that keeps track of which planet your spaceship is on.
- Create a function that returns your spaceship back to your home planet.
- Add more destinations and allow for traveling between different planets.
 **/


func greetPlanet(planet string) {
	fmt.Println("Welcome to Planet", planet)
}

func cantFly() {
	fmt.Println("We do not have the available fuel to fly there.")
}

func flyToPlanet(planet string, fuel int) int {
	var fuelRemaining, fuelCost int
	fuelRemaining = fuel
	fuelCost = calculateFuel(planet)

	if fuelRemaining >= fuelCost {
		(greetPlanet(planet))
		fuelRemaining -= fuelCost
	} else {
		cantFly()
	}
	return fuelRemaining
}

func main() {
	var fuel int = 100000
	var planetChoice string
	fmt.Print("Select a planet to fly: \n", "Venus, Mercury, Mars\n")
	fmt.Scan(&planetChoice)
	fuel = flyToPlanet(planetChoice, fuel)
	fuelGauge(fuel)
}
