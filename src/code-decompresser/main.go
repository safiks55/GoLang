package main

import (
	"bufio"
	"fmt"
	"log"
	"math"
	"os"
	"strconv"
)

// finds the bitsize for binary conversion
func findBitsize(decimal []int) int {

	var temp, max int
	for _, element := range decimal {
		if element > temp {
			temp = element
			max = temp
		}
	}
	max = int(math.Ceil(math.Log2(float64(max))))
	return max
}

// adds zero padding to the binary numbers
func addZeros(result string, x int) string {
	size := x - len(result)
	if x != len(result) {
		for i := 0; i < size; i++ {
			temp1 := result
			result = "0"
			result += temp1
		}
	}
	return result
}

// converts decimals to binary numbers
func decToBin(decimal int, bSize int) string {
	num := int64(decimal)
	temp := strconv.FormatInt(num, 2)
	res := addZeros(temp, bSize)
	return res
}

// splits the message in n parts
func splitString(message string, n int) []string {
	length := len(message)
	parts := length / n
	var slc []string
	for x := 0; x < length; x += parts {
		temp := message[x : x+parts]
		slc = append(slc, temp)
	}
	return slc
}

// decodes the message
func decodeBinary(mp map[string]string, splits []string) string {
	decodedMessage := ""
	var temp string
	for i := range splits {
		temp = mp[splits[i]]
		decodedMessage += temp
	}
	return decodedMessage
}

//finds frequency of a character within the string
func freqMap(arr []string) map[string]int {
	freq := make(map[string]int)
	for _, num := range arr {
		freq[num] = freq[num] + 1
	}
	return freq
}

//print function to print the output
func printResult(alpha []string, bins []string)
func main() {
	var (
		alpha   []string // store alphabets here
		dec     []int    // store decimals here
		message string   // store the binary message here
		temp    string   // store the input word by word
		size    int      // number of alphabets/decimals
	)

	input, err := os.Open("input3.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer input.Close()

	// reading and storing the input
	scanner := bufio.NewScanner(input)
	for scanner.Scan() {
		temp = scanner.Text()
		if _, err := strconv.Atoi(temp); err == nil {
			size, _ = strconv.Atoi(temp)
		} else {
			if len(alpha) < size && len(dec) < size {
				alphabet := string(temp[0])
				decimal, _ := strconv.Atoi(temp[2:])
				alpha = append(alpha, alphabet)
				dec = append(dec, decimal)
			}
		}
		message = temp
	}
	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}
	N := findBitsize(dec)
	var bins []string // stores the initial decimals bin values
	for i := range dec {
		temp := decToBin(dec[i], N)
		bins = append(bins, temp)
	}
	n := len(message) / findBitsize(dec)
	splitString := splitString(message, n)

	myMap := make(map[string]string)
	for i := 0; i < len(bins); i++ {
		myMap[bins[i]] = alpha[i]
	}
	decodedMessage := decodeBinary(myMap, splitString)

	fmt.Println(decodedMessage)

}