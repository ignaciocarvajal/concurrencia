package main

import (
	"fmt"
)

//Gorutinas
func main() {
	name := "Rappi"
	done := make(chan bool, 1)
	go sayHello(name, done)
	<-done
	sayBye(name)
}

func sayHello(name string, done chan bool) {
	fmt.Printf("Hello %s \n", name)
	done <- true
}

func sayBye(name string) {
	fmt.Printf("Bye %s \n", name)

}
