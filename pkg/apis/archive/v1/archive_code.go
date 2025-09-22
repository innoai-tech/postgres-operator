package v1

import (
	"strings"
	"time"
)

const ArchiveCodeTimestampFormat = "20060102150405"

func NewArchiveCode(t time.Time, cat string) ArchiveCode {
	return ArchiveCode(t.Format(ArchiveCodeTimestampFormat) + "." + cat)
}

type ArchiveCode string

func (code ArchiveCode) Time() time.Time {
	t, _ := time.Parse(ArchiveCodeTimestampFormat, strings.Split(string(code), ".")[0])
	return t
}

func (code ArchiveCode) Cat() string {
	if i := strings.Index(string(code), "."); i > 0 {
		return string(code[i+len("."):])
	}
	return ""
}
