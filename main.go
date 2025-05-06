package main

import (
	"fmt"
	"math/rand"
	"reflect"
)

func PrintMessage(message string) {
	fmt.Println(message)
}

func magic(a int, b int) (int, int) {
	//retourne la somme et le produit des 2 entiers passe en paramètre
	return a + b, a * b
}

var number int = 5 

func main() {
	PrintMessage("hello word again !")

	sum, prod := magic(2, 5)

	fmt.Println("Somme :", sum)
	fmt.Println("Produit :", prod)

	fmt.Println("Type of  :", reflect.TypeOf(sum))

	//variables
	var num int;
	fmt.Println(num)

	var decision bool = true

	number = 10 

	a := rand.Int()
	b := rand.Int()

	fmt.Printf("New Value of number: %d\n",number)
    fmt.Printf("Value of decision: %t\n"  ,decision)

    fmt.Printf("Random  %d et %d " ,a,b)

}
