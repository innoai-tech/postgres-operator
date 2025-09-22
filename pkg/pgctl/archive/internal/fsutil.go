package internal

import (
	"io/fs"
	"path/filepath"
)

type WalkFunc func(path string, entry fs.DirEntry) bool

func Include(fsys fs.FS, include WalkFunc) fs.FS {
	return &partialFS{
		FS:      fsys,
		include: include,
	}
}

func Exclude(fsys fs.FS, include WalkFunc) fs.FS {
	return &partialFS{
		FS: fsys,
		include: func(path string, entry fs.DirEntry) bool {
			return !include(path, entry)
		},
	}
}

type partialFS struct {
	fs.FS
	include WalkFunc
}

func (efs *partialFS) ReadDir(name string) ([]fs.DirEntry, error) {
	entries, err := fs.ReadDir(efs.FS, name)
	if err != nil {
		return nil, err
	}

	filteredEntries := make([]fs.DirEntry, 0, len(entries))
	for _, entry := range entries {
		fullPath := filepath.Join(name, entry.Name())
		if efs.include(fullPath, entry) {
			filteredEntries = append(filteredEntries, entry)
		}
	}
	return filteredEntries, nil
}
