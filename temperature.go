package main

type Celsius float32
type Fahrenheit float32

func toFahrenheit(t Celsius) Fahrenheit {
	return Fahrenheit(t*9/5) + 32
}
