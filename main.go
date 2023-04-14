package main

import (
	"github.com/YakovAkk/FizzBuzz/app/service"
)

func main() {

	var usersNum = 100

	myService := service.FizzBuzzService{}

	myService.Calculate(usersNum)
}
