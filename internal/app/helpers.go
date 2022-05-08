package app

import (
	"fmt"
	"go/ast"
	"go/parser"
	"go/token"
	"io/fs"
	"os"
	"path"
	"path/filepath"
	"regexp"
)

func fullInputDirPath() string {
	p, err := os.Getwd()
	if err != nil {
		panic(err)
	}

	fullpath := path.Join(p, appFlags.InputDir)
	if _, err := os.Stat(fullpath); os.IsNotExist(err) {
		fmt.Printf("Error: incorrect path `%s`\n", fullpath)
		os.Exit(0)
	}

	return fullpath
}

func fullOutputDirPath() string {
	p, _ := os.Getwd()
	return path.Join(p, appFlags.OutputDir)
}

func openEachFile(inDir string, pattern string, fn func(*ast.File) error) error {
	re, err := regexp.Compile(pattern)
	if err != nil {
		return err
	}

	return filepath.WalkDir(inDir, func(path string, d fs.DirEntry, err error) error {
		fs := token.NewFileSet()
		if !d.IsDir() && re.MatchString(d.Name()) {
			af, err := parser.ParseFile(fs, path, nil, parser.ParseComments)
			if err != nil {
				return err
			}

			if fn(af) != nil {
				return err
			}
		}
		return err
	})
}

func scanFileComments(af *ast.File, fn func(comment string)) error {
	for _, commentGroup := range af.Comments {
		for _, comment := range commentGroup.List {
			line := []byte(comment.Text)
			if len(line) > 3 && string(line[:3]) == "//@" {
				fn(string(line[3:]))
			}
		}
	}
	return nil
}
