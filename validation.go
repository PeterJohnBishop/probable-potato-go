package main

import (
	"fmt"
	"strconv"
)

func ConvertInput(cardNumber string) [16]int {
	var arr [16]int

	// Iterate through each character in the string and convert it to an integer
	for i, char := range cardNumber {
		// Convert the character to an integer
		num, err := strconv.Atoi(string(char))
		if err != nil {
			fmt.Println("Error converting string to integer:", err)
		}
		// Assign the integer to the corresponding index in the array
		arr[i] = num
	}

	// Print the result
	fmt.Println(arr)
	return arr
}

func Validate(ccSlice [16]int) int {

	var arrSum int
	for i, num := range ccSlice {
		if i == 0 || i%2 == 0 {
			// double ever other number starting at ccSlice[0]
			n := num * 2
			if n <= 9 {
				// if result is 9 or less replace with result at position
				ccSlice[i] = n
			} else {
				// if result is 10 or greater sum first + second digit
				// replace with result at position
				numStr := strconv.Itoa(n)
				var nArr [2]int
				for i, char := range numStr {
					// Convert each character to an integer and store in the array
					nArr[i], _ = strconv.Atoi(string(char))
				}
				ccSlice[i] = nArr[0] + nArr[1]
			}
		}
	}
	for _, num := range ccSlice {
		arrSum = arrSum + num
	}
	numStrToo := strconv.Itoa(arrSum)
	var nArrToo [2]int
	for i, char := range numStrToo {
		// Convert each character to an integer and store in the array
		nArrToo[i], _ = strconv.Atoi(string(char))
	}
	if nArrToo[1] == 0 {
		return nArrToo[1]
	} else {
		return nArrToo[1]
	}
}
