package fizzbuzz

import (
	"strconv"
)

type Rule interface {
	Apply(n int) string
}

type FizzBuzz struct {
	Rules []Rule
}

func New(rules []Rule) *FizzBuzz {
	return &FizzBuzz{Rules: rules}
}

func (this FizzBuzz) Apply(n int) string {
	var result string = ""

	for i := 0; i < len(this.Rules); i++ {
		result += this.Rules[i].Apply(n)
	}

	if result == "" {
		return strconv.FormatInt(int64(n), 10)
	}

	return result
}
