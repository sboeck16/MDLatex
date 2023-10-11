package latex

/*
List holds an itemize latex list
*/
type List struct {
	OptionHolder
	// Items for printing
	Items []LatexStr
	// Name holds internal name, maybe make it displayable?
	Name string
}

/*
NewList, constructor method
*/
func NewList(name string) *List {
	return &List{Name: name}
}

/*
AddItem adds a printable list item
*/
func (l *List) AddItem(item LatexStr) {
	l.Items = append(l.Items, item)
}

/*
String stringifies the list
*/
func (l *List) String() string {
	ret := ""
	for _, elem := range l.Items {
		ret += Item
		// wrap options here
		ret +=  EnsureSpaceBefore(l.WrapWithOptions(elem.String()))
		ret += LineBr
	}
	// doesnt work
	// ret = l.WrapWithOptions(ret)
	return Block(Itemize, ret)
}
