package main

import (
	"fmt"
	"sync"
)

func main() {
	var wg sync.WaitGroup
	wg.Add(1)
	go func() {
		defer wg.Done()
		grindCoffeeBeans()
	}()
	wg.Add(1)
	go func() {
		defer wg.Done()
		frothMilk()
	}()
	wg.Wait()
}

func grindCoffeeBeans() {
	fmt.Println("Grinding coffee beans...")
}

func frothMilk() {
	fmt.Println("Frothing milk...")
}
