package integers

import (
	"fmt"
	"reflect"
	"testing"
)

func TestAdder(t *testing.T) {
	got := Add(4, 5)
	want := 9

	if got != want {
		t.Errorf("got %d want %d", got, want)
	}
}

func TestSum(t *testing.T) {
	numbers := []int{1, 2, 3, 4, 5}

	got := Sum(numbers)
	want := 15

	if got != want {
		t.Errorf("got %d want %d", got, want)
	}
}

func TestSumAll(t *testing.T) {

	number1 := []int{1, 2, 3}
	number2 := []int{4, 5, 6}

	got := SumAll(number1, number2)
	want := []int{6, 15}
	isEqual := reflect.DeepEqual(got, want)
	if !isEqual {
		t.Errorf("got %v want %v", got, want)
	}

}

func TestSumAllTails(t *testing.T) {
	number1 := []int{1, 2}
	number2 := []int{2, 3}

	got := SumAllTails(number1, number2)
	want := []int{2, 3}

	isEqual := reflect.DeepEqual(got, want)
	if !isEqual {
		t.Errorf("got %v want %v", got, want)
	}
}

func ExampleAdd() {
	sum := Add(1, 5)
	fmt.Println(sum)
	// Output: 6
}

func ExampleSum() {
	numbers := []int{1, 2, 3, 4, 5}
	sum := Sum(numbers)
	fmt.Println(sum)
	// Output: 15
}

func ExampleSumAll() {
	number1 := []int{1, 2}
	number2 := []int{2, 3}

	sum := SumAll(number1, number2)
	fmt.Println(sum)
	// Output  : [3, 5]

}

func ExampleSumAllTails() {
	number1 := []int{1, 2}
	number2 := []int{2, 3}

	sum := SumAllTails(number1, number2)
	fmt.Println(sum)
	// Output: [2, 3]
}
