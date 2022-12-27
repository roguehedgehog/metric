package importdata

import (
	"fmt"
	"os"
)

type ImportType string

const (
	ImportTypeLongHistory = "endsong"
)

type ImportFile struct {
	file *os.File
	what ImportType
}

func NewImportFile(file *os.File, what ImportType) (*ImportFile, error) {
	if what != ImportTypeLongHistory {
		return nil, fmt.Errorf("import of %s is not supported, supported type %s", what, ImportTypeLongHistory)
	}

	return &ImportFile{
		file, what,
	}, nil
}
