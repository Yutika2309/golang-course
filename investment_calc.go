package main

import (
	"fmt"
	"math"
)

// func main() {
// 	var investedAmount = 1000
// 	var expectedReturnRate = 5.5
// 	var years = 10

// 	var futureValue = float64(investedAmount) * math.Pow(1+expectedReturnRate/100, float64(years)) // we are typecasting a few variables here

// 	fmt.Println("Here is the future value:", futureValue) // printing the value in the terminal with a NEW LINE (ln does that with Print)
// }

// ================================================================================================== //

// Here we explicilty mention the type of variable as the function
// above required us to mention the types of the variables in the equation; which bloats the equation
// also there's a way to explicitly mention variables:

func main() {
	investedAmount := 1000.00 // example: if we want to explicitly also declare a variable without kw: var; we can do so with ':=' assignment operator
	var expectedReturnRate float64 = 5.5
	var years float64 = 10

	var futureValue = investedAmount * math.Pow(1+expectedReturnRate/100, years) // we are typecasting a few variables here

	fmt.Println("Here is the future value (rounded):", math.Round(futureValue)) // printing the value in the terminal with a NEW LINE (ln does that with Print)
}
