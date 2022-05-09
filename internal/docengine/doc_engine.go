package docengine

import (
	"go/ast"
	"os"
)

type DocEngine struct {
	Meta *DocEngineMeta

	compiler *DocCompiler
	funcs    map[string]DocEngineAddFuncData

	outDir string

	Structs map[string]*ast.StructType
}

func NewDocEngine(outDir string) *DocEngine {
	return &DocEngine{
		Meta:     NewDocEngineMeta(),
		compiler: NewDocCompiler(),
		funcs:    make(map[string]DocEngineAddFuncData),
		Structs:  make(map[string]*ast.StructType),
		outDir:   outDir,
	}
}

func (d *DocEngine) Compile() error {
	d.compiler = NewDocCompiler()

	//create groups
	d.compiler.initGroups(d.Meta.Groups)

	//distribution of items by groups and subgroups
	d.compiler.initItems(d.Meta.Items, d.Structs)

	//set values to MainInfo
	d.compiler.initMainInfo(d.Meta)

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
