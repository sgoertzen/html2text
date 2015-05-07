package html2text

import (
	"bytes"
	"fmt"
	"golang.org/x/net/html"
	"strings"
)

func Textify(body string) string {
	r := strings.NewReader(body)
	doc, err := html.Parse(r)
	if err != nil {
		// ...
	}
	
	var breakers = make(map[string]bool)
	breakers["br"] = true
	breakers["div"] = true
	breakers["tr"] = true
	breakers["li"] = true
	
	var f func(*html.Node, *bytes.Buffer)
	f = func(n *html.Node, b *bytes.Buffer) {
		processChildren := true
		
		if n.Type == html.ElementNode && n.Data == "head" {
			return
		} else if n.Type == html.ElementNode && n.Data == "a" && n.FirstChild != nil {
			for c := n.FirstChild; c != nil; c = c.NextSibling {
				f(c, b)
			}
			b.WriteString(fmt.Sprintf(" (link: %s)", n.Attr[0].Val))
			processChildren = false
		} else if n.Type == html.TextNode {
			b.WriteString(n.Data)
		} 
		if processChildren {
			for c := n.FirstChild; c != nil; c = c.NextSibling {
				f(c, b)
			}
		}
		if n.Type == html.ElementNode && breakers[n.Data] {
			b.WriteString("\n")
		}
	}
	var buffer bytes.Buffer
	f(doc, &buffer)
	return strings.TrimSpace(strings.Replace(buffer.String(), "\u00a0", " ", -1))
}