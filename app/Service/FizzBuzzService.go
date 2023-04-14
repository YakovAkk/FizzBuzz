package service

import (
	"fmt"

	fizzbuzz "github.com/YakovAkk/FizzBuzz/app/FizzBuzz"
)

type FizzBuzzService struct{}

func New() *FizzBuzzService {
	return &FizzBuzzService{}
}

func (this FizzBuzzService) Calculate(usersNum int, rules []fizzbuzz.Rule) {

	fizzBuzz := fizzbuzz.New(rules)

	for i := 1; i <= usersNum; i++ {
		fmt.Print(fizzBuzz.Apply(i) + " ")
	}

}