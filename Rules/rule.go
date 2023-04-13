package rules

type Rule struct {
	Word   string
	Number int
}

func (this Rule) Apply(n int) string {
	if n%this.Number == 0 {
		return this.Word
	}
	return ""
}

func New(word string, num int) *Rule {
	return &Rule{Word: word, Number: num}
}
