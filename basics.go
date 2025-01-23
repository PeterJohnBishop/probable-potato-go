package main

func Hello(name string) string {
	return "Hello, " + name
}

func Add(n1 int, n2 int) int {
	return n1 + n2
}

func Repeat(s string) string {
	var response string
	for i := 0; i < 5; i++ {
		response = response + s
	}
	return response
}

func Sum(numbers [5]int) int {
	sum := 0
	// for i := 0; i < 5; i++ {
	// 	sum += numbers[i]
	// }

	for _, number := range numbers {
		sum += number
	}
	return sum
}

func SumSlice(numbers []int) int {
	sum := 0
	// for i := 0; i < 5; i++ {
	// 	sum += numbers[i]
	// }

	for _, number := range numbers {
		sum += number
	}
	return sum
}
