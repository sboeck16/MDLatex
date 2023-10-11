package latex

// Node holds a part of the document. Node can be part of another Node
// depending how the document is build. All Nodes must be linked to Doc.
// And Doc could be a node as well, but we want to keep package etc. simple
type Node struct {
	// ParentClass for content managing
	ContentHolder
	/* Level for latex sectioning:
	Part:          0,
	Chapter:       1,
	Section:       2,
	SubSection:    3,
	SubSubSection: 4,
	Paragraph:     5,
	SubParagraph:  6,
	*/
	Level int
	// Headline to be printed for this node
	Headline LatexStr
}

/*
NewSection is the constructor for a section. This section is unlinked from
Document!
*/
func NewNode(l int, hl LatexStr) *Node {
	return &Node{Level: l, Headline: hl}
}

/*
AddChild will return a child one latex depth below this node. If no depth below
is defined depth for child node will be the same as for parent.
*/
func (n *Node) AddChild(hl LatexStr) *Node {
	ret, err := n.AddNode(n.Level+1, hl)
	if err != nil {
		ret, _ = n.AddNode(n.Level, hl)
	}
	return ret
}

/*
String stringifies node to be printed into latex document
*/
func (n *Node) String() string {
	ret := ""

	// HeadLine
	ret += LC + lvlToName[n.Level] + Brackets(n.Headline.String()) + LineBr
	for _, opt := range n.GetHeadlineOptions() {
		ret = opt.WrapString(ret)
	}

	// Content
	ret += n.ContentString()

	// wrap all with outer options if any
	for _, opt := range n.GetOuterOptions() {
		ret = opt.WrapString(ret)
	}

	return ret
}
