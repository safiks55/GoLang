package main

import (
	"fmt"
	pkg "packages"
	"time"
)

func timesThree(number int, ch chan int) {
	pkg.ExecTime(time.Now(), "timesThree")
	result := number * 3
	fmt.Println(number * 3)
	ch <- result
}

func main() {
	fmt.Println("We are executing a goroutine")
	ch := make(chan int)
	go timesThree(5, ch)
	result := <-ch
	fmt.Printf("The result is: %v", result)
}
