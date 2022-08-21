package main

import (
	"fmt"
	"sync"
	"time"
)

var wg sync.WaitGroup

//Gorutinas
func main() {
	name := "Rappi"
	wg.Add(2)

	var hello, bye string

	go func() {
		defer wg.Done()
		hello = sayHello(name)
	}()

	go func() {
		defer wg.Done()
		bye = sayBye(name)
	}()

	wg.Wait()
	fmt.Println(hello)
	fmt.Println(bye)
}

func sayHello(name string) string {
	time.Sleep(time.Second * 3)
	return fmt.Sprintf("Hello %s", name)
}

func sayBye(name string) string {
	time.Sleep(time.Second * 1)
	return fmt.Sprintf("Bye %s", name)
}
