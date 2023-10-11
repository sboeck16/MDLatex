package latex

import "regexp"

var (
	escapeLatex  = regexp.MustCompile(`(&|%|\$|#|_|{|}|~|\^)`)
	escapeQuotes = regexp.MustCompile(`"`)
)

/*
Text holds basic text strings to use in latex. Escapes text when printed
*/
type Text struct {
	OptionHolder
	text string
	// If set can be used to add direct latex cmds
	IsRaw bool
}

/*
NewText returns a Tex
*/
func NewText(text string) *Text {
	return &Text{text: text}
}

func NewTextRaw(text string) *Text {
	return &Text{text: text, IsRaw: true}
}

/*
AddText adds to existing text. It ensures a whitespace between old and new text
*/
func (t *Text) AddText(add string) {
	t.text = EnsureSpaceAfter(t.text) + add
}

/*
String returns printable text.
*/
func (t Text) String() string {
	if t.IsRaw {
		return t.text
	}
	// clear unwanted characters
	ret := escapeLatex.ReplaceAllString(t.text, `\$1`)
	// quotes
	ret = escapeQuotes.ReplaceAllString(ret, `''`)
	return t.WrapWithOptions(ret)
}

/*
Makes a text bold. Used to set Header or Headlines
*/
func BoldText(text string) *Text {
	ret := NewText(text)
	ret.AddOption(Bold, nil, nil, true, "", false)
	return ret
}
