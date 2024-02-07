package modutils

import (
	"bytes"
	"io"
	"os"
)

// ModFile represents an ESM, ESP, or ESL file, based on the NetImmerse IFF format.
type ModFile struct {
	Size     uint
	Metadata *Record
	Groups   []*Group

	rawData io.Reader
}

// LoadModFile will parse and return the given file.
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
