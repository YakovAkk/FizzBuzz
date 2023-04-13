package ruleBase

type Rule interface{
	Word   string
	Number int
	Apply(n int)
}