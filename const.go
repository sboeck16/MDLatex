package latex

const (
	// Symbols
	// ------

	// line break. windows builds may need ann additional `\r`?
	LineBr = "\n"
	// Space
	WhiteSpace = ` `
	// LC holds the begin of a Latex Command
	LC = `\`
	// OptBegin and OptEnd for options in `[]`
	ParamsBegin = `[`
	ParamsEnd   = `]`
	// Option delimiter
	ParamsDelim = `, `
	// Latex uses curly brackets in many places. If needed this constants
	// needs to be split/exported/redone?
	OpenBr  = `{`
	CloseBr = `}`

	// Control
	// -------
	// Begin LaTeX keyword
	Begin = LC + `begin`
	// End
	End = LC + `end`

	// Document
	// --------
	// DocumentClass
	DocumentClass = LC + `documentclass`
	// TOC
	TableOfContents = LC + `tableofcontents`
	// Package
	UsePackage = LC + `usepackage`
	// Document name for document block
	Document = `document`

	// List
	// ----
	// Itemize latex comand
	Itemize = `itemize`
	// Item single latex item
	Item = LC + `item`

	// Table
	// -----
	// Tabular is used for tables (deprecated?)
	Tabular = `tabular`
	// TabularX
	TabularX = `tabularx`
	// TableName
	TableName = `table`
	// TableCaption
	TableCaption = `caption`
	// TableHere
	TableHere = `h`
	// TextWidthOpt to ensure table stayson page
	TextWidthOpt = "textwidth"
	// DefaultTabularXOption
	DefaultTabularXOption = `X`

	// TableOptionsWrap holds Latex symbol to describe the tabular options
	TableOptionsWrap = `|`
	// DefaultTabularOption
	DefaultTabularOption = `m`
	// CenterTabularOpt
	CenterTabularOpt = `c`
	// LeftTabularOpt
	LeftTabularOpt = `l`
	// RightTabularOpt
	RightTabularOpt = `r`
	// TableElemDelimiter
	TableCellDelim = `&`
	// TableRowDelim (with added linebreak for visibility)
	TableRowDelim = `\\` + LineBr

	// HLine for horizontal line (with added line break)
	HLine = LC + `hline` + LineBr

	// Formatting
	// ----------
	Bold = `textbf`
	// size
	Small  = `small`
	Normal = `Normal`
	// behaviour
	NewPage = LC + `newpage` + LineBr

	// Section, parts, etc
	// -------------------
	Part          = `part`
	Chapter       = `chapter`
	Section       = `section`
	SubSection    = `subsection`
	SubSubSection = `subsubsection`
	Paragraph     = `paragraph`
	SubParagraph  = `subparagraph`
	TableType     = `tabular`
	TextType      = `text`
	ListType      = `itemize`
	DocumentType  = `document`
)

var (
	// Identify mapping. If ever needed a setter/getter might be added
	lvlToName = map[int]string{
		0: Part,
		1: Chapter,
		2: Section,
		3: SubSection,
		4: SubSubSection,
		5: Paragraph,
		6: SubParagraph,
	}
	nameToLvl = map[string]int{
		Part:          0,
		Chapter:       1,
		Section:       2,
		SubSection:    3,
		SubSubSection: 4,
		Paragraph:     5,
		SubParagraph:  6,
	}
)
