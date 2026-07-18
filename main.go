package main

import (
	"fmt"
	"math"
	"math/rand"
	"sync"
)

func main() {
	var wg sync.WaitGroup
	intChan := make(chan int)
	powChan := make(chan int)

	wg.Add(2)
	go func() {
		defer wg.Done()
		createRandomSlice(intChan)
	}()
	go func() {
		defer wg.Done()
		for value := range intChan {
			powChan <- int(math.Pow(float64(value), 2))
		}
	}()

	go func() {
		wg.Wait()
		close(powChan)
	}()

	for value := range powChan {
		fmt.Println(value)
	}
}

func createRandomSlice(ch chan<- int) {
	var slice []int

	for range 10 {
		number := rand.Intn(100)
		ch <- number
		slice = append(slice, number)
	}
	close(ch)
}
