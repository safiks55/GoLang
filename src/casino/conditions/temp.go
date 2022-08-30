package main

import (
	snow "casino/pkg"
	"fmt"
	"github.com/bwmarrin/snowflake"
)

type lottos struct {
	megaA [5]int
	megaB [1]int

	powerA [5]int
	powerB [1]int

	tx    [6]int
	cash5 [5]int

	pick3  [3]int
	daily4 [4]int
}

type user struct {
	user_id snowflake.ID
	tickets []lottos
}

func generateN(numTickets int, buyer user) {
	buyer.user_id = snow.Generate()

}

func main() {
	var input int
	for ok := true; ok; ok = input != 0 {
		fmt.Println("Sussy's lottery")
		fmt.Println("Choose from the following options")
		fmt.Println("1.Buy Powerball tickets")
		fmt.Println("2.Buy Megamillions tickets")
		fmt.Println("3.Buy Lotto-texas tickets")
		fmt.Println("4.Buy Cash-5 tickets")
		fmt.Println("5.Buy Pick-3 tickets")
		fmt.Println("6.Buy Daily-4 tickets")
		fmt.Println("7.Check Jackpots")
		fmt.Println("8.Previous winning numbers")
		fmt.Println("9.How to play")
		fmt.Println("0. Exit")

		n, err := fmt.Scanln(&input)
		if n < 1 || err != nil {
			fmt.Println("invalid input")
			break
		}

	}

}
