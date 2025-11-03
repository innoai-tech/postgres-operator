package v1

import (
	metav1 "github.com/octohelm/objectkind/pkg/apis/meta/v1"
)

type ColumnCode string

// +gengo:objectkind
type Column struct {
	metav1.TypeMeta
	metav1.Codable[ColumnCode]

	Spec ColumnSpec `json:"spec"`
}

type ColumnSpec struct {
	Type string `json:"type"`
}
