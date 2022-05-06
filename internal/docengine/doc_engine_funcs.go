package docengine

import (
	"fmt"
	"strings"
)

type DocEngineFunc func(meta *DocEngineMeta, args []string) error

type DocEngineFuncData struct {
	ArgsCount uint8
	Func      DocEngineFunc
}

func (d *DocEngine) AddFunc(fnName string, argsCount uint8, fn DocEngineFunc) {
	d.funcs[fnName] = DocEngineFuncData{
		ArgsCount: argsCount,
		Func:      fn,
	}
}

func (d *DocEngine) ExecFunc(fnName string, args []string) error {
	if funcData, exists := d.funcs[fnName]; exists {
		if funcData.ArgsCount == uint8(len(args)) {
			return funcData.Func(d.Meta, args)
		} else {
			fmt.Printf("!Warning: invalid number of input arguments: have %d, need %d.\n", len(args), funcData.ArgsCount)
			fmt.Printf("func: %s; args: [%s].\n", fnName, strings.Join(args, ", "))
		}
	} else {
		fmt.Printf("!Warning: undefined func `@%s`\n", fnName)
	}
	return nil
}
