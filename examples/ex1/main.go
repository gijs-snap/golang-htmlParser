package main 

import (
	"fmt"
	"strings"
	"github.com/gijs-snap/golang-htmlParser"
)

var exampleHtml = `
<html>
<body>
  <h1>Hello!</h1>
  <a href="/other-page">A link to another page</a>
  <a href="/other-other-page">A link to another other page</a>
  <div>
	<a href="/other-div-link">A link in a div</a>
  </div>
</body>
</html>
`

func main() {
	r := strings.NewReader(exampleHtml)
	links, err := link.Parse(r)

	if err != nil {
		panic(err)
	}
	fmt.Printf("%+v\n", links)
}