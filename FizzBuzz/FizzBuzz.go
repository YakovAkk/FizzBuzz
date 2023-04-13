package fizzbuzz

import (
	"github.com/YakovAkk/FizzBuzz/rules"
)

type FizzBuzz struct {
	Rules []rules.Rule
}

func New(rules []rules.Rule) *FizzBuzz {
	return &FizzBuzz{Rules: rules}
}

func (this *FizzBuzz) Apply(n int) string {
	var result string = ""

	for i := 0; i < len(this.Rules); i++ {
		result += this.Rules[i].Apply(n)
	}

	if result == "" {
		return string(rune(n))
	}

	return result
}