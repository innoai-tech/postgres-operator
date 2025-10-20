package v1

import metav1 "github.com/octohelm/objectkind/pkg/apis/meta/v1"

// +gengo:objectkind
type Result struct {
	metav1.TypeMeta

	Columns []*ResultColumn `json:"columns,omitzero"`
	Data    [][]any         `json:"data"`
}

type ResultColumn struct {
	Code ColumnCode `json:"code"`
	Type string     `json:"type"`
}
