package main

import (
	"fmt"

	"github.com/YakovAkk/FizzBuzz/fizzbuzz"
	"github.com/YakovAkk/FizzBuzz/rules"
)

func main() {
	var usersNum = 10

	rule1 := rules.Rule{Word: "Fizz", Number: 3}
	rule2 := rules.Rule{Word: "Buzz", Number: 5}

	fizzBuzz := fizzbuzz.FizzBuzz{Rules: []rules.Rule{rule1, rule2}}

	for i := 1; i <= usersNum; i++ {
		fmt.Print(fizzBuzz.apply(i))
	}
}
