package modutils

import (
	"bytes"
	"encoding/binary"
	"testing"
)

func TestReadBytes(t *testing.T) {
	b := []byte{0, 1, 2, 3}
	r := bytes.NewReader(b)
	_, err := readBytes(4, r)
	if err != nil {
		t.Error(err)
	}
}

func TestReadBytesOversized(t *testing.T) {
	b := []byte{0, 1, 2}
	r := bytes.NewReader(b)
	_, err := readBytes(4, r)
	if err == nil {
		t.Error("failed to detect error when reading too many bytes")
	}
}

func TestReadUint16(t *testing.T) {
	in := uint16(1234)
	b := make([]byte, 2)
	binary.LittleEndian.PutUint16(b, in)
	r := bytes.NewReader(b)
	out, err := readUint16(r)
	if err != nil {
		t.Error(err)
	}
	if in != out {
		t.Error("expected:", in, "got:", out)
	}
}

func TestReadUint16Oversized(t *testing.T) {
	b := make([]byte, 1)
	r := bytes.NewReader(b)
	_, err := readUint16(r)
	if err == nil {
		t.Error("failed to detect error when reading too many bytes")
	}
}

func TestReadUint32(t *testing.T) {
	in := uint32(1234)
	b := make([]byte, 4)
	binary.LittleEndian.PutUint32(b, in)
	r := bytes.NewReader(b)
	out, err := readUint32(r)
	if err != nil {
		t.Error(err)
	}
	if in != out {
		t.Error("expected:", in, "got:", out)
	}
}

func TestReadUint32Oversized(t *testing.T) {
	b := make([]byte, 3)
	r := bytes.NewReader(b)
	_, err := readUint32(r)
	if err == nil {
		t.Error("failed to detect error when reading too many bytes")
	}
}
