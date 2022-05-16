package docengine

import "go/ast"

type DocVisitor struct {
	CurrentPackageName string
	Structs            DocEngineStructs
	lastIdent          *ast.Ident
}

func NewDocVisitor() *DocVisitor {
	return &DocVisitor{
		Structs: make(DocEngineStructs),
	}
}

func (v *DocVisitor) Visit(n ast.Node) ast.Visitor {
	switch val := n.(type) {
	case *ast.File:
		v.CurrentPackageName = val.Name.Name
	case *ast.StructType:
		v.Structs.Add(v.CurrentPackageName, v.lastIdent.Name, val)
	case *ast.Ident:
		v.lastIdent = val
	default:
		if val != nil {
			v.lastIdent = &ast.Ident{}
		}
	}

	return v
}
