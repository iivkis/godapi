package docengine

import "encoding/json"

type DocEngineMeta struct {
	AppName      string
	AppVersion   string
	DefaultGroup string

	Groups    []*MetaGroup
	Subgroups []*MetaSubgroup
	Items     []*MetaItem
}

type MetaGroup struct {
	Name        string
	Description string
	Hidden      bool
}

type MetaSubgroup struct {
	Name string
}

type MetaItem struct {
	Annotation  string
	Description []string

	Group    string
	Subgroup string

	Method string
	Route  string

	Params []*MetaItemParam
}

type MetaItemParam struct {
	Located    string //body, query
	StructName string
}

func NewDocEngineMeta() *DocEngineMeta {
	docs := &DocEngineMeta{
		Groups:    make([]*MetaGroup, 0),
		Subgroups: make([]*MetaSubgroup, 0),
		Items:     make([]*MetaItem, 0),
	}

	docs.Groups = append(docs.Groups, &MetaGroup{
		Name:        "main",
		Description: "Main group",
	})

	return docs
}

func NewMetaItem() *MetaItem {
	return &MetaItem{
		Description: make([]string, 0),
		Params:      make([]*MetaItemParam, 0),
	}
}

func (m *MetaItem) ToString() string {
	s, err := json.MarshalIndent(m, "", "  ")
	if err != nil {
		panic(err)
	}
	return string(s)
}
