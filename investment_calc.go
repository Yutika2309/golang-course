package main

import (
	"fmt"
	"math"
)

func main() {
	var investedAmount = 1000
	var expectedReturnRate = 5.5
	var years = 10

	var futureValue = float64(investedAmount) * math.Pow(1+expectedReturnRate/100, float64(years)) // we are typecasting a few variables here

	fmt.Println("Here is the future value:", futureValue) // printing the value in the terminal with a NEW LINE (ln does that with Print)
}
