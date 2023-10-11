package latex

// TOC describes table of contents
type TOC struct {
	NoNewPage bool
}

// NewTOC is constructor for table of contents
func NewTOC() *TOC {
	return &TOC{}
}

// String is output/interface function to print contents in latex
func (toc *TOC) String() string {
	ret := TableOfContents + LineBr
	if !toc.NoNewPage {
		ret += NewPage
	}
	return ret
}
