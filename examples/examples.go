package main

import (
	"fmt"
	"log"

	"github.com/bnse/temperature"
)

// Usage:
func do(t temperature.Celsius) temperature.Celsius {
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
		if !t.IsValid() {

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
	do(37.1234)

	fmt.Println("Celsius:")
	c1NotValid := temperature.Celsius(-273.16)
	fmt.Println(c1NotValid.IsValid())
	c1 := temperature.Celsius(37.123)
	fmt.Println(c1.IsValid())
	fmt.Println(c1)
	fmt.Println(c1.FixDecimal(1))
	fmt.Println(c1.Fahrenheit())
	fmt.Println(c1.Fahrenheit().FixDecimal(2))
	fmt.Println(c1.Kelvin())
	fmt.Println(c1.Kelvin().FixDecimal(2))

	fmt.Println("Fahrenheit:")
	f1NotValid := temperature.Fahrenheit(-459.90)
	fmt.Println(f1NotValid.IsValid())
	f1 := temperature.Fahrenheit(98.8214)
	fmt.Println(f1.IsValid())
	fmt.Println(f1)
	fmt.Println(f1.FixDecimal(2))
	fmt.Println(f1.AbsoluteZero())
	fmt.Println(f1.AbsoluteZero().IsValid())
	fmt.Println(f1.Celsius())
	fmt.Println(f1.Celsius().FixDecimal(2))
	fmt.Println(f1.Kelvin())
	fmt.Println(f1.Kelvin().FixDecimal(2))

	fmt.Println("Kelvin")
	k1 := temperature.Kelvin(421.12341)
	k1NotValid := temperature.Kelvin(-1)
	fmt.Println(k1)
	fmt.Println(k1.FixDecimal(2))
	fmt.Println(k1.AbsoluteZero())
	fmt.Println(k1.AbsoluteZero().IsValid())
	fmt.Println(k1.IsValid())
	fmt.Println(k1NotValid.IsValid())
	fmt.Println(k1.Celsius())
	fmt.Println(k1.Fahrenheit())
}
