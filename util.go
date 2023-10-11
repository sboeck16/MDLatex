package latex

import (
	"fmt"
	"regexp"
	"strings"
)

var (
	removeWhiteSpaces = regexp.MustCompile(`^\s*(.*?)\s*$`)
)

/*
Brackets puts the text in latex standard brackets
*/
func Brackets(text string) string {
	return OpenBr + text + CloseBr
}

/*
JoinOpts joins options with delim and brackets. If opts is empty an empty
string will be returned
*/
func JoinParams(params ...string) string {
	if len(params) > 0 {
		return ParamsBegin + strings.Join(params, ParamsDelim) + ParamsEnd
	}
	return ""
}

/*
Block wrap the text with a \begin{blocktype} and \end{blocktype}.
*/
func Block(blocktype, text string) string {
	ret := Begin + Brackets(blocktype) + LineBr
	ret += CheckAddLineBreak(text)
	ret += End + Brackets(blocktype) + LineBr
	return ret
}

/*
CheckAddLineBreak will ensure last character in text is a line break.
*/
func CheckAddLineBreak(text string) string {
	if text != "" && text[len(text)-1:] == LineBr {
		return text
	}
	return text + LineBr
}

/*
EnsureSpaceAfter
*/
func EnsureSpaceAfter(text string) string {
	if text != "" && text[len(text)-1:] == WhiteSpace {
		return text
	}
	return text + WhiteSpace
}

/*
EnsureSpaceBefore
*/
func EnsureSpaceBefore(text string) string {
	if text != "" && text[0:1] == WhiteSpace {
		return text
	}
	return WhiteSpace + text

}

/*
TrimWhiteSpace trims all whitespace at the begin and end of text
*/
func TrimWhiteSpace(text string) string {
	return removeWhiteSpaces.ReplaceAllString(text, `$1`)
}

/*
WhiteSpaceWrap ensures that text are wrapped in whitespaces
*/
func WhiteSpaceWrap(text string) string {
	return EnsureSpaceAfter(EnsureSpaceBefore(text))
}

// utility debug printer...
func deb(i ...any) {
	fmt.Println(i...)
}
