package modutils

import (
	"bytes"
	"io"
	"os"
)

type ModFile struct {
	Size     uint
	Metadata *Record
	Groups   []*Group

	rawData io.Reader
}

func LoadModFile(filename string) (*ModFile, error) {
	m := &ModFile{}

	b, err := os.ReadFile(filename)
	if err != nil {
		return nil, err
	}
	m.Size = uint(len(b))
	m.rawData = bytes.NewReader(b)

	m.Metadata, err = ReadRecord(m.rawData)
	if err != nil {
		return nil, err
	}

	if err := m.readAllGroups(); err != nil {
		return nil, err
	}

	return m, nil
}

func (m *ModFile) readAllGroups() error {
	groups := make([]*Group, 0)
	var bytesRead uint = 24 + uint(m.Metadata.Size)

	for bytesRead < uint(m.Size) {
		g, err := ReadGroup(m.rawData)
		if err != nil {
			return err
		}
		groups = append(groups, g)
		bytesRead += uint(g.Size)
	}

	m.Groups = groups
	return nil
}
