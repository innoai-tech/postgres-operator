package v1

import (
	metav1 "github.com/octohelm/objectkind/pkg/apis/meta/v1"
	sqltypetime "github.com/octohelm/storage/pkg/sqltype/time"
)

type ArchiveList = metav1.List[Archive]

// Archive
// +gengo:objectkind
type Archive struct {
	metav1.Codable[ArchiveCode]
	metav1.Metadata

	Files []*File `json:"files,omitzero"`
}

type File struct {
	Name           string                `json:"name"`
	Size           int64                 `json:"size,omitzero"`
	LastModifiedAt sqltypetime.Timestamp `json:"lastModifiedAt,omitzero"`
}
