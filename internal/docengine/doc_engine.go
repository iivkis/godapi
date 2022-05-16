package docengine

import (
	"go/ast"
	"os"
)

// key - "[package].[struct_name]"
type DocEngineStructs map[string]*ast.StructType

func (s DocEngineStructs) Add(pkgName string, structName string, val *ast.StructType) {
	s[pkgName+"."+structName] = val
}

func (s DocEngineStructs) Get(pkgName string, structName string) *ast.StructType {
	return s[pkgName+"."+structName]

}

type DocEngine struct {
	Meta    *DocEngineMeta
	Structs DocEngineStructs

	compiler *DocCompiler
	funcs    map[string]DocEngineAddFuncData
	outDir   string
}

func NewDocEngine(outDir string) *DocEngine {
	return &DocEngine{
		Meta:    NewDocEngineMeta(),
		Structs: make(map[string]*ast.StructType),

		compiler: NewDocCompiler(),
		funcs:    make(map[string]DocEngineAddFuncData),
		outDir:   outDir,
	}
}

func (d *DocEngine) Compile() error {
	d.compiler = NewDocCompiler()

	//create groups
	d.compiler.initGroups(d.Meta)

	//distribution of items by groups and subgroups
	if err := d.compiler.initItems(d.Meta, d.Structs); err != nil {
		return err
	}

	//set values to MainInfo
	if err := d.compiler.initMainInfo(d.Meta); err != nil {
		return err
	}

	return nil
}

func (d *DocEngine) ClearDir() error {
	if err := os.RemoveAll(d.outDir); err != nil {
		return err
	}

	if err := os.MkdirAll(d.outDir, os.ModePerm); err != nil {
		return err
	}

	return nil
}

func (d *DocEngine) SaveJSON() error {
	builder := NewDocJSONBuidler(d.compiler)
	return builder.Save(d.outDir)
}

func (d *DocEngine) SaveHTML() error {
	builder := NewHTMLBuilder(d.compiler)
	return builder.Save(d.outDir)
}
