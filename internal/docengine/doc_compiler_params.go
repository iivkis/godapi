package docengine

import (
	"go/ast"
	"go/types"
	"strings"

	"github.com/fatih/structtag"
)

type DocCompilerItemParam struct {
	Name      string                 `json:"name"`
	Located   string                 `json:"located"`
	FieldsMap map[string]interface{} `json:"fields"`

	fields []*DocCompilerItemParamField `json:"-"`
}

type DocCompilerItemParamField struct {
	Name string
	Tag  string

	Type struct {
		Name string
		Expr ast.Expr
	}

	Node *DocCompilerItemParam `json:"node"`
}

func NewDocCompilerItemParam(name string, pkg string, located string, structs DocEngineStructs) *DocCompilerItemParam {
	s, ok := structs[pkg+"."+name]
	if !ok {
		return nil
	}

	var param = &DocCompilerItemParam{
		Name:    name,
		Located: located,
		fields:  make([]*DocCompilerItemParamField, 0, len(s.Fields.List)),
	}

	for _, field := range s.Fields.List {
		if !field.Names[0].IsExported() {
			continue
		}

		//create new field
		f := &DocCompilerItemParamField{}

		//set Tag
		if field.Tag != nil {
			f.Tag = strings.Trim(field.Tag.Value, "`")
		}

		//set Name
		{
			f.Name = field.Names[0].Name

			//if have json tag
			if f.Tag != "" {
				tag, err := structtag.Parse(f.Tag)
				if err != nil {
					panic(err)
				}

				j, err := tag.Get("json")
				if err != nil {
					panic(err)
				}
				f.Name = j.Name
			}
		}

		//set Type.Name
		{
			switch field.Type.(type) {
			case *ast.Ident:
				f.Type.Name = field.Type.(*ast.Ident).Name
			case *ast.StarExpr:
				f.Type.Name = types.ExprString(field.Type.(*ast.StarExpr).X)
			case *ast.ArrayType:
				f.Type.Name = types.ExprString(field.Type.(*ast.ArrayType).Elt)
			}
		}

		//set Type.Expr
		f.Type.Expr = field.Type

		//set node if exists
		f.Node = NewDocCompilerItemParam(f.Type.Name, pkg, located, structs)

		//push field
		param.fields = append(param.fields, f)
	}

	param.FieldsMap = param.toMap()
	return param
}

func (d *DocCompilerItemParam) toMap() map[string]interface{} {
	m := make(map[string]interface{})
	for _, field := range d.fields {
		if field.Node == nil {
			switch field.Type.Expr.(type) {
			case *ast.Ident: //simple type
				m[field.Name] = field.Type.Name
			case *ast.StarExpr: //type with pointer
				m[field.Name] = field.Type.Name
			default: //arrays, maps and other types
				m[field.Name] = types.ExprString(field.Type.Expr)
			}
		} else {
			switch field.Type.Expr.(type) {
			case *ast.Ident: //simple type
				m[field.Name] = field.Node.toMap()
			case *ast.StarExpr: //type with pointer
				m[field.Name] = field.Node.toMap()
			default: //arrays, maps and other types
				m[field.Name] = []interface{}{field.Node.toMap()}
			}
		}
	}
	return m
}
