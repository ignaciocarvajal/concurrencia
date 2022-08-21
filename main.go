package main

import (
	"fmt"
	"time"
)

//Gorutinas
func main() {
	name := "Falabella Financiero"
	hello := make(chan string)
	defer close(hello)
	bye := make(chan string)
	defer close(bye)

	go sayHello(name, hello)
	go sayBye(name, bye)

	fmt.Println(<-hello)
	fmt.Println(<-bye)
}

func sayHello(name string, channel chan string) {
	time.Sleep(time.Second * 3)
	channel <- fmt.Sprintf("Hello %s", name)
}

func sayBye(name string, channel chan string) {
	time.Sleep(time.Microsecond * 1)
	channel <- fmt.Sprintf("Bye %s", name)
}
