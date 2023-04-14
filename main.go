package main

import (
	"fmt"

	fizzbuzz "github.com/YakovAkk/FizzBuzz/app/FizzBuzz"
	rules "github.com/YakovAkk/FizzBuzz/app/Rules"
	"github.com/YakovAkk/FizzBuzz/app/service"
)

func main() {

	var usersNum = 15

	myService := service.FizzBuzzService{}

	rule1 := rules.New("Fizz", 3)
	rule2 := rules.New("Buzz", 5)
	rules := []fizzbuzz.Rule{rule1, rule2}

	res := myService.Calculate(usersNum, rules)

	for _, el := range res {
		fmt.Printf(el + " ")
	}
}
