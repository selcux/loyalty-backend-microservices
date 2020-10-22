package packer

import (
	"archive/tar"
	"bytes"
	"compress/gzip"
	"io/ioutil"
	"os"
	"path"
)

type Packer interface {
	Add(inputFilePaths ...string)
	AddCompress(name string, inputFilePaths ...string) error
	Flush(outFilePath string) error
	compressToMemory(inFilePaths ...string) ([]byte, error)
	prepareData(inFilePath string) (*tarFileData, error)
}

type tarFileData struct {
	name string
	size int64
	body []byte
}

type CcPack struct {
	files           []string
	compressedFiles map[string][]byte
}

func NewCcPack() *CcPack {
	return &CcPack{
		files:           make([]string, 0),
		compressedFiles: make(map[string][]byte),
	}
}

func (c *CcPack) Add(inputFilePaths ...string) {
	for _, filePath := range inputFilePaths {
		c.files = append(c.files, filePath)
	}
}

func (c *CcPack) AddCompress(name string, inputFilePaths ...string) error {
	archive, err := c.compressToMemory(inputFilePaths...)
	if err != nil {
		return err
	}

	var buf bytes.Buffer
	zw := gzip.NewWriter(&buf)

	_, err = zw.Write(archive)
	if err != nil {
		return err
	}

	if err = zw.Close(); err != nil {
		return err
	}

	c.compressedFiles[name] = buf.Bytes()

	return nil
}

func (c *CcPack) Flush(outFilePath string) error {
	var tarbuf bytes.Buffer
	tw := tar.NewWriter(&tarbuf)

	for _, fPath := range c.files {
		fileData, err := c.prepareData(fPath)
		if err != nil {
			return err
		}

		hdr := &tar.Header{
			Name: fileData.name,
			Mode: 0644,
			Size: fileData.size,
		}
		if err := tw.WriteHeader(hdr); err != nil {
			return err
		}
		if _, err := tw.Write(fileData.body); err != nil {
			return err
		}
	}

	for arcName, arcBody := range c.compressedFiles {
		hdr := &tar.Header{
			Name: arcName,
			Mode: 0644,
			Size: int64(len(arcBody)),
		}
		if err := tw.WriteHeader(hdr); err != nil {
			return err
		}
		if _, err := tw.Write(arcBody); err != nil {
			return err
		}
	}

	if err := tw.Close(); err != nil {
		return err
	}

	var zbuf bytes.Buffer
	zw := gzip.NewWriter(&zbuf)

	_, err := zw.Write(tarbuf.Bytes())
	if err != nil {
		return err
	}

	if err = zw.Close(); err != nil {
		return err
	}

	return c.WriteTo(outFilePath, zbuf.Bytes())
}

func (c *CcPack) compressToMemory(inFilePaths ...string) ([]byte, error) {
	var buf bytes.Buffer
	tw := tar.NewWriter(&buf)

	for _, inFile := range inFilePaths {
		fileData, err := c.prepareData(inFile)
		if err != nil {
			return nil, err
		}

		hdr := &tar.Header{
			Name: fileData.name,
			Mode: 0644,
			Size: fileData.size,
		}
		if err := tw.WriteHeader(hdr); err != nil {
			return nil, err
		}
		if _, err := tw.Write(fileData.body); err != nil {
			return nil, err
		}
	}

	if err := tw.Close(); err != nil {
		return nil, err
	}

	return buf.Bytes(), nil
}

func (c *CcPack) WriteTo(outFilePath string, data []byte) error {
	return ioutil.WriteFile(outFilePath, data, 0644)
}

func (c *CcPack) prepareData(inFilePath string) (*tarFileData, error) {
	f, err := os.Open(inFilePath)
	if err != nil {
		return nil, err
	}

	_, name := path.Split(inFilePath)
	body, err := ioutil.ReadAll(f)
	if err != nil {
		return nil, err
	}

	data := &tarFileData{
		name: name,
		size: int64(len(body)),
		body: body,
	}

	return data, nil
}
