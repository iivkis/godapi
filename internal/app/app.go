package app

import (
	"flag"
	"fmt"
	"go/ast"
	"godapi/internal/docengine"
	"os"
)

var appFlags struct {
	InputDir  string
	OutputDir string

	GenHTML bool //generate html static files
}

func init() {
	flag.StringVar(&appFlags.InputDir, "in", "./", "folder with files for which documentation is needed")
	flag.StringVar(&appFlags.OutputDir, "out", "./gdocs", "folder with files for which documentation is needed")
	flag.BoolVar(&appFlags.GenHTML, "html", false, "generate html static files")

	flag.Parse()
}

func Launch() {
	docs := docengine.NewDocEngine(fullOutputDirPath())

	//set funcs
	setDocFuncs(docs)

	//open each go file in InputDir, parse comments & execute funcs
	if err := pasrseGoFilesToDocs(docs); err != nil {
		panic(err)
	}

	//compile docs
	if err := docs.Compile(); err != nil {
		panic(err)
	}

	//clear dir
	if err := docs.ClearDir(); err != nil {
		panic(err)
	}

	//save docs to json files
	if err := docs.SaveJSON(); err != nil {
		panic(err)
	}

	//save docs to html files
	if appFlags.GenHTML {
		if err := docs.SaveHTML(); err != nil {
			panic(err)
		}
	}
}

func setDocFuncs(docs *docengine.DocEngine) {
	setDocGlobalFuncs(docs)
	setDocItemFuncs(docs)
}

func pasrseGoFilesToDocs(docs *docengine.DocEngine) error {
	return openEachFile(fullInputDirPath(), `^.*\.go$`, func(af *ast.File) error {
		visitor := docengine.NewDocVisitor() //get all file structs
		ast.Walk(visitor, af)

		//execute each comment
		if err := scanFileComments(af, func(comment string) {
			obj := tokensToObject(tokenize(comment))
			if err := docs.ExecFunc(&docengine.DocEngineExecFuncArgs{
				Method:  obj.Method,
				Args:    obj.Args,
				Package: visitor.CurrentPackageName,
			}); err != nil {
				fmt.Printf("Execute error: %s", err.Error())
				os.Exit(0)
			}
		}); err != nil {
			return err
		}

		//set found structs
		for key, val := range visitor.Structs {
			docs.Structs[key] = val
		}
		return nil
	})
}
