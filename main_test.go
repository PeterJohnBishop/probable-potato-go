package main

import (
	"fmt"
	"probable-potato-go/main.go/basic"
	"probable-potato-go/main.go/validation"
	"testing"
)

func TestHello(t *testing.T) {

	fmt.Println(basic.Hello("Peter"))

	got := basic.Hello(("Chris"))
	want := "Hello, Chris"

	if got != want {
		t.Errorf("got %q want %q", got, want)
	}
}

func TestAdder(t *testing.T) {
	sum := basic.Add(2, 2)
	expected := 4

	fmt.Println(basic.Add(2, 2))

	if sum != expected {
		t.Errorf("expected '%d' but got '%d'", expected, sum)
	}
}

func TestRepeat(t *testing.T) {
	repeated := basic.Repeat("a")
	expected := "aaaaa"

	fmt.Println(basic.Repeat(("a")))

	if repeated != expected {
		t.Errorf("expected %q but got %q", expected, repeated)
	}
}

func TestSum(t *testing.T) {

	numbers := [5]int{1, 2, 3, 4, 5}

	fmt.Println(basic.Sum((numbers)))

	got := basic.Sum(numbers)
	want := 15

	if got != want {
		t.Errorf("got %d want %d given, %v", got, want, numbers)
	}
}

func TestSumGroup(t *testing.T) {

	t.Run("collection of 5 numbers", func(t *testing.T) {
		numbers := [5]int{1, 2, 3, 4, 5}

		fmt.Println(basic.Sum((numbers)))

		got := basic.Sum(numbers)
		want := 15

		if got != want {
			t.Errorf("got %d want %d given, %v", got, want, numbers)
		}
	})

	t.Run("collection of any size", func(t *testing.T) {
		numbers := []int{1, 2, 3}

		fmt.Println(basic.SumSlice(numbers))

		got := basic.SumSlice(numbers)
		want := 6

		if got != want {
			t.Errorf("got %d want %d given, %v", got, want, numbers)
		}
	})

}

func TestConvertInput(t *testing.T) {
	exampleCC := "1036787567886619"

	got := validation.ConvertInput(exampleCC)
	want := [16]int{1, 0, 3, 6, 7, 8, 7, 5, 6, 7, 8, 8, 6, 6, 1, 9}

	if got != want {
		fmt.Printf("got %d want %d", got, want)
		t.Errorf("got %v want %v", got, want)
	}

}

func TestValidate(t *testing.T) {
	// Example credit card array
	exampleCC := [16]int{1, 0, 3, 6, 7, 8, 7, 5, 6, 7, 8, 8, 6, 6, 1, 9}

	// Call Validate function
	got := validation.Validate(exampleCC)

	// Expected result
	want := 00

	if got != want {
		t.Errorf("Test failed: got %d, want %d", got, want)
	} else {
		fmt.Println("Credit Card Number validated by Luhn's Algorithm")
	}
}
