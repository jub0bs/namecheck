package main

import (
	"fmt"
	"sync"
)

func grindCoffeeBeans(grams float64) {
	fmt.Printf("Grinding %f grams of coffee beans...\n", grams)
}

func frothMilk(ml float64) {
	fmt.Printf("Frothing %f ml of milk...\n", ml)
}

func main() {
	var wg sync.WaitGroup
	wg.Go(func() { grindCoffeeBeans(7) })
	wg.Go(func() { frothMilk(30) })
	wg.Wait()
}
