package docengine

type DocEngineMeta struct {
	AppName      string
	AppVersion   string
	DefaultGroup string

	Groups    []*MetaGroup
	Subgroups []*MetaSubgroup
	Items     []*MetaItem
}

func NewDocEngineMeta() *DocEngineMeta {
	doc := &DocEngineMeta{
		Groups:    make([]*MetaGroup, 0),
		Subgroups: make([]*MetaSubgroup, 0),
		Items:     make([]*MetaItem, 0),
	}

	doc.Groups = append(doc.Groups, &MetaGroup{
		Name:        "main",
		Description: "Main group",
	})

	return doc
}
