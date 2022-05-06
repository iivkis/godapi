package docengine

import (
	"fmt"
)

type DocCompiler struct {
	MainInfo *DocCompilerMainInfo
	Groups   map[string]*DocCompilerGroup
}

//File with base info
type DocCompilerMainInfo struct {
	AppName    string `json:"app_name"`
	AppVersion string `json:"app_version"`

	Groups       []string `json:"groups"`
	DefaultGroup string   `json:"default_group"`
}

//File for each group
type DocCompilerGroup struct {
	Name        string               `json:"name"`
	Description string               `json:"description"`
	Subgroups   DocCompilerSubgroups `json:"subgroups"`

	hidden bool
}

// key 1 - subgroup name; key 2 - item route
type DocCompilerSubgroups map[string][]*DocCompilerSubgroupItem

//Item data
type DocCompilerSubgroupItem struct {
	Annotation  string   `json:"annotation"`
	Description []string `json:"description"`
	Method      string   `json:"method"`
	Route       string   `json:"route"`
}

func NewDocCompiler() *DocCompiler {
	return &DocCompiler{
		MainInfo: &DocCompilerMainInfo{
			Groups: make([]string, 0),
		},
		Groups: make(map[string]*DocCompilerGroup),
	}
}

func (b *DocCompiler) initGroups(g []*MetaGroup) {
	for _, group := range g {
		b.Groups[group.Name] = &DocCompilerGroup{
			Name:        group.Name,
			Description: group.Description,
			Subgroups:   make(DocCompilerSubgroups),

			hidden: group.Hidden,
		}
	}
}

func (b *DocCompiler) initItems(items []*MetaItem) {
	for _, item := range items {
		//set default group
		if item.Group == "" {
			item.Group = "main"
		}

		//set default subgroup
		if item.Subgroup == "" {
			item.Subgroup = "default"
		}

		//check group exists
		if g := b.Groups[item.Group]; g != nil {
			if g.hidden {
				continue
			}
		} else {
			fmt.Printf("@Group: undefined group `%s`\n", item.Group)
			fmt.Println(item.ToString())
			return
		}

		//create subgroup if not exists
		// if b.Groups[item.Group].Subgroups[item.Subgroup] == nil {
		// 	b.Groups[item.Group].Subgroups[item.Subgroup] = make(map[string]*DocCompilerSubgroupItem)
		// }

		//add item to [group][subgroup]
		b.Groups[item.Group].Subgroups[item.Subgroup] = append(b.Groups[item.Group].Subgroups[item.Subgroup], &DocCompilerSubgroupItem{
			Annotation:  item.Annotation,
			Description: item.Description,
			Method:      item.Method,
			Route:       item.Route,
		})
	}
}

func (b *DocCompiler) initMainInfo(meta *DocEngineMeta) {
	//AppName
	b.MainInfo.AppName = meta.AppName

	//AppVersion
	b.MainInfo.AppVersion = meta.AppVersion

	//Clean & append groups to main
	for key, group := range b.Groups {
		if len(group.Subgroups) == 0 {
			delete(b.Groups, key)
			continue
		}

		b.MainInfo.Groups = append(b.MainInfo.Groups, key)
		fmt.Printf("Add group `%s`\n", key)
	}

	//DefaultGroup. If empty then "main"
	if meta.DefaultGroup == "" {
		b.MainInfo.DefaultGroup = "main"
	} else {
		if _, exists := b.Groups[meta.DefaultGroup]; exists {
			b.MainInfo.DefaultGroup = meta.DefaultGroup
		} else {
			fmt.Printf("@DefaultGroup: undefined group `%s`\n", meta.DefaultGroup)
		}
	}
}
