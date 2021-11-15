package main

import (
	"fmt"
)

//Gorutinas
func main() {
	name := "Rappi"
	sayHello(name)
	sayBye(name)
}

func sayHello(name string) {
	fmt.Printf("Hello %s \n", name)
}

func sayBye(name string) {
	fmt.Printf("Bye %s \n", name)
}
