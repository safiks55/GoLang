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
// we apply the formulae on the max element to find the bitsize for rest of the input
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
// this function adds leading zeros (01=>010) depending on the bitsize requirement
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
// we use formatInt function to convert the decimal number to a binary
func decToBin(decimal int, bSize int) string {
	num := int64(decimal)
	temp := strconv.FormatInt(num, 2)
	res := addZeros(temp, bSize)
	return res
}

// splits the message in n parts
//we iterate over x.number of elements and split the message and append it to a string slice
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
// we compare the map-values to the data in the slice and append the corresponding keys to a string
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
//we pass the decoded message to this function to get the frequency of each character in the message
func findFrequency(str string) map[string]int {
	mp := make(map[string]int)
	for _, char := range str {
		mp[string(char)] = mp[string(char)] + 1
	}
	return mp
}

//Character: a, Code: 010, Frequency: 3
//print function to print the output
func printResult(alpha []string, mp map[string]int, mp2 map[string]string, decodedMsg string) {
	fmt.Println("Alphabet:")
	for _, i := range alpha {
		fmt.Println("Character: ", i, ",", "Code: ", mp2[i], ",", " Frequency:", mp[i])
	}
	fmt.Println("")
	fmt.Println("Decompressed Message: ", decodedMsg)
}
func main() {
	var (
		alpha   []string // store alphabets here
		dec     []int    // store decimals here
		message string   // store the binary message here
		temp    string   // store the input word by word
		size    int      // number of alphabets/decimals
	)
	// opening the text file
	input, err := os.Open("input3.txt")
	if err != nil {
		log.Fatal(err)
	}
	// opening the text file
	defer input.Close()

	// reading and storing the input
	scanner := bufio.NewScanner(input)
	for scanner.Scan() {
		temp = scanner.Text()
		// if first thing in the input is integer then we know its the size
		if _, err := strconv.Atoi(temp); err == nil {
			size, _ = strconv.Atoi(temp) // storing the size
		} else {
			//keep taking input until the size == len of dec and alpha slice
			if len(alpha) < size && len(dec) < size {
				alphabet := string(temp[0])
				decimal, _ := strconv.Atoi(temp[2:])
				alpha = append(alpha, alphabet)
				dec = append(dec, decimal)
			}
		}
		// storing the binary message
		message = temp
	}
	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}

	// finding the bit-length using the ceil(log2(max_decimal_in_input)
	N := findBitsize(dec)

	// stores the initial decimals bin values in string format
	var bins []string
	for i := range dec {
		temp := decToBin(dec[i], N)
		bins = append(bins, temp)
	}

	// finding number of partitions we will get after splitting the message
	n := len(message) / findBitsize(dec)
	splitString := splitString(message, n)

	// initialize a  map and "map-binary numbers to the respective alphabet they represent"
	myMap := make(map[string]string)
	for i := 0; i < len(bins); i++ {
		myMap[bins[i]] = alpha[i]
	}
	// decode the message using the helper function
	decodedMessage := decodeBinary(myMap, splitString)

	//making temporary map so we can print the output in the same for loop
	dummy := make(map[string]string)
	for i := 0; i < len(bins); i++ {
		dummy[alpha[i]] = bins[i]
	}
	//finding the frequency of each alphabet in the message using the helper function
	tempMap := findFrequency(decodedMessage)

	//prinitng the result using the helper function
	printResult(alpha, tempMap, dummy, decodedMessage)

}
