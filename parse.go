package link

import (
	"io"
	"golang.org/x/net/html"
	"strings"
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
	nodes := getLinks(doc)
	var links []Link
	for _, node := range nodes {
		links = append(links, buildLinks(node))
	}
	return links, nil
}

// Gets all links from a doc
func getLinks(n *html.Node) []*html.Node {
	if n.Type == html.ElementNode && n.Data == "a" {
		return []*html.Node{n}
	}
	var ret []*html.Node

	for c:= n.FirstChild; c!= nil; c = c.NextSibling {
		ret = append(ret, getLinks(c)...)
	}

	return ret
}

// Builds a Link (Struct) from passed nodes which are links
func buildLinks(n *html.Node) Link {
	var ret Link
	for _, attr := range n.Attr {
		if attr.Key == "href" {
			ret.Href = attr.Val
			break
		}
	}
	ret.Text = getTextFromLink(n)
	return ret
}

// Extracts the text from a link
func getTextFromLink(n *html.Node) string {
	if n.Type == html.TextNode {
		return n.Data
	}
	if n.Type != html.ElementNode {
		return ""
	}
	var ret string 
	for c := n.FirstChild; c!= nil; c = c.NextSibling {
		ret += getTextFromLink(c)
	}
	return strings.Join(strings.Fields(ret), " ")
}