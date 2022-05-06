package app

import "flag"

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
