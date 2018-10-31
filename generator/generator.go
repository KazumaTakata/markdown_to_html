package generator

import "markdown_to_html/ast"

func Generator(node interface{}) string {
	switch node := node.(type) {
	case *ast.Program:
		return evalStatements(node.Statements)
	}
}

func evalStatements(node interface{}) string {

}
