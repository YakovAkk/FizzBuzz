package tests

import (
	"reflect"
	"testing"

	fizzbuzz "github.com/YakovAkk/FizzBuzz/app/FizzBuzz"
	rules "github.com/YakovAkk/FizzBuzz/app/Rules"
	service "github.com/YakovAkk/FizzBuzz/app/Service"
)

func testIfNumber3(t *testing.T) {
	//Arrange
	expectedResult := []string{"1", "2", "Fizz"}

	Number := 3

	rule1 := rules.New("Fizz", 3)
	rule2 := rules.New("Buzz", 5)
	rules := []fizzbuzz.Rule{rule1, rule2}

	myService := service.FizzBuzzService{}
	// Act
	result := myService.Calculate(Number, rules)

	// Assert
	if reflect.DeepEqual(result, expectedResult) {
		t.Error("Incorrect data when num = 3")
	} else {
		t.Logf("PASS")
	}
}

func testIfNumber15(t *testing.T) {
	//Arrange
	expectedResult := []string{"1", "2", "Fizz", "4", "Buzz", "Fizz", "7", "8",
		"Fizz", "Buzz", "11", "Fizz", "13", "14", "FizzBuzz"}

	Number := 3

	rule1 := rules.New("Fizz", 3)
	rule2 := rules.New("Buzz", 5)
	rules := []fizzbuzz.Rule{rule1, rule2}

	myService := service.FizzBuzzService{}
	// Act
	result := myService.Calculate(Number, rules)

	// Assert
	if reflect.DeepEqual(result, expectedResult) {
		t.Error("Incorrect data when num = 3")
	} else {
		t.Logf("PASS")
	}
}
