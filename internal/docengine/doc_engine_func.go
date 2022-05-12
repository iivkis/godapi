package docengine

import (
	"fmt"
	"strings"
)

type DocEngineFunc func(meta *DocEngineMeta, args []string) error

type DocEngineAddFuncData struct {
	ArgsCount uint8
	Func      DocEngineFunc
}

type DocEngineExecFuncArgs struct {
	Method  string
	Package string
	Args    []string
}

func (d *DocEngine) AddFunc(fnName string, argsCount uint8, fn DocEngineFunc) {
	d.funcs[fnName] = DocEngineAddFuncData{
		ArgsCount: argsCount,
		Func:      fn,
	}
}

func (d *DocEngine) ExecFunc(a *DocEngineExecFuncArgs) error {
	if funcData, exists := d.funcs[a.Method]; exists {
		if funcData.ArgsCount == uint8(len(a.Args)) {
			a.Args = append(a.Args, a.Package)
			return funcData.Func(d.Meta, a.Args)
		} else {
			fmt.Printf("Warning: invalid number of input arguments: have %d, need %d.\n", len(a.Args), funcData.ArgsCount)
			fmt.Printf("func: %s; args: [%s].\n", a.Method, strings.Join(a.Args, ", "))
		}
	} else {
		fmt.Printf("Warning: undefined func `@%s`\n", a.Method)
	}
	return nil
}
