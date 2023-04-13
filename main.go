package main

import (
	"fmt"

	fizzbuzz "github.com/YakovAkk/FizzBuzz/FizzBuzz"
	dividedrules "github.com/YakovAkk/FizzBuzz/dividedrules"
)

func main() {
	var usersNum = 15

	rule1 := dividedrules.New("Fizz", 3)
	rule2 := dividedrules.New("Buzz", 5)

	rules := []dividedrules.DividedRule{*rule1, *rule2}

	fizzBuzz := fizzbuzz.New(rules)

	for i := 1; i <= usersNum; i++ {
		fmt.Print(fizzBuzz.Apply(i) + " ")
	}
}
