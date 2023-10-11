package latex

import "strings"

/*
Table holds data and options to build a table
*/
type Table struct {
	OptionHolder
	// Name of table (maybe make it displayable)
	Name string
	// Options holds table options
	Options string
	// Rows holds the talbe data, first row on index 0 is the table head.
	Rows [][]LatexStr
	// Horizontal Line Options
	NoOuterHLine    bool
	NoInnerHLine    bool
	NoHeadBodyHLine bool
	NoDoubleHLine   bool
}

/*
NewTable constructor for table with latex table options
*/
func NewTable(name string, opts string) *Table {
	return &Table{Options: opts, Name: name}
}

/*
AddHeader shifts a Row into first place that is used as header. Header
is first row and can be alternativly set via first call of AddRow.
*/
func (t *Table) AddHeader(header []LatexStr) {
	t.Rows = append([][]LatexStr{header}, t.Rows...)
}

/*
AddRow appends a Row to Table.
*/
func (t *Table) AddRow(row []LatexStr) {
	t.Rows = append(t.Rows, row)
}

/*
AddHLine adds an empty row which leads to an additional hline
*/
func (t *Table) AddHLine() {
	t.AddRowStr()
}

/*
String renders table and make its printable
*/
func (t *Table) String() string {
	// some sanity checks
	// content is here
	if len(t.Rows) == 0 {
		return ""
	}
	// broken table head
	if len(t.Rows[0]) == 0 {
		return ""
	}
	// no options -> use default
	if t.Options == "" {
		defOpts := make([]string, len(t.Rows[0]))
		for ind := range t.Rows[0] {
			defOpts[ind] = DefaultTabularOption
		}
		t.Options = TableOptionsWrap + strings.Join(
			defOpts, TableOptionsWrap) + TableOptionsWrap
	}

	// Table start
	ret := ""
	ret += Begin + Brackets(Tabular)
	ret += Brackets(t.Options) + LineBr
	// table is wrapped with lines.
	if !t.NoOuterHLine {
		ret += HLine
	}

	// build table, skipping rows with mismatching length
	for rowInd, row := range t.Rows {
		elemStr := make([]string, 0)
		for _, elem := range row {
			elemStr = append(elemStr, elem.String())
		}
		if len(elemStr) > 0 {
			ret += strings.Join(elemStr, TableCellDelim) + TableRowDelim
		}

		// Horizontal LINES
		// ----------------
		if !t.NoInnerHLine && rowInd < len(t.Rows)-1 && rowInd > 0 {
			ret += HLine
		}
		if !t.NoOuterHLine && rowInd == len(t.Rows)-1 {
			ret += HLine
		}
		// Add another hline for dividing head from body
		if rowInd == 0 && !t.NoHeadBodyHLine {
			ret += HLine
			if !t.NoDoubleHLine {
				ret += HLine
			}
		}
	}

	// Table end
	ret += End + Brackets(Tabular) + LineBr

	return t.WrapWithOptions(ret)
}

// #############################################################################
// #							Util/Wrapper
// #############################################################################

/*
AddRowStr wraps input strings in text and adds them as row.
*/
func (t *Table) AddRowStr(strs ...string) {
	lStr := []LatexStr{}
	for _, str := range strs {
		lStr = append(lStr, NewText(str))
	}
	t.AddRow(lStr)
}

/*
AddHeaderStr adds string headline and sets them to bold.
*/
func (t *Table) AddHeaderStr(strs ...string) {
	lStr := []LatexStr{}
	for _, str := range strs {
		lStr = append(lStr, BoldText(str))
	}
	t.AddHeader(lStr)
}
