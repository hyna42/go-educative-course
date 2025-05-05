package main

import (
	"fmt"
	"reflect"
)

func PrintMessage(message string) {
	fmt.Println(message)
}

func magic(a int, b int) (int, int) {
	//retourne la somme et le produit des 2 entiers passe en paramètre
	return a + b, a * b
}

func main() {
	PrintMessage("hello word again !")
	sum, prod := magic(2,5)

	fmt.Println("Somme :", sum)
	fmt.Println("Produit :", prod)


	
	fmt.Println("Type of  :", reflect.TypeOf(sum))

}
