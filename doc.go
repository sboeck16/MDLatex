package latex

// Doc is the latex document
type Doc struct {
	Node
	class    string
	packages []*LatexPackage
	toc      *TOC
}

// NewDoc creates a new document. Class needs to be specified
func NewDoc(class string) *Doc {
	return &Doc{class: class}
}

// AddPackage adds a package to document with options options
func (d *Doc) AddPackage(name string, ops []string) {
	d.packages = append(d.packages, NewLatexPackage(name, ops))
}

// AddTOC adds table of contents
func (d *Doc) AddTOC() {
	d.toc = NewTOC()
}

// String is output/interface function to print contents in latex
func (d *Doc) String() string {

	ret := ""

	// add class
	ret += DocumentClass + Brackets(d.class) + LineBr

	// add packages
	for _, pack := range d.packages {
		ret += pack.String()
	}

	// Add TOC
	toc := ""
	if d.toc != nil {
		toc = d.toc.String()
	}

	// add content in document block
	ret += Block(Document, toc+d.ContentString())

	// return latex document string
	return ret
}

// #############################################################################
// #							Utility
// #############################################################################

/*
GetDefaultDoc returns doc with packages i find useful
*/
func GetDefaultDoc() *Doc {
	ret := NewDoc("article")
	ret.SetDefaultPackages()
	return ret
}

func (d *Doc) SetDefaultPackages() {
	d.AddPackage("inputenc", []string{"utf8"})
	d.AddPackage("fontenc", []string{"T1"})
	d.AddPackage("babel", []string{"ngerman"})
	d.AddPackage("geometry", []string{
		"a4paper", "left=2cm", "bottom=15mm", "top=2cm"})
	d.AddPackage("hyperref", []string{})
	d.AddPackage("multicol", []string{})
	d.AddPackage("array", []string{})
	d.AddTOC()
}
