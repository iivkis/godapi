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
	if val, ok := n.(*ast.File); ok {
		v.CurrentPackageName = val.Name.Name
	} else if val, ok := n.(*ast.StructType); ok {
		v.Structs.Add(v.CurrentPackageName, v.lastIdent.Name, val)
	} else if val, ok := n.(*ast.Ident); ok {
		v.lastIdent = val
	} else if val != nil {
		v.lastIdent = &ast.Ident{}
	}
	return v
}
