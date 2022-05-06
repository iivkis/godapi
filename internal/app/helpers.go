package app

import (
	"bufio"
	"fmt"
	"io/fs"
	"os"
	"path"
	"path/filepath"
	"regexp"
	"strings"
)

func fullInputDirPath() string {
	p, err := os.Getwd()
	if err != nil {
		panic(err)
	}

	fullpath := path.Join(p, appFlags.InputDir)

	if _, err := os.Stat(fullpath); os.IsNotExist(err) {
		fmt.Printf("Incorrect path `%s`\n", fullpath)
		os.Exit(0)
	}

	return fullpath
}

func fullOutputDirPath() string {
	p, _ := os.Getwd()
	return path.Join(p, appFlags.OutputDir)
}

func openEachFile(inDir string, pattern string, fn func(*os.File) error) error {
	re, err := regexp.Compile(pattern)
	if err != nil {
		return err
	}

	return filepath.WalkDir(inDir, func(path string, d fs.DirEntry, err error) error {
		if !d.IsDir() && re.MatchString(d.Name()) {
			file, err := os.OpenFile(path, os.O_RDONLY, os.ModePerm)
			if err != nil {
				panic(err)
			}
			defer file.Close()

			if fn(file) != nil {
				return err
			}
		}
		return err
	})
}

func scanFileComments(file *os.File, fn func(comment string)) error {
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		if err := scanner.Err(); err != nil {
			return err
		}

		line := []byte(strings.TrimSpace(scanner.Text()))
		if len(line) > 3 && string(line[:3]) == "//@" {
			fn(string(line[3:]))
		}
	}
	return nil
}
