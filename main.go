package main

import (
	"fmt"

	"github.com/YakovAkk/FizzBuzz/fizzbuzz"
	"github.com/YakovAkk/FizzBuzz/rules"
)

func main() {
	var usersNum = 10

	rule1 := rules.New("Fizz", 3)
	rule2 := rules.New("Buzz", 5)

	rules := []rules.Rule{*rule1, *rule2}

	fizzBuzz := fizzbuzz.New(rules)

	for i := 1; i <= usersNum; i++ {
		fmt.Print(fizzBuzz.Apply())
	}
}
