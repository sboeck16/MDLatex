# LaTeX Writer for Golang

Project was created to transform markdown files into latex. This project provides:

* Very simple definitions for a latex document in golang.
* Provides String methods for different structs that result in LaTex strings.
* A markdown to latex parser is provided, see example.
* Parsing latex is not part of this repository!

## Usage latex structs

```
    // get default document
    testDoc := GetDefaultDoc()

	// get some formatting options (provide your own or from repo, or ommit)
	testDoc.GetDefaultsFromFile("test_opts.json")

    // add a section
	testSection := testDoc.AddSection(NewText("TestSection"))

    // add a subsection test section
	testSubSection1 := testSection.AddChild(BoldText("Sub1"))
    // add some text 
	testSubSection1.AddText(`Lorem ipsum dolor sit amet, consetetur sadipscing
elitr, sed diam nonumy eirmod tempor invidunt ut labore et dolore magna aliquyam
erat, sed diam voluptua. At vero eos et accusam et justo duo dolores et ea
rebum. Stet clita kasd gubergren, no sea takimata sanctus est Lorem ipsum dolor
sit amet. Lorem ipsum dolor sit amet, consetetur sadipscing elitr, sed diam
nonumy eirmod tempor invidunt ut labore et dolore magna aliquyam erat,
sed diam voluptua. At vero eos et accusam et justo duo dolores et ea rebum.
Stet clita kasd gubergren, no sea takimata sanctus est Lorem`)

    // add a subsubsection via child method
	testSubSection2 := testSection.AddChild(NewText("Sub2"))

    // add a list with items
	testList := testSubSection2.AddList("example list")
	testList.AddItem(NewText("item1"))
	testList.AddItem(NewText("item2"))
	testList.AddItem(NewText("item3"))

    // add a table with no specific name and no options
	table := testSection.AddTable("", nil)
    // set head
	table.AddHeaderStr("FOO", "BAR")
    // set body
	table.AddRowStr("A", "B")
	table.AddRowStr("C", "D")

    // print resulting latex document
	fmt.Println(testDoc.String())
```

## Markdown converter

```
    lines := strings.Split(getExampleMD(), "\n")
    doc := GetDefaultDoc()
    WriteMDToLatexDoc(lines, doc)
    fmt.Println(doc.String())
```
