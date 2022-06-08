package main

import (
	"fmt"
	"log"

	"github.com/influxdata/flux/ast"
	"github.com/influxdata/flux/ast/astutil"
)

func main() {
	file := &ast.File{
		Name: "myfile",
		Body: []ast.Statement{
			&ast.VariableAssignment{
				ID: &ast.Identifier{
					Name: "foo",
				},
				Init: ast.IntegerLiteralFromValue(4),
			},
		},
	}

	s, err := astutil.Format(file)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(s)
}
