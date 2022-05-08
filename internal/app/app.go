package app

import (
	"fmt"
	"go/ast"
	"godapi/internal/docengine"
	"os"
)

func Launch() {
	docs := docengine.NewDocEngine(fullOutputDirPath())

	//set funcs
	setDocFuncs(docs)

	//open each go file in InputDir, parse comments & execute funcs
	if err := openEachFile(fullInputDirPath(), `^.*\.go$`, func(af *ast.File) (err error) {
		err = scanFileComments(af, func(comment string) {
			obj := tokensToObject(tokenize(comment))
			if err := docs.ExecFunc(obj.Method, obj.Args); err != nil {
				fmt.Printf("Execute error: %s", err.Error())
				os.Exit(0)
			}
		})
		if err != nil {
			return err
		}
		return nil
	}); err != nil {
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
