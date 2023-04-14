package main

import (
	"github.com/YakovAkk/FizzBuzz/app/service"
)

func main() {

	var usersNum = 100

	service := service.New()

	service.FizzBuzzService.calculate(usersNum)
}
