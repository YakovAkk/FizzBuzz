package main

import (
	"fmt"

	fizzbuzz "github.com/YakovAkk/FizzBuzz/app/FizzBuzz"
	"github.com/YakovAkk/FizzBuzz/app/rules"
)

func main() {
	var usersNum = 15
	rule1 := rules.New("Fizz", 3)
	rule2 := rules.New("Buzz", 5)

	rules := []rules.DividedRule{*rule1, *rule2}

	fizzBuzz := fizzbuzz.New(rules)

	for i := 1; i <= usersNum; i++ {
		fmt.Print(fizzBuzz.Apply(i) + " ")
	}
}
