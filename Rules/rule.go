package dividedrules

type DividedRule struct {
	Word   string
	Number int
}

func (this *DividedRule) Apply(n int) string {
	if n%this.Number == 0 {
		return this.Word
	}
	return ""
}

func New(word string, num int) *DividedRule {
	return &DividedRule{Word: word, Number: num}
}
