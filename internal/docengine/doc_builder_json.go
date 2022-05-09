package docengine

import (
	"encoding/json"
	"fmt"
	"os"
	"path"
)

type DocJSONBuilder struct {
	compiler *DocCompiler
}

func NewDocJSONBuidler(compiler *DocCompiler) *DocJSONBuilder {
	return &DocJSONBuilder{compiler: compiler}
}

func (d *DocJSONBuilder) saveMainInfo(outDir string) error {
	js, err := json.MarshalIndent(d.compiler.MainInfo, "", "  ")
	if err != nil {
		return err
	}

	file, err := os.Create(path.Join(outDir, "_info.json"))
	if err != nil {
		return err
	}
	defer file.Close()

	if _, err := file.Write(js); err != nil {
		return err
	}

	fmt.Println("Successfully saved", file.Name())
	return nil
}

func (d *DocJSONBuilder) saveGroupsInfo(outDir string) error {
	for _, group := range d.compiler.Groups {
		js, err := json.MarshalIndent(group, "", "  ")
		if err != nil {
			return err
		}

		file, err := os.Create(path.Join(outDir, fmt.Sprintf("./%s.json", group.Name)))
		if err != nil {
			return err
		}
		defer file.Close()

		if _, err := file.Write(js); err != nil {
			return err
		}

		fmt.Println("Successfully saved", file.Name())
	}
	return nil
}

func (d *DocJSONBuilder) Save(outDir string) error {
	if err := d.saveMainInfo(outDir); err != nil {
		return err
	}

	if err := d.saveGroupsInfo(outDir); err != nil {
		return err
	}

	return nil
}
