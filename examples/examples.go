package main

import (
	"fmt"
	"log"

	"github.com/bnse/temperature"
)

// Usage:
func do[T temperature.Temperature](t T) T {
	// dosomething
	fmt.Println("dosomething", t)
	return t
}

func main() {
	// define
	lambda := func(f float64) {
		// 1. get a template form whatever
		t := temperature.Celsius(f)

		// 2. check the tempalte is valid, if NotValid, return
		if !temperature.IsValid(t) {

			log.Printf("temperature %v is not valid\n", t)
			return
		}
		// 3. process the main logic
		fmt.Printf("temperature %v is valid, normal process the main logic\n", t.FixDecimal(2))
		fmt.Printf("%v = %v\n", t, t.Fahrenheit())
		fmt.Printf("%v = %v\n", t, t.Kelvin())
	}
	// use
	lambda(37.1234)
	do(temperature.Celsius(37.1234))
	do[temperature.Celsius](37.1234)

	{

		fmt.Println("Celsius:")
		c1NotValid := temperature.Celsius(-273.16)
		fmt.Println(temperature.IsValid(c1NotValid))
		c1 := temperature.Celsius(37.123)
		fmt.Println(temperature.IsValid(c1))
		fmt.Println(temperature.AbsoluteZero(c1))
		fmt.Println(c1)
		fmt.Println(c1.FixDecimal(1))
		fmt.Println(c1.Fahrenheit())
		fmt.Println(c1.Fahrenheit().FixDecimal(2))
		fmt.Println(c1.Kelvin())
		fmt.Println(c1.Kelvin().FixDecimal(2))

		fmt.Println("Fahrenheit:")
		f1NotValid := temperature.Fahrenheit(-459.90)
		fmt.Println(temperature.IsValid(f1NotValid))
		f1 := temperature.Fahrenheit(98.8214)
		fmt.Println(temperature.IsValid(f1))
		fmt.Println(f1)
		fmt.Println(f1.FixDecimal(2))
		fmt.Println(temperature.AbsoluteZero(f1))
		fmt.Println(temperature.IsValid(f1))
		fmt.Println(f1.Celsius())
		fmt.Println(f1.Celsius().FixDecimal(2))
		fmt.Println(f1.Kelvin())
		fmt.Println(f1.Kelvin().FixDecimal(2))

		fmt.Println("Kelvin")
		k1 := temperature.Kelvin(421.12341)
		k1NotValid := temperature.Kelvin(-1)
		fmt.Println(k1)
		fmt.Println(k1.FixDecimal(2))
		fmt.Println(temperature.AbsoluteZero(k1))
		fmt.Println(temperature.AbsoluteZero(k1) > k1)
		fmt.Println(temperature.IsValid(k1))
		fmt.Println(temperature.IsValid(k1NotValid))
		fmt.Println(k1.Celsius())
		fmt.Println(k1.Fahrenheit())
	}
}
