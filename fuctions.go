package main

func Greeting(name *string) {
	println("In greeting: Hi first ", *name)
	*name = "Johnny"
	println("In greeting: Hi again", *name)
}
