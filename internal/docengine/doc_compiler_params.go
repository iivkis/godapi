package docengine

import (
	"go/ast"
	"go/types"
)

type DocCompilerItemParam struct {
	Name    string
	Located string // in body, query
	Fields  []*DocCompilerItemParamField
}

type DocCompilerItemParamField struct {
	Name string `json:"name"`
	Type string `json:"type"`
	Tags string `json:"-"`
}

func NewDocCompilerItemParam(name string, located string, s *ast.StructType) *DocCompilerItemParam {
	p := &DocCompilerItemParam{
		Name:    name,
		Fields:  make([]*DocCompilerItemParamField, 0, len(s.Fields.List)),
		Located: located,
	}

	for _, field := range s.Fields.List {
		f := &DocCompilerItemParamField{}

		f.Name = field.Names[0].Name
		f.Type = types.ExprString(field.Type)

		if field.Tag != nil {
			f.Tags = field.Tag.Value
		}

		p.Fields = append(p.Fields, f)
	}

	return p
}
