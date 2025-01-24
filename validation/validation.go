package validation

import (
	"fmt"
	"strconv"
)

type Credit struct {
	TextCardNumber string
	CardNumberArr  [16]int
	Valid          bool
}

func (cc *Credit) ConvertInput() [16]int {

	// Iterate through each character in the string and convert it to an integer
	for i, char := range cc.TextCardNumber {
		// Convert the character to an integer
		num, err := strconv.Atoi(string(char))
		if err != nil {
			fmt.Println("Error converting string to integer:", err)
		}
		// Assign the integer to the corresponding index in the array
		cc.CardNumberArr[i] = num
	}

	// Print the result
	fmt.Println(cc.CardNumberArr)
	return cc.CardNumberArr
}

func (cc *Credit) Validate() bool {

	var arrSum int
	for i, num := range cc.CardNumberArr {
		if i == 0 || i%2 == 0 {
			// double ever other number starting at ccSlice[0]
			n := num * 2
			if n <= 9 {
				// if result is 9 or less replace with result at position
				cc.CardNumberArr[i] = n
			} else {
				// if result is 10 or greater sum first + second digit
				// replace with result at position
				numStr := strconv.Itoa(n)
				var nArr [2]int
				for i, char := range numStr {
					// Convert each character to an integer and store in the array
					nArr[i], _ = strconv.Atoi(string(char))
				}
				cc.CardNumberArr[i] = nArr[0] + nArr[1]
			}
		}
	}
	for _, num := range cc.CardNumberArr {
		arrSum = arrSum + num
	}
	numStrToo := strconv.Itoa(arrSum)
	var nArrToo [2]int
	for i, char := range numStrToo {
		// Convert each character to an integer and store in the array
		nArrToo[i], _ = strconv.Atoi(string(char))
	}
	if nArrToo[1] == 0 {
		cc.Valid = true
		return cc.Valid
	} else {
		cc.Valid = false
		return cc.Valid
	}
}
