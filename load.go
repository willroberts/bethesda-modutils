package modutils

import "os"

type ModFile struct {
	Metadata *Record
	Groups   []*Group

	rawBytes []byte
}

func Load(filename string) (*ModFile, error) {
	b, err := os.ReadFile(filename)
	if err != nil {
		return nil, err
	}

	return &ModFile{
		rawBytes: b,
	}, nil
}
