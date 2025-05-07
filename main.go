package main

import (
	"fmt"
	"math/rand"
	"reflect"
	"strings"
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
	var num int
	fmt.Println(num)

	var decision bool = true

	number = 10

	a := rand.Int()
	b := rand.Int()

	fmt.Printf("New Value of number: %d\n", number)
	fmt.Printf("Value of decision: %t\n", decision)

	fmt.Printf("Random  %d et %d\n ", a, b)

	//test challenge temperature conversion
	temp := Celsius(100)

	fmt.Printf("%.2f°C = %2.f°F\n", temp, toFahrenheit(temp))

	//string
	str1 := " Hello test1 JOHN DOE " + "re hello    "

	fmt.Printf("string str 1 : , %s\n, lenStr1 ==> %d\n", str1, len(str1))

	fmt.Printf("str HasPrefix : %t\n", strings.HasPrefix(str1, "Hello"))

	fmt.Printf("str Index() : %d\n", strings.Index(str1, "hello"))

	str2 := strings.Replace(str1, "test1", "test2", -1)
	fmt.Printf("str2 Replace() : %s\n", strings.ToUpper(str2))

	trimmed := strings.TrimSpace(str2)

	fmt.Printf("str2 TrimSpace() : %s\n, lenStr2 [after Trim] ==> %d\n ", trimmed, len(trimmed))

	words := strings.Fields(trimmed)

	fmt.Printf("str2 : Split() : %v\n ", words)

	//test string8string.go functions
	phrasetestId := "fid:42"
	phrasetestEmail := "email:johdoe@educative.io"

	test1 := IdentifyPrefixPostfix(phrasetestId, phrasetestEmail)

	test2 := ContainsEducative(phrasetestEmail)

	test3 := MaskUserName("daddy@gmail.com")

	test4 := IndexOfAtSymbol("doe@gmail.com")

	test5 := TrimAndSplitUserID("id-                   52hF ")

	test6 := ConvertStringToInt("f123")

	fmt.Printf("test 1 : IdentifyPrefixPostfix() : %t\n ", test1)

	fmt.Printf("test 2 : ContainsEducative() : %t\n ", test2)

	fmt.Printf("test 3 : MaskUserName() : %s\n ", test3)

	fmt.Printf("test 4 : IndexOfAtSymbol() : %d\n ", test4)

	fmt.Printf("test 5 : TrimAndSplitUserID() : %s\n ", test5)

	fmt.Printf("test 6 : ConvertStringToInt() : %d\n ", test6)

}
