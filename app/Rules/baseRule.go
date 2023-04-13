package rules

type Rule interface {
	Apply(n int)
}
