package docengine

import (
	"bytes"
	"html/template"
	"io"
	"io/fs"
	"os"
	"path"
	"path/filepath"

	"github.com/tdewolff/minify/v2"
	"github.com/tdewolff/minify/v2/css"
	"github.com/tdewolff/minify/v2/html"
)

type DocHTMLBuilder struct {
	compiler  *DocCompiler
	templates *template.Template
	minify    *minify.M
}

func NewHTMLBuilder(builder *DocCompiler) *DocHTMLBuilder {
	w := &DocHTMLBuilder{compiler: builder}
	w.loadHTMLTempl()
	w.initMinifier()
	return w
}

func (w *DocHTMLBuilder) loadHTMLTempl() {
	t, err := template.ParseGlob("./web/*.html")
	if err != nil {
		panic(err)
	}
	w.templates = t
}

func (w *DocHTMLBuilder) initMinifier() {
	w.minify = minify.New()

	w.minify.Add("text/html", &html.Minifier{
		KeepEndTags:             true,
		KeepComments:            true,
		KeepConditionalComments: true,
		KeepDefaultAttrVals:     true,
		KeepDocumentTags:        true,
		KeepQuotes:              true,
		KeepWhitespace:          true,
	})

	w.minify.Add("text/css", &css.Minifier{
		KeepCSS2:  true,
		Precision: 1,
	})
}

func (w *DocHTMLBuilder) makeOutDir(outDir string) error {
	return os.MkdirAll(path.Join(outDir, "./html/src"), os.ModePerm)
}

func (w *DocHTMLBuilder) renderGroups(outDir string) error {
	for _, group := range w.compiler.Groups {
		var buf bytes.Buffer

		f, err := os.Create(path.Join(outDir, "./html/", group.Name+".html"))
		if err != nil {
			return err
		}
		defer f.Close()

		if err := w.templates.ExecuteTemplate(&buf, "group.html", map[string]interface{}{
			"current_group": group,
			"main":          w.compiler.MainInfo,
		}); err != nil {
			return err
		}

		if err := w.minify.Minify("text/html", f, &buf); err != nil {
			return err
		}
	}
	return nil
}

func (w *DocHTMLBuilder) copySrc(outDir string) error {
	return filepath.Walk("./web/src", func(fpath string, info fs.FileInfo, err error) error {
		if info.IsDir() {
			return err
		}

		fIn, err := os.Open(fpath)
		if err != nil {
			return err
		}
		defer fIn.Close()

		fOut, err := os.Create(path.Join(outDir, "./html/src", info.Name()))
		if err != nil {
			return err
		}
		defer fOut.Close()

		switch filepath.Ext(info.Name()) {
		case ".css":
			if err := w.minify.Minify("text/css", fOut, fIn); err != nil {
				return err
			}
		default:
			if _, err := io.Copy(fOut, fIn); err != nil {
				return err
			}
		}

		return err
	})
}

func (w *DocHTMLBuilder) Save(outDir string) error {
	if err := w.makeOutDir(outDir); err != nil {
		return err
	}

	if err := w.copySrc(outDir); err != nil {
		return err
	}

	if err := w.renderGroups(outDir); err != nil {
		return err
	}
	return nil
}
