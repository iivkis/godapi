package app

import (
	"fmt"
	"os"
	"strings"

	"github.com/iivkis/godapi/internal/docengine"
)

var ALLOWED_PARAM_LOCATION = map[string]byte{
	"body":  0,
	"query": 0,
}

/*GLOBAL FUNCS*/
func setDocGlobalFuncs(doc *docengine.DocEngine) {
	//arg_1 - name for your app
	doc.AddFunc("@AppName", 1, func(meta *docengine.DocEngineMeta, args []string) error {
		meta.AppName = args[0]
		return nil
	})

	//arg_1 - documentation version
	doc.AddFunc("@AppVersion", 1, func(meta *docengine.DocEngineMeta, args []string) error {
		meta.AppVersion = args[0]
		return nil
	})

	//arg_1 - group name
	//arg_2 - group descroption
	doc.AddFunc("@AddGroup", 2, func(meta *docengine.DocEngineMeta, args []string) error {
		meta.Groups = append(meta.Groups, &docengine.MetaGroup{
			Name:        args[0],
			Description: args[1],
		})
		return nil
	})

	//arg_1 - group name
	doc.AddFunc("@HideGroup", 1, func(meta *docengine.DocEngineMeta, args []string) error {
		for _, group := range meta.Groups {
			if group.Name == args[0] {
				group.Hidden = true
				return nil
			}
		}
		fmt.Printf("@HideGroup: incorrect group name `%s`", args[0])
		os.Exit(0)

		return nil
	})
}

/*FUNC FOR CURRENT ITEM*/
func setDocItemFuncs(doc *docengine.DocEngine) {
	var item *docengine.MetaItem

	//arg_1 - method annatation
	doc.AddFunc("New", 1, func(meta *docengine.DocEngineMeta, args []string) error {
		item = docengine.NewMetaItem()
		item.Annotation = args[0]
		return nil
	})

	//arg_1 - description for method or param
	doc.AddFunc("Desc", 1, func(meta *docengine.DocEngineMeta, args []string) error {
		if len(item.Params) == 0 {
			item.Description = append(item.Description, args[0])
		} else {
			item.Params[len(item.Params)-1].Description = append(item.Params[len(item.Params)-1].Description, args[0])
		}
		return nil
	})

	//arg_1 - group ("main" if group empty)
	//arg_2 - subgroup ("default" if subgroup empty)
	doc.AddFunc("Group", 2, func(meta *docengine.DocEngineMeta, args []string) error {
		item.Group = args[0]
		item.Subgroup = args[1]
		return nil
	})

	//arg_1 - query/body
	//arg_2 - struct name
	doc.AddFunc("Param", 2, func(meta *docengine.DocEngineMeta, args []string) error {
		if _, ok := ALLOWED_PARAM_LOCATION[args[0]]; ok {
			item.Params = append(item.Params, &docengine.MetaItemParam{
				Located:        args[0],
				StructName:     args[1],
				CurrentPackage: args[2],
			})
		}
		return nil
	})

	//arg_1 - request method (GET, POST, PUT, DELETE, etc.)
	//arg_2 - route (example: /users)
	doc.AddFunc("Route", 2, func(meta *docengine.DocEngineMeta, args []string) error {
		item.Method = strings.ToUpper(args[0])
		item.Route = args[1]

		//push item
		meta.Items = append(meta.Items, item)
		return nil
	})
}
