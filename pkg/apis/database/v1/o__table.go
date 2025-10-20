package v1

import metav1 "github.com/octohelm/objectkind/pkg/apis/meta/v1"

type TableCode string

// +gengo:objectkind
type Table struct {
	metav1.TypeMeta
	metav1.Codable[TableCode]

	Spec     TableSpec `json:"spec,omitzero"`
	Database *Database `json:"database,omitzero"`
}

type TableSpec struct {
	Columns     []*Column     `json:"columns,omitzero"`
	Constraints []*Constraint `json:"constraints,omitzero"`
}
