package modutils

import (
	"fmt"
	"io"
)

type Group struct {
	RecordType []byte
	Size       uint32
	Label      []byte
	GroupType  []byte
	Timestamp  uint16
	VCSInfo    uint16
	Unknown    uint32
	Data       []byte
}

func ReadGroup(r io.Reader) (*Group, error) {
	g := &Group{}
	var err error

	g.RecordType, err = readBytes(4, r)
	if err != nil {
		return nil, err
	}

	s := string(g.RecordType)
	if s != "GRUP" {
		return nil, fmt.Errorf("expected record type GRUP; got %s", s)
	}

	g.Size, err = readUint32(r)
	if err != nil {
		return nil, err
	}

	g.Label, err = readBytes(4, r)
	if err != nil {
		return nil, err
	}

	g.GroupType, err = readBytes(4, r)
	if err != nil {
		return nil, err
	}

	g.Timestamp, err = readUint16(r)
	if err != nil {
		return nil, err
	}

	g.VCSInfo, err = readUint16(r)
	if err != nil {
		return nil, err
	}

	g.Unknown, err = readUint32(r)
	if err != nil {
		return nil, err
	}

	g.Data, err = readBytes(uint(g.Size)-24, r)
	if err != nil {
		return nil, err
	}

	return g, nil
}

func (g *Group) Print() {
	fmt.Println("Group Record Type:", string(g.RecordType))
	fmt.Println("Group Size:", g.Size)
	fmt.Println("Group Label:", string(g.Label))
	fmt.Println("Group Type:", string(g.GroupType))
	fmt.Println("Group Timestamp:", g.Timestamp)
	fmt.Println("Group VCSInfo:", g.VCSInfo)
	//fmt.Println("Group Data:", string(g.Data))
}
