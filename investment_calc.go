package main

import (
	"fmt"
	"math"
)

const inflationRate = 5.5

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
	// factoring also inflation rate
	// const inflationRate = 2.5
	// investedAmount := 1000.00 // example: if we want to explicitly also declare a variable without kw: var; we can do so with ':=' assignment operator

	// if we want to dynamically allocate (thru user input):
	var investedAmount float64
	fmt.Print("Enter the invested amount here:")
	fmt.Scan(&investedAmount)

	var expectedReturnRate float64 = 5.5
	var years float64 = 10

	// var futureValue = investedAmount * math.Pow(1+expectedReturnRate/100, years) // we are typecasting a few variables here
	// var futureValueReal = futureValue / math.Pow(1+inflationRate/100, years)

	futureValue, futureValueReal := calculateFutureValues(investedAmount, expectedReturnRate, years) // calling the function here

	fmt.Println("Here is the future value (rounded):", math.Round(futureValue)) // printing the value in the terminal with a NEW LINE (ln does that with Print)
	fmt.Println("Here is the inflated value (rounded):", math.Round(futureValueReal))
	fmt.Printf("Future value: %f\nInflated value: %f", futureValue, futureValueReal) // Printf allows us to have multiple vars
}

// Helper Functions
// 1) all variables within parantheses can be heterogeneous, its because all the vars below are homogeneous hence float64 has been declared only once
// 2) the output variable types also need to be declared since there are multiple outputs returned

func calculateFutureValues(investedAmount, expectedReturnRate, years float64) (float64, float64) {
	fv := investedAmount * math.Pow(1+expectedReturnRate/100, years)
	rfv := fv / math.Pow(1+inflationRate/100, years)
	return fv, rfv
}

// NOTES:

// %v - for normal values, %f - for float values
