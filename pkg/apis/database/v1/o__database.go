package v1

import (
	metav1 "github.com/octohelm/objectkind/pkg/apis/meta/v1"
)

type DatabaseCode string

// +gengo:objectkind
type Database struct {
	metav1.TypeMeta
	metav1.Codable[DatabaseCode]

	Spec DatabaseSpec `json:"spec,omitzero"`
}

type DatabaseSpec struct {
	CharacterType    string `json:"characterType,omitzero"`
	Collation        string `json:"collation,omitzero"`
	CollationVersion string `json:"collationVersion,omitzero"`
}
