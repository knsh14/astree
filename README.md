astree
---

tree command for Go AST

# install
```
go get -u github.com/knsh14/astree/cmd/astree
```

# options

# Example
https://play.golang.org/p/ucQjFbdeNcQ

```
package main

import (
	"go/parser"
	"go/token"
	"log"
	"os"

	"github.com/knsh14/astree"
)

func main() {
	code := `
package main

import (
	"fmt"
)

func main() {
	fmt.Println("Hello, playground")
}`
	fs := token.NewFileSet()
	f, err := parser.ParseFile(fs, "", code, 0)
	if err != nil {
		log.Fatal(err)
	}
	astree.File(os.Stdout, f)
}
```
