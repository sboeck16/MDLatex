package latex

import (
	"regexp"
	"strings"
)

const (
	// Lvl1Headline for biggest import headline in markdown `# Headline`
	Lvl1Headline = "#"
	// Lvl2Headline
	Lvl2Headline = "##"
	// Lvl3Headline
	Lvl3Headline = "###"
	// Lvl4Headline
	Lvl4Headline = "####"
	// Lvl5Headline is the smallest supported. Any more `#` will use this.
	Lvl5Headline = "#####"

	mdTableSep = `|`
)

var (
	// HeadlineToLatexSection holds amount of # to section name in latex doc
	HeadlineToLatexSection = map[string]string{
		Lvl1Headline: Section,
		Lvl2Headline: SubSection,
		Lvl3Headline: SubSubSection,
		Lvl4Headline: Paragraph,
		Lvl5Headline: SubParagraph,
	}
	// parsing regexes
	readHeadline = regexp.MustCompile(`^\s*(#+)\s+(.*)$`)
	readTable    = regexp.MustCompile(`^\s*\|`)
	readSep      = regexp.MustCompile(`^\s*\|-`) // we only need to identify
	readList     = regexp.MustCompile(`^\*\s(.*)$`)
	readEmpty    = regexp.MustCompile(`^\s*$`)
	readEmptyRow = regexp.MustCompile(`^[| ]+$`)
	readNewPage  = regexp.MustCompile(`^\s*!newpage`)

	// makro regex
	makroNextTable = regexp.MustCompile(`§§table§([^§]*)§([^§]*)§([^§]*)$`)
)

/*
TransformMDToLatex is a quick and dirty transformer. it adds all section and
subsection to document amd tries to get a parent node. provide markdown by
lines without line breaks.
*/
func WriteMDToLatexDoc(lines []string, document *Doc) {

	var lastTable *Table
	var lastList *List
	var lastNode *Node
	lastNode = &document.Node
	nodeParents := map[int]*Node{0: lastNode}
	var lastItem *Text
	nextTableOpts := []string{"", "", "", ""}

	for _, line := range lines {

		// add a new page
		if readNewPage.MatchString(line) {
			lastNode.AddRaw(NewPage)
			continue
		}

		// reading a makro line
		if m := makroNextTable.FindStringSubmatch(line); len(m) > 2 {
			nextTableOpts = m
			continue
		}

		// Empty line ends table and list or paragraph (Maybe use paragraph here?)
		if readEmpty.MatchString(line) {
			lastTable = nil
			lastList = nil
			// add empty line
			lastNode.AddText("")
			continue
		}

		//		 TABLE
		// -----------------
		// markdown table seperator holds no value here
		// Table is directly inserted to document and will not use last
		// nodes options!
		if readSep.MatchString(line) {
			continue
		}
		// we read a table line
		if matches := readTable.FindStringSubmatch(line); len(matches) > 0 {
			// havent started a table yet
			tableStrs := []string{}
			for _, elem := range strings.Split(line, mdTableSep) {
				if elem == "" {
					continue
				}
				tableStrs = append(tableStrs, TrimWhiteSpace(elem))
			}
			if lastTable == nil {
				lastTable = lastNode.AddTable(
					nextTableOpts[1], nextTableOpts[2])
				if nextTableOpts[3] != "" {
					lastTable.ShortOpt(nextTableOpts[3])
				}
				nextTableOpts = []string{"", "", "", ""}

				lastTable.AddHeaderStr(tableStrs...)
			} else {
				if !readEmptyRow.MatchString(line) {
					lastTable.AddRowStr(tableStrs...)
				} else {
					lastTable.AddHLine()
				}

			}
			// line is used continue
			continue
		} else {
			// safeguard for ending table (md tables should be ended by empty
			// line)
			lastTable = nil
		}

		// 		LIST
		// --------------
		// we have no safeguard for ending list as they contain all to last
		// item until an empty line is reached
		if matches := readList.FindStringSubmatch(line); len(matches) > 0 {
			if lastList == nil {
				lastList = lastNode.AddList("")
			}
			lastItem = NewText(matches[1])
			lastList.AddItem(lastItem)
			continue
		}
		if lastList != nil {
			lastItem.AddText(line)
			continue
		}

		//		Sections
		// --------------------
		if matches := readHeadline.FindStringSubmatch(line); len(matches) > 0 {
			// More ##### than defined, MAYBE add an error/logging/else channel
			if len(matches[1]) > len(HeadlineToLatexSection) {
				continue
			}
			for i := len(matches[1]) - 1; i >= 0; i-- {
				if val, ok := nodeParents[i]; ok {
					lastNode, _ = val.AddNodeLvlName(
						HeadlineToLatexSection[matches[1]],
						NewText(matches[2]))
					for a := len(matches[1]); a <= 5; a++ {
						nodeParents[a] = lastNode
					}

					break
				}
			}
			continue
		}

		// 		Text
		// ---------------
		lastNode.AddText(line)
	}
}
