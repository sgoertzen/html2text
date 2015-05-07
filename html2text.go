package html2text

import (
	"bytes"
	"errors"
	"fmt"
	"golang.org/x/net/html"
	"strings"
)

var breakers = map[string]bool{
	"br":  true,
	"div": true,
	"tr":  true,
	"li":  true,
}

func Textify(body string) (string, error) {
	r := strings.NewReader(body)
	doc, err := html.Parse(r)
	if err != nil {
		return "", errors.New("Unable to parse the html")
	}
	var buffer bytes.Buffer
	process(doc, &buffer)

	return strings.TrimSpace(buffer.String()), nil
}

func process(n *html.Node, b *bytes.Buffer) {
	processChildren := true

	if n.Type == html.ElementNode && n.Data == "head" {
		return
	} else if n.Type == html.ElementNode && n.Data == "a" && n.FirstChild != nil {
		processChildren = anchor(n, b)
	} else if n.Type == html.TextNode {
		b.WriteString(cleanup(n.Data))
	}
	if processChildren {
		for c := n.FirstChild; c != nil; c = c.NextSibling {
			process(c, b)
		}
	}
	if b.Len() > 0 {
		last := b.Bytes()[b.Len()-1]
		if last != '\n' && n.Type == html.ElementNode && breakers[n.Data] {
			b.WriteString("\n")
		} else if last != ' ' && last != '\n' {
			b.WriteString(" ")
		}
	}
}

func cleanup(text string) string {
	return strings.TrimSpace(strings.Replace(text, "\u00a0", " ", -1))
}

func anchor(n *html.Node, b *bytes.Buffer) bool {
	start := b.Len()
	for c := n.FirstChild; c != nil; c = c.NextSibling {
		process(c, b)
	}
	// check to see if the link text has been written out already
	bytes := b.Bytes()
	end := b.Len()
	link := n.Attr[0].Val
	recentlyWrittenBytes := bytes[start:end]
	if !strings.Contains(string(recentlyWrittenBytes), link) {
		b.WriteString(fmt.Sprintf("(link: %s)", link))
	}
	return false
}
