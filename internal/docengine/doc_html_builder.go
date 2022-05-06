package docengine

import (
	"html/template"
	"os"
	"path"
)

type DocHTMLBuilder struct {
	compiler   *DocCompiler
	templGroup *template.Template
}

func NewHTMLBuilder(builder *DocCompiler) *DocHTMLBuilder {
	w := &DocHTMLBuilder{compiler: builder}
	w.loadHTMLTempl()
	return w
}

func (w *DocHTMLBuilder) loadHTMLTempl() {
	//load group
	tGroup, err := template.ParseFiles("./web/group.html")
	if err != nil {
		panic(err)
	}
	w.templGroup = tGroup
}

func (w *DocHTMLBuilder) renderGroups(outDir string) error {
	if err := os.Mkdir(path.Join(outDir, "./html"), os.ModePerm); err != nil {
		return err
	}

	for _, group := range w.compiler.Groups {
		f, err := os.Create(path.Join(outDir, "./html/", group.Name+".html"))
		if err != nil {
			return err
		}
		defer f.Close()

		if err := w.templGroup.Execute(f, map[string]interface{}{
			"current": group,
			"main":    w.compiler.MainInfo,
		}); err != nil {
			return err
		}
	}

	return nil
}

func (w *DocHTMLBuilder) Save(outDir string) error {
	if err := w.renderGroups(outDir); err != nil {
		return err
	}
	return nil
}
