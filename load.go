package modutils

import (
	"bytes"
	"io"
	"os"
)

type ModFile struct {
	Metadata *Record
	Groups   []*Group

	rawBytes io.Reader
}

func Load(filename string) (*ModFile, error) {
	m := &ModFile{}

	b, err := os.ReadFile(filename)
	if err != nil {
		return nil, err
	}
	m.rawBytes = bytes.NewReader(b)

	m.Metadata, err = ReadRecord(m.rawBytes)
	if err != nil {
		return nil, err
	}

	return m, nil
}
