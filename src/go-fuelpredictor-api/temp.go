package main

import (
	"fmt"
	"math/rand"
	"sync"
	"time"
)

var goRoutine sync.WaitGroup

func main() {
	rand.Seed(time.Now().Unix())

	marks := make(chan string, 10)
	goRoutine.Add(3)
	for i := 1; i <= 3; i++ {
		go student(marks, i)
	}

	for j := 1; j <= 3; j++ {
		marks <- fmt.Sprintf("Work:%d", j)
	}
	close(marks)
	goRoutine.Wait()
}

func student(marks chan string, student int) {
	sleep := rand.Int63n(30)
	time.Sleep(time.Duration(sleep) * time.Millisecond)
	fmt.Println("\n time to sleep", sleep, "ms\n")
	
}
