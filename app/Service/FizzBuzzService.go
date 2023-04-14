package service

import (
	"fmt"

	fizzbuzz "github.com/YakovAkk/FizzBuzz/app/FizzBuzz"
)

ype FizzBuzzService struct{}

unc New() *FizzBuzzService {
	return &FizzBuzzService{}
}

unc (this FizzBuzzService) Calculate(usersNum int, rules []fizzbuzz.Rule) {

	fizzBuzz := fizzbuzz.New(rule)

for i := 1; i <= usersNum; i++ {
		fmt.Print(fizzBuzz.Apply(i) + " ")
}

