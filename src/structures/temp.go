package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"strconv"
)

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan()
	text := scanner.Text()

	num, _ := strconv.ParseFloat(text, 64)

	if num == 0 {
		fmt.Println(0)
		return
	}
	count := 1
	for num > 1 {
		count++
		tmp := math.Round(math.Log2(num))
		num = num - math.Pow(tmp, float64(2))
	}
	if math.Round(float64(count%3)) == 1 {
		fmt.Println(1)
		return
	}

	if math.Round(float64(count%3)) == 2 {
		fmt.Println(2)
		return
	}

	if math.Round(float64(count%3)) == 3 {
		fmt.Println(0)
		return
	}

	fmt.Println(0)
	return
}
