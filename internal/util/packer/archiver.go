package packer

import (
	"archive/tar"
	"bytes"
	"compress/gzip"
)

type ArchiveArgs struct {
	Name string
	Data []byte
}

type Archiver interface {
	Append(args ...ArchiveArgs)
	Archive() ([]byte, error)
}

type Arch struct {
	inputList []ArchiveArgs
}

func NewArch() Archiver {
	return &Arch{}
}

func (a *Arch) Append(args ...ArchiveArgs) {
	a.inputList = append(a.inputList, args...)
}

func (a *Arch) Archive() ([]byte, error) {
	var buf bytes.Buffer
	tw := tar.NewWriter(&buf)

	for _, arg := range a.inputList {
		hdr := &tar.Header{
			Name: arg.Name,
			Mode: 0644,
			Size: int64(len(arg.Data)),
		}
		if err := tw.WriteHeader(hdr); err != nil {
			return nil, err
		}
		if _, err := tw.Write(arg.Data); err != nil {
			return nil, err
		}
	}

	if err := tw.Close(); err != nil {
		return nil, err
	}

	return buf.Bytes(), nil
}

func Comporess(data []byte) ([]byte, error) {
	var buf bytes.Buffer
	zw := gzip.NewWriter(&buf)

	_, err := zw.Write(data)
	if err != nil {
		return nil, err
	}

	if err = zw.Close(); err != nil {
		return nil, err
	}

	return buf.Bytes(), nil
}
