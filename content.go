package latex

import "fmt"

/*
LatexStr is the interface for stringifiable latex objects
*/
type LatexStr interface {
	String() string
}

/*
Container for multiple strings. Parent for sections and document
*/
type ContentHolder struct {
	OptionHolder
	Content []LatexStr
}

func (ch *ContentHolder) ContentString() string {
	ret := ""
	// empty
	if ch.Content == nil {
		return ret
	}

	// add contents
	for _, cont := range ch.Content {
		ret += CheckAddLineBreak(cont.String())
	}

	// add options
	ret = ch.WrapWithOptions(ret)

	return ret
}

// #############################################################################
// 							GENERIC Content
// #############################################################################

/*
AppendContent adds to the container. Must be stringifiable items
*/
func (ch *ContentHolder) AppendContent(addCont ...LatexStr) {
	ch.Content = append(ch.Content, addCont...)
}

/*
ShiftContent adds to the container. Elements ared added in front.
Must be stringifiable items.
*/
func (ch *ContentHolder) ShiftContent(addCont ...LatexStr) {
	ch.Content = append(addCont, ch.Content...)
}

/*
adds node without checking if lvl is defined
*/
func (ch *ContentHolder) addNewNode(lvl int, hl LatexStr) *Node {
	ret := NewNode(lvl, hl)

	ret.SetNameTypeAndParent(hl.String(), lvlToName[lvl], &ch.OptionsByName)
	ch.AppendContent(ret)
	return ret
}

// #############################################################################
// #							Wrapper and Utility add
// #############################################################################

/*
AddNode adds a node like chapter or section and returns it. If lvl constant
is unknown an error will be returned.
*/
func (ch *ContentHolder) AddNode(lvl int, hl LatexStr) (*Node, error) {
	if _, ok := lvlToName[lvl]; !ok {
		return nil, fmt.Errorf("unknown chapter/section/level")
	}
	return ch.addNewNode(lvl, hl), nil
}

/*
AddNodeLvlName for human readable section string name/constant
*/
func (ch *ContentHolder) AddNodeLvlName(lvl string, hl LatexStr) (*Node, error) {
	if _, ok := nameToLvl[lvl]; !ok {
		return nil, fmt.Errorf("unknown chapter/section/level")
	}
	return ch.AddNode(nameToLvl[lvl], hl)
}

/*
AddPart adds a part with given headline
*/
func (ch *ContentHolder) AddPart(hl LatexStr) *Node {
	return ch.addNewNode(nameToLvl[Part], hl)
}

/*
AddChapter adds a chapter with given headline
*/
func (ch *ContentHolder) AddChapter(hl LatexStr) *Node {
	return ch.addNewNode(nameToLvl[Chapter], hl)
}

/*
AddSection
*/
func (ch *ContentHolder) AddSection(hl LatexStr) *Node {
	return ch.addNewNode(nameToLvl[Section], hl)
}

/*
AddSubSection
*/
func (ch *ContentHolder) AddSubSection(hl LatexStr) *Node {
	return ch.addNewNode(nameToLvl[SubSection], hl)
}

/*
AddSubSubSection
*/
func (ch *ContentHolder) AddSubSubSection(hl LatexStr) *Node {
	return ch.addNewNode(nameToLvl[SubSubSection], hl)
}

/*
AddParagraph
*/
func (ch *ContentHolder) AddParagraph(hl LatexStr) *Node {
	return ch.addNewNode(nameToLvl[Paragraph], hl)
}

/*
AddSubParagraph
*/
func (ch *ContentHolder) AddSubParagraph(hl LatexStr) *Node {
	return ch.addNewNode(nameToLvl[SubSection], hl)
}

/*
AddText
*/
func (ch *ContentHolder) AddText(text string) *Text {
	ret := NewText(text)
	ret.SetNameTypeAndParent("", TextType, &ch.OptionsByName)
	ch.Content = append(ch.Content, ret)
	return ret
}

/*
AddRaw adds raw latex commands that are not escaped
*/
func (ch *ContentHolder) AddRaw(text string) *Text {
	ret := NewTextRaw(text)
	ret.SetNameTypeAndParent("", TextType, &ch.OptionsByName)
	ch.Content = append(ch.Content, ret)
	return ret
}

/*
AddList
*/
func (ch *ContentHolder) AddList(name string) *List {
	ret := NewList(name)
	ret.SetNameTypeAndParent(name, ListType, &ch.OptionsByName)
	ch.Content = append(ch.Content, ret)
	return ret
}

/*
AddTable
*/
func (ch *ContentHolder) AddTable(name string, opts string) *Table {
	ret := NewTable(name, opts)
	ret.SetNameTypeAndParent(name, TableType, &ch.OptionsByName)
	ch.Content = append(ch.Content, ret)
	return ret
}

/*
AddTableNoBorder
*/
func (ch *ContentHolder) AddTableNoBorder(name string, opts string) *Table {
	ret := NewTable(name, opts)
	ret.SetNameTypeAndParent(name, TableType, &ch.OptionsByName)
	ret.NoHeadBodyHLine = true
	ret.NoOuterHLine = true
	ret.NoInnerHLine = true
	ch.Content = append(ch.Content, ret)
	return ret
}
