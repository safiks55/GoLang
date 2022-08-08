package main

import "fmt"

func fib(c, quit chan int) {
	x, y := 0, 1
	for {
		select {
		case c <- x:
			x, y = y, x+y

		case quit <- x:
			fmt.Println("-------------End-------------")
		}
	}
}

func main() {
	//temp := make(chan int)
	//quit := make(chan int)
	var n int
	fmt.Println("Please input the number till you want the sequence:")
	fmt.Scan(n)

	/*go func() {
		for i := 0; i < n; i++ {
			fmt.Println(<-temp)
		}
		quit <- 0
	}()
	fib(temp, quit)*/
}
