package docengine

import (
	"bytes"
	"fmt"
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
	// s, ok := structs[pkg+"."+name]
	s := structs.Get(pkg, name)
	if s == nil {
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
		f.Name = field.Names[0].Name

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

	param.FieldsMap = param.ToMap("json")
	return param
}

func (d *DocCompilerItemParam) ToMap(tagKey string) map[string]interface{} {
	m := make(map[string]interface{})
	for _, field := range d.fields {
		//if have key tag
		if tagKey != "" && field.Tag != "" {
			tag, err := structtag.Parse(field.Tag)
			if err != nil {
				panic(err)
			}

			s, err := tag.Get(tagKey)
			if err != nil {
				panic(err)
			}
			field.Name = s.Name
		}

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
				m[field.Name] = field.Node.ToMap(tagKey)
			case *ast.StarExpr: //type with pointer
				m[field.Name] = field.Node.ToMap(tagKey)
			default: //arrays, maps and other types
				m[field.Name] = []interface{}{field.Node.ToMap(tagKey)}
			}
		}
	}
	return m
}

func (d *DocCompilerItemParam) Text() string {
	m := d.marshal(d.FieldsMap, 1)
	return string(m.String())
}

func (d *DocCompilerItemParam) marshal(m map[string]interface{}, deep int) (b bytes.Buffer) {
	b.WriteString("{")
	for name, data := range m {
		b.WriteString(fmt.Sprintf("\n%s%s: ", strings.Repeat("    ", deep), name))

		//print type
		switch data := data.(type) {
		case string:
			b.WriteString(data)
		case map[string]interface{}:
			m := d.marshal(data, deep+1)
			b.WriteString(m.String())
		case []interface{}:
			m := d.marshal(data[0].(map[string]interface{}), deep+1)
			b.WriteString(fmt.Sprintf("[%s]", m.String()))
		}
	}
	b.WriteString(fmt.Sprintf("\n%s}", strings.Repeat("    ", deep-1)))
	return b
}
