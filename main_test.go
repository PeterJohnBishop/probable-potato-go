package main

import (
	"fmt"
	"testing"
)

func TestHello(t *testing.T) {

	fmt.Println(Hello("Peter"))

	got := Hello("Chris")
	want := "Hello, Chris"

	if got != want {
		t.Errorf("got %q want %q", got, want)
	}
}

func TestAdder(t *testing.T) {
	sum := Add(2, 2)
	expected := 4

	fmt.Println(Add(2, 2))

	if sum != expected {
		t.Errorf("expected '%d' but got '%d'", expected, sum)
	}
}

func TestRepeat(t *testing.T) {
	repeated := Repeat("a")
	expected := "aaaaa"

	fmt.Println(Repeat(("a")))

	if repeated != expected {
		t.Errorf("expected %q but got %q", expected, repeated)
	}
}

func TestSum(t *testing.T) {

	numbers := [5]int{1, 2, 3, 4, 5}

	fmt.Println(Sum((numbers)))

	got := Sum(numbers)
	want := 15

	if got != want {
		t.Errorf("got %d want %d given, %v", got, want, numbers)
	}
}

func TestSumGroup(t *testing.T) {

	t.Run("collection of 5 numbers", func(t *testing.T) {
		numbers := [5]int{1, 2, 3, 4, 5}

		fmt.Println(Sum((numbers)))

		got := Sum(numbers)
		want := 15

		if got != want {
			t.Errorf("got %d want %d given, %v", got, want, numbers)
		}
	})

	t.Run("collection of any size", func(t *testing.T) {
		numbers := []int{1, 2, 3}

		fmt.Println(SumSlice(numbers))

		got := SumSlice(numbers)
		want := 6

		if got != want {
			t.Errorf("got %d want %d given, %v", got, want, numbers)
		}
	})

}
