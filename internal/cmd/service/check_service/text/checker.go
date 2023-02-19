package text

type TextChecker struct {
	text string
}

func NewTextChecker(text string) *TextChecker {
	return &TextChecker{
		text: text,
	}
}

func (t TextChecker) Check() bool {
	return t.text == "ok"
}
