package link

import (
	"io"
	"fmt"
	"golang.org/x/net/html"
)

// Link represents a link in an HTML document aka <a href=""... />
type Link struct {
	Href string
	Text string
}


// Parse takes an HTML document and returns a slice of links parsed from it
func Parse (r io.Reader) ([]Link, error) {
	doc, err := html.Parse(r)
	if err != nil {
		return nil, err
	}
	var f func(*html.Node)
	f = func(n *html.Node) {
		if n.Type == html.ElementNode && n.Data == "a" {
			fmt.Println(n)
		}
		for c := n.FirstChild; c != nil; c = c.NextSibling {
			f(c)
		}
	}
	f(doc)
	return nil, nil
}