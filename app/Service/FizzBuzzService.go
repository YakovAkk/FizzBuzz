package service

import (
	"fmt"

	fizzbuzz "github.com/YakovAkk/FizzBuzz/app/FizzBuzz"
	rules "github.com/YakovAkk/FizzBuzz/app/Rules"
)

type FizzBuzzService struct{}

func New() *FizzBuzzService {
	return &FizzBuzzService{}
}

func (this FizzBuzzService) calculate(usersNum int) {
	rule1 := rules.New("Fizz", 3)
	rule2 := rules.New("Buzz", 5)

	rules := []fizzbuzz.Rule{rule1, rule2}

	fizzBuzz := fizzbuzz.New(rules)

	for i := 1; i <= usersNum; i++ {
		fmt.Print(fizzBuzz.Apply(i) + " ")
	}
}
