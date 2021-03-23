package flag

import (
	"strings"
)

// ListFlags is dynamic flag types, as a slice
type ListFlags []string

func (fields *ListFlags) String() string {
	return strings.Join(*fields, ", ")
}

// Set is used to add field into fields
func (fields *ListFlags) Set(value string) error {
	*fields = append(*fields, value)

	return nil
}
