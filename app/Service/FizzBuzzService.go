package service

import (
	fizzbuzz "github.com/YakovAkk/FizzBuzz/app/FizzBuzz"
)

type FizzBuzzService struct{}

func New() *FizzBuzzService {
	return &FizzBuzzService{}
}

func (this *FizzBuzzService) Calculate(usersNum int, rules []fizzbuzz.Rule) []string {
	fizzBuzz := fizzbuzz.New(rules)

	result := make([]string, usersNum)

	for i := 1; i <= usersNum; i++ {
		result[i-1] = fizzBuzz.Apply(i)
	}

	return result
}
