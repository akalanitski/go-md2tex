package markdown

import (
	"regexp"
	"strings"
)

// cleanup replace sequnces of spaces to one space and sequence of EOL more the
// 2 in a row by 2 line breaks
func cleanup(input string) string {
	multiSpace := regexp.MustCompile(`\s+`)
	multiBreak := regexp.MustCompile(`\n{3,}`)
	ouput := multiBreak.ReplaceAllString(input, "\n\n")
	output := multiSpace.ReplaceAllString(ouput, " ")
	return output
}

func escapeTex(input string) string {
	e := regexp.MustCompile(`[\{}#%&_^~]`)
	return e.ReplaceAllString(input, "\\$0")
}

func MarkdownToTex(input string) string {
	input = cleanup(input)
	input = escapeTex(input)
	boldItalicPattern := regexp.MustCompile(`[*_]{3}(.+?)[*_]{3}`)
	boldPattern := regexp.MustCompile(`[*_]{2}(.+?)[*_]{2}`)
	italicPattern := regexp.MustCompile(`[*_]{1}(.+?)[*_]{1}`)
	codePattern := regexp.MustCompile("`(.+?)`")
	urlPattern := regexp.MustCompile(`\[(.+?)\]\((.+?)\)`)

	output := input
	output = boldItalicPattern.ReplaceAllString(output, "\\textbf{\\textit{$1}}")
	output = boldPattern.ReplaceAllString(output, "\\textbf{$1}")
	output = italicPattern.ReplaceAllString(output, "\\textit{$1}")
	output = codePattern.ReplaceAllString(output, "\\texttt{$1}")
	output = urlPattern.ReplaceAllString(output, "\\href{$2}{$1}")
	return output
}

func MarkdownToHTML(input string) string {
	input = cleanup(input)
	blocks := strings.Split(input, "\n\n")
	boldItalicPattern := regexp.MustCompile(`[*_]{3}(.+?)[*_]{3}`)
	boldPattern := regexp.MustCompile(`[*_]{2}(.+?)[*_]{2}`)
	italicPattern := regexp.MustCompile(`[*_]{1}(.+?)[*_]{1}`)
	codePattern := regexp.MustCompile("`(.+?)`")
	urlPattern := regexp.MustCompile(`\[(.+?)\]\((.+?)\)`)

	for _, p := range blocks {
		p = boldItalicPattern.ReplaceAllString(p, "<em><strong>$1</strong></em>")
		p = boldPattern.ReplaceAllString(p, "<strong>$1</strong>")
		p = italicPattern.ReplaceAllString(p, "<em>$1<em>")
		p = codePattern.ReplaceAllString(p, "<code>$1</code>")
		p = urlPattern.ReplaceAllString(p, "<a href='$2'>$1</a>")
		p = "<p>" + p + "</p>"
	}

	return strings.Join(blocks, "\n")
}
