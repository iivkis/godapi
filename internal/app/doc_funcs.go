package app

import (
	"fmt"
	"godapi/internal/docengine"
	"os"
	"strings"
)

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

	//arg_1 - description for method
	doc.AddFunc("Desc", 1, func(meta *docengine.DocEngineMeta, args []string) error {
		item.Description = append(item.Description, args[0])
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
		switch args[0] {
		case "body":
			item.Params.BodyStructName = args[len(args)-1] + "." + args[1]
		case "quert":
			item.Params.QueryStructName = args[len(args)-1] + "." + args[1]
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
