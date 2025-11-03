package v1

import (
	metav1 "github.com/octohelm/objectkind/pkg/apis/meta/v1"
)

type ConstraintCode string

// +gengo:objectkind
type Constraint struct {
	metav1.TypeMeta
	metav1.Codable[ConstraintCode]

	Spec ConstraintSpec `json:"spec"`
}

type ConstraintSpec struct {
	Columns []*ConstraintColumn `json:"columns"`
	Method  string              `json:"method,omitzero"`
	Unique  bool                `json:"unique,omitzero"`
	Primary bool                `json:"primary,omitzero"`
}

type ConstraintColumn struct {
	metav1.Codable[ColumnCode]

	Options []string `json:"options,omitzero"`
}
