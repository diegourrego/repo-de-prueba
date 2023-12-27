package main

import "fmt"

func main()  {
	fmt.Println("Esto es una prueba")
	fmt.Println("Esto es otra prueba")
	greetTest("Diego")
}

func greetTest(name string) {
	fmt.Println("Hola", name)
}
