package pkg

import (
	"math/rand"
	"time"
)

func iterativeDigitsCount(number int) int {
	count := 0
	for number != 0 {
		number /= 10
		count += 1
	}
	return count
}

/*
- Powerball
- 2 sets
- set 1 pick 5 numbers (range 1-69) cannot repeat
- set 2 pick 1 number (range 1-26)
- MegaMillions
- 2 sets
- set 1 pick 5 numbers (range 1-70) cannot repeat
- set 2 pick 1 number (range 1-25)
*/

func Big_Lotto(maxA, maxB int) ([5]int, [1]int) {
	rand.Seed(time.Now().UnixNano())
	var pickA [5]int
	var pickB [1]int
	unique := -1
	var number int
	for i := 0; i < 5; i++ {
		for {
			number = rand.Intn(maxA-1) + 1
			unique = 1
			for j := 0; j < i; j++ {

				if pickA[j] == number {
					unique = 0
				}
			}
			pickA[i] = number
			if unique != 0 {
				break
			}
		}
	}
	pickB[0] = rand.Intn(maxB-1) + 1
	return pickA, pickB
}

//Texas lotto pick 6 numbers range(1-54) cannot repeat

func Lotto_Texas(max int) [6]int {
	rand.Seed(time.Now().UnixNano())
	var pick [6]int
	unique := -1
	var number int
	for i := 0; i < 6; i++ {
		for {
			number = rand.Intn(max-1) + 1
			unique = 1
			for j := 0; j < i; j++ {

				if pick[j] == number {
					unique = 0
				}
			}
			pick[i] = number
			if unique != 0 {
				break
			}
		}
	}
	return pick
}

// CashFive pick 5 numbers range(1-35) cannot repeat

func CashFive(max int) [5]int {
	rand.Seed(time.Now().UnixNano())
	var pick [5]int
	unique := -1
	var number int
	for i := 0; i < 5; i++ {
		for {
			number = rand.Intn(max-1) + 1
			unique = 1
			for j := 0; j < i; j++ {

				if pick[j] == number {
					unique = 0
				}
			}
			pick[i] = number
			if unique != 0 {
				break
			}
		}
	}
	return pick
}

//pick3 pick 3 numbers each within 1-10 range can repeat

func Pick_3(max, min int) [3]int {
	rand.Seed(time.Now().UnixNano())
	var pick3 [3]int
	for i := 0; i < 3; i++ {
		pick3[i] = rand.Intn(max-min) + min
	}
	return pick3
}

// daily4 pick 4 numbers each within 1-10 range can repeat

func Daily_4(max, min int) [4]int {
	rand.Seed(time.Now().UnixNano())
	var pick3 [4]int
	for i := 0; i < 4; i++ {
		pick3[i] = rand.Intn(max-min) + min
	}
	return pick3
}

//All-or-Nothing pick 12 numbers range(1-24) cannot repeat

func All_or_Nothing(max int) [12]int {
	rand.Seed(time.Now().UnixNano())
	var pick [12]int
	unique := -1
	var number int
	for i := 0; i < 12; i++ {
		for {
			number = rand.Intn(max-1) + 1
			unique = 1
			for j := 0; j < i; j++ {

				if pick[j] == number {
					unique = 0
				}
			}
			pick[i] = number
			if unique != 0 {
				break
			}
		}
	}
	return pick
}

// Roll 2 dice and land 1 on both of them

func SnakeEyes(max, min int) (int, int) {
	rand.Seed(time.Now().UnixNano())
	dice1 := rand.Intn(max-min) + min
	dice2 := rand.Intn(max-min) + min
	return dice1, dice2
}
