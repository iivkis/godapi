package docengine

import (
	"encoding/json"
)

type MetaItem struct {
	Annotation  string
	Description []string

	Group    string
	Subgroup string

	Method string
	Route  string

	Params *MetaItemParams
}

type MetaItemParams struct {
	BodyStructName  string
	QueryStructName string
}

func NewMetaItem() *MetaItem {
	return &MetaItem{
		Description: make([]string, 0),
		Params:      &MetaItemParams{},
	}
}

func (m *MetaItem) ToString() string {
	s, err := json.MarshalIndent(m, "", "  ")
	if err != nil {
		panic(err)
	}
	return string(s)
}
