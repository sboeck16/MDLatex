package latex

// LatexPackage holds definition for latex package with options
type LatexPackage struct {
	name    string
	options []string
}

// NewLatexPackage is the constructor, better use AddPackage for Doc
func NewLatexPackage(n string, ops []string) *LatexPackage {
	return &LatexPackage{name: n, options: ops}
}

// String is output/interface function to print contents in latex
func (lp *LatexPackage) String() string {
	ret := UsePackage
	ret += JoinParams(lp.options...)
	ret += Brackets(lp.name) + LineBr
	return ret
}
