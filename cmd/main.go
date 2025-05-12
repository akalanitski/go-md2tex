package main

import (
	"fmt"
	"go-md2tex/pkg/markdown"
)

func main() {
	text := `
A Go string is a *read-only* slice & of _ bytes. The language
and the standard library treat strings specially - as
containers of text encoded in [UTF-8](https://en.wikipedia.org/wiki/UTF-8).

In other languages, strings are made of "characters".
In Go, the concept of a character is called a rune - it's
an integer that represents a Unicode code point.
[This Go blog post](https://go.dev/blog/strings) is a good
introduction to the topic.`

	fmt.Printf("TOKENS:\n" + markdown.MarkdownToTex(text))

}
