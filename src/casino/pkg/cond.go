package pkg

/*
Power-Ball
-1 -Jackpot - if all numbers match both pickA and pickB
-2 $ x * 1 million - if 5 numbers match of pick A
-3 $ 50k - if 4 out of 5 match for pick A and match pickB
-4 $ 1k - if 4 of 5 match for pickA
-5 $ 5 - if pickB matches

each normal ticket
*/

func Powerball(white [5]int, red [1]int) int {
	count := 0
	pickA, pickB := Big_Lotto(69, 26)
	if pickA == white && pickB == red {
		return -1
	} else {
		for i := 0; i < 5; i++ {
			for j := 0; j < 5; j++ {
				if pickA[i] == white[j] {
					count += 1
				}
			}
		}
		if count == 5 && pickB != red {
			return -2
		}
		if count == 4 && pickB == red {
			return -3
		}
		if count == 4 && pickB != red {
			return -4
		}
		if count == 0 && pickB == red {
			return -5
		}
		return count
	}
}

/*
Mega-Millions
-1 -Jackpot - if all numbers match both pickA and pickB
-2 $ x * 1 million - if 5 numbers match of pick A
-3 $ 50k - if 4 out of 5 match for pick A and match pickB
-4 $ 1k - if 4 of 5 match for pickA
-5 $ 5 - if pickB matches

each normal ticket
*/

func MegaMillions(white [5]int, red [1]int) int {
	count := 0
	pickA, pickB := Big_Lotto(70, 25)
	if pickA == white && pickB == red {
		return -1
	} else {
		for i := 0; i < 5; i++ {
			for j := 0; j < 5; j++ {
				if pickA[i] == white[j] {
					count += 1
				}
			}
		}
		if count == 5 && pickB != red {
			return -2
		}
		if count == 4 && pickB == red {
			return -3
		}
		if count == 4 && pickB != red {
			return -4
		}
		if count == 0 && pickB == red {
			return -5
		}
		return count
	}
}

/*
lotto-Texas
-1 Jackpot - 6 of 6 numbers match
-2 $ 10k - 5 of 6 numbers match
-3 $ 1k - 4 of 6 numbers match
-4 $ 5 - 3 of 6 numbers match
*/

func LottoTexas(white [6]int) int {
	count := 0
	pickA := Lotto_Texas(54)
	if pickA == white {
		return -1
	} else {
		for i := 0; i < 6; i++ {
			for j := 0; j < 6; j++ {
				if pickA[i] == white[j] {
					count += 1
				}
			}
		}
		if count == 5 {
			return -2
		}
		if count == 4 {
			return -3
		}
		if count == 3 {
			return -4
		}
		return count
	}
}

/*
Cash-Five
-1 jackpot $ 50K - 5 of 5 numbers match
-2 $ 1k - 4 of 5 numbers match
-3 $ 100 - 3 of 5 numbers match
-4 Free ticket - 2 of 5 numbers match
*/

func Cash5(white [5]int) int {
	count := 0
	pickA := CashFive(35)

	if pickA == white {
		return -1
	} else {
		for i := 0; i < 5; i++ {
			for j := 0; j < 5; j++ {
				if pickA[i] == white[j] {
					count += 1
				}
			}
		}
		if count == 4 {
			return -2
		}
		if count == 3 {
			return -3
		}
		if count == 2 {
			return -2
		}
		return count
	}
}

/*
All-or-Nothing
-1 $250k - if 12/12 match or 0/12 match
-2 $500 - if 11/12 match or 1/12 match
-3 $100 - if 10/12 match or 2/12 match
-4 $50 - if 9/12 match or 3/12 match
*/

func AllorNothing(white [12]int) int {
	count := 0
	pickA := All_or_Nothing(24)
	if pickA == white {
		return -1
	} else {
		for i := 0; i < 12; i++ {
			for j := 0; j < 12; j++ {
				if pickA[i] == white[j] {
					count += 1
				}
			}
		}
		if count == 11 || count == 1 {
			return -2
		}
		if count == 10 || count == 2 {
			return -3
		}
		if count == 9 || count == 3 {
			return -4
		}
		return count
	}
}

/*
Pick-3
prize-combination-price
$500 - exact $1
$250 - exact/any - $1
$250 - exact - $.50
$50 - any - $1
*/

func Pick3(b [3]int) int {
	count := 0
	a := Pick_3(10, 0)
	if a == b {
		return -1
	} else {
		for i := 0; i < 3; i++ {
			for j := 0; j < 3; j++ {
				if a[i] == b[j] {
					count += 1
				}
			}
		}
		if count == 3 {
			return -2
		}
	}
	return 0
}

/*
Daily-4
$5000 - exact - $1
$1200 - box - 	$1
$50  - 2 way - $1
*/

func Daily4(b [4]int) int {
	count := 0
	a := Daily_4(10, 0)
	if a == b {
		return -1
	} else {
		for i := 0; i < 4; i++ {
			for j := 0; j < 4; j++ {
				if a[i] == b[j] {
					count += 1
				}
			}
		}
		if count == 4 {
			return -2
		}
	}
	return 0
}
