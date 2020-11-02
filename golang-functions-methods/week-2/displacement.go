package main

import "fmt"

func main() {
	a := TakeInput("acceleration")
	v0 := TakeInput("initial velocity")
	s0 := TakeInput("displacement")

	fn := GenDisplaceFn(a, v0, s0)
	t := TakeInput("time")

	fmt.Println("The displacement is", fn(t))
}

func TakeInput(varName string) float64 {
	var num float64

	fmt.Printf("Please, enter initial value if %s:\n", varName)

	_, err := fmt.Scan(&num)

	if err != nil {
		panic(err)
	}

	return num
}

func GenDisplaceFn(a, v0, s0 float64) func(float64) float64 {
	return func(t float64) float64 {
		return 0.5 * a * t * t + (v0 * t) + s0
	}
}
