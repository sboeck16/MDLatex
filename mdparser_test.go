package latex

import (
	"strings"
	"testing"
)

var (
	runExample = true
)

/*
More an example than a test. just run everything thorugh here and check it
with hardcoded output...
*/
func TestParser(t *testing.T) {
	lines := strings.Split(getExampleMD(), "\n")
	doc := GetDefaultDoc()
	WriteMDToLatexDoc(lines, doc)
	if doc.String() != expectedResult() {
		t.Error("rendered latex mismatched expected result")
		deb(doc.String())
	}
}

func getExampleMD() string {
	return `# Test Markdown
foobar

## another section

foobar
foobar
foobar

### table

§§table§TestTable§|m{2cm}|m{3cm}|m{4cm}|§small
| H1 | H2 | H3 |
|----|----|----|
| 1 | 2 | 3 |
| A | B | C |
| | | |
| X | Y | Z |

### list

* abcd
* de
fg
* xyz

not part of list

## more foobar

foobar
foobar
foobar

### baz

baz
baz
`
}

func expectedResult() string {
	return `\documentclass{article}
\usepackage[utf8]{inputenc}
\usepackage[T1]{fontenc}
\usepackage[ngerman]{babel}
\usepackage[a4paper, left=2cm, bottom=15mm, top=2cm]{geometry}
\usepackage{hyperref}
\usepackage{multicol}
\usepackage{array}
\begin{document}
\tableofcontents
\newpage
\section{Test Markdown}
foobar

\subsection{another section}

foobar
foobar
foobar

\subsubsection{table}

\begin{small}
\begin{tabular}{|m{2cm}|m{3cm}|m{4cm}|}
\hline
\textbf{H1}&\textbf{H2}&\textbf{H3}\\
\hline
\hline
1&2&3\\
\hline
A&B&C\\
\hline
\hline
X&Y&Z\\
\hline
\end{tabular}
\end{small}

\subsubsection{list}

\begin{itemize}
\item abcd
\item de fg
\item xyz
\end{itemize}

not part of list

\subsection{more foobar}

foobar
foobar
foobar

\subsubsection{baz}

baz
baz

\end{document}
`
}
